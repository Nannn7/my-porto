package migration

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

var upFilePattern = regexp.MustCompile(`^([0-9]{14})_(.+)\.up\.sql$`)
var nameSanitizer = regexp.MustCompile(`[^a-z0-9_]+`)

type FileMigration struct {
	Version  int64
	Name     string
	UpPath   string
	DownPath string
	UpSQL    string
	DownSQL  string
}

type StatusItem struct {
	Version   int64
	Name      string
	State     string
	AppliedAt *time.Time
}

type appliedRecord struct {
	Version   int64     `gorm:"column:version"`
	Name      string    `gorm:"column:name"`
	AppliedAt time.Time `gorm:"column:applied_at"`
}

type Migrator struct {
	db  *gorm.DB
	dir string
}

func NewMigrator(db *gorm.DB, dir string) (*Migrator, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}

	if strings.TrimSpace(dir) == "" {
		return nil, errors.New("migration directory is required")
	}

	return &Migrator{
		db:  db,
		dir: dir,
	}, nil
}

func (m *Migrator) EnsureMigrationsTable() error {
	query := `
CREATE TABLE IF NOT EXISTS schema_migrations (
  version BIGINT PRIMARY KEY,
  name TEXT NOT NULL,
  applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)`

	return m.db.Exec(query).Error
}

func (m *Migrator) LoadMigrations() ([]FileMigration, error) {
	pattern := filepath.Join(m.dir, "*.up.sql")
	upFiles, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	migrations := make([]FileMigration, 0, len(upFiles))
	for _, upFile := range upFiles {
		base := filepath.Base(upFile)
		matches := upFilePattern.FindStringSubmatch(base)
		if len(matches) != 3 {
			continue
		}

		version, parseErr := strconv.ParseInt(matches[1], 10, 64)
		if parseErr != nil {
			return nil, fmt.Errorf("invalid migration version in %s: %w", base, parseErr)
		}

		name := matches[2]
		downFile := filepath.Join(m.dir, fmt.Sprintf("%s_%s.down.sql", matches[1], name))
		downSQL := ""
		if content, readErr := os.ReadFile(downFile); readErr == nil {
			downSQL = string(content)
		}

		upContent, readErr := os.ReadFile(upFile)
		if readErr != nil {
			return nil, fmt.Errorf("failed to read %s: %w", upFile, readErr)
		}

		migrations = append(migrations, FileMigration{
			Version:  version,
			Name:     name,
			UpPath:   upFile,
			DownPath: downFile,
			UpSQL:    string(upContent),
			DownSQL:  downSQL,
		})
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

func (m *Migrator) AppliedMigrations() ([]appliedRecord, error) {
	if err := m.EnsureMigrationsTable(); err != nil {
		return nil, err
	}

	var records []appliedRecord
	if err := m.db.Raw(
		"SELECT version, name, applied_at FROM schema_migrations ORDER BY version ASC",
	).Scan(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (m *Migrator) Up(steps int) (int, error) {
	if err := m.EnsureMigrationsTable(); err != nil {
		return 0, err
	}

	migrations, err := m.LoadMigrations()
	if err != nil {
		return 0, err
	}

	applied, err := m.AppliedMigrations()
	if err != nil {
		return 0, err
	}

	appliedMap := make(map[int64]struct{}, len(applied))
	for _, item := range applied {
		appliedMap[item.Version] = struct{}{}
	}

	appliedCount := 0
	for _, migration := range migrations {
		if _, exists := appliedMap[migration.Version]; exists {
			continue
		}

		if strings.TrimSpace(migration.UpSQL) == "" {
			return appliedCount, fmt.Errorf("migration %d_%s has empty up SQL", migration.Version, migration.Name)
		}

		tx := m.db.Begin()
		if tx.Error != nil {
			return appliedCount, tx.Error
		}

		if err := tx.Exec(migration.UpSQL).Error; err != nil {
			tx.Rollback()
			return appliedCount, fmt.Errorf("failed applying up migration %d_%s: %w", migration.Version, migration.Name, err)
		}

		if err := tx.Exec(
			"INSERT INTO schema_migrations (version, name, applied_at) VALUES (?, ?, NOW())",
			migration.Version,
			migration.Name,
		).Error; err != nil {
			tx.Rollback()
			return appliedCount, fmt.Errorf("failed recording migration %d_%s: %w", migration.Version, migration.Name, err)
		}

		if err := tx.Commit().Error; err != nil {
			return appliedCount, err
		}

		appliedCount++
		if steps > 0 && appliedCount >= steps {
			break
		}
	}

	return appliedCount, nil
}

func (m *Migrator) Down(steps int) (int, error) {
	if err := m.EnsureMigrationsTable(); err != nil {
		return 0, err
	}

	if steps <= 0 {
		steps = 1
	}

	migrations, err := m.LoadMigrations()
	if err != nil {
		return 0, err
	}

	byVersion := make(map[int64]FileMigration, len(migrations))
	for _, item := range migrations {
		byVersion[item.Version] = item
	}

	var applied []appliedRecord
	if err := m.db.Raw(
		"SELECT version, name, applied_at FROM schema_migrations ORDER BY version DESC",
	).Scan(&applied).Error; err != nil {
		return 0, err
	}

	if len(applied) == 0 {
		return 0, nil
	}

	if steps < len(applied) {
		applied = applied[:steps]
	}

	rolledBack := 0
	for _, record := range applied {
		migration, exists := byVersion[record.Version]
		if !exists {
			return rolledBack, fmt.Errorf("cannot rollback version %d: migration file missing", record.Version)
		}

		if strings.TrimSpace(migration.DownSQL) == "" {
			return rolledBack, fmt.Errorf("cannot rollback version %d: down SQL file missing or empty", record.Version)
		}

		tx := m.db.Begin()
		if tx.Error != nil {
			return rolledBack, tx.Error
		}

		if err := tx.Exec(migration.DownSQL).Error; err != nil {
			tx.Rollback()
			return rolledBack, fmt.Errorf("failed applying down migration %d_%s: %w", migration.Version, migration.Name, err)
		}

		if err := tx.Exec("DELETE FROM schema_migrations WHERE version = ?", record.Version).Error; err != nil {
			tx.Rollback()
			return rolledBack, err
		}

		if err := tx.Commit().Error; err != nil {
			return rolledBack, err
		}

		rolledBack++
	}

	return rolledBack, nil
}

func (m *Migrator) Status() ([]StatusItem, error) {
	migrations, err := m.LoadMigrations()
	if err != nil {
		return nil, err
	}

	applied, err := m.AppliedMigrations()
	if err != nil {
		return nil, err
	}

	appliedMap := make(map[int64]appliedRecord, len(applied))
	for _, item := range applied {
		appliedMap[item.Version] = item
	}

	statuses := make([]StatusItem, 0, len(migrations)+len(applied))
	fileVersions := make(map[int64]struct{}, len(migrations))

	for _, migration := range migrations {
		fileVersions[migration.Version] = struct{}{}
		if record, ok := appliedMap[migration.Version]; ok {
			appliedAt := record.AppliedAt
			statuses = append(statuses, StatusItem{
				Version:   migration.Version,
				Name:      migration.Name,
				State:     "applied",
				AppliedAt: &appliedAt,
			})
			continue
		}

		statuses = append(statuses, StatusItem{
			Version: migration.Version,
			Name:    migration.Name,
			State:   "pending",
		})
	}

	for _, record := range applied {
		if _, exists := fileVersions[record.Version]; exists {
			continue
		}

		appliedAt := record.AppliedAt
		statuses = append(statuses, StatusItem{
			Version:   record.Version,
			Name:      record.Name,
			State:     "orphaned",
			AppliedAt: &appliedAt,
		})
	}

	sort.Slice(statuses, func(i, j int) bool {
		return statuses[i].Version < statuses[j].Version
	})

	return statuses, nil
}

func CreateMigrationFiles(dir, rawName string) (string, string, error) {
	name := normalizeMigrationName(rawName)
	if name == "" {
		return "", "", errors.New("migration name is invalid")
	}

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", "", err
	}

	version := time.Now().UTC().Format("20060102150405")
	base := fmt.Sprintf("%s_%s", version, name)

	upPath := filepath.Join(dir, base+".up.sql")
	downPath := filepath.Join(dir, base+".down.sql")

	if _, err := os.Stat(upPath); err == nil {
		return "", "", fmt.Errorf("file already exists: %s", upPath)
	}

	if _, err := os.Stat(downPath); err == nil {
		return "", "", fmt.Errorf("file already exists: %s", downPath)
	}

	upTemplate := "-- Write your UP migration SQL here.\n"
	downTemplate := "-- Write your DOWN migration SQL here.\n"

	if err := os.WriteFile(upPath, []byte(upTemplate), 0o644); err != nil {
		return "", "", err
	}

	if err := os.WriteFile(downPath, []byte(downTemplate), 0o644); err != nil {
		return "", "", err
	}

	return upPath, downPath, nil
}

func normalizeMigrationName(rawName string) string {
	normalized := strings.ToLower(strings.TrimSpace(rawName))
	normalized = strings.ReplaceAll(normalized, "-", "_")
	normalized = strings.ReplaceAll(normalized, " ", "_")
	normalized = nameSanitizer.ReplaceAllString(normalized, "_")
	normalized = strings.Trim(normalized, "_")
	return normalized
}
