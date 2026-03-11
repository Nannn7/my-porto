package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"myporto-backend/config"
	"myporto-backend/migration"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "up", "down", "status":
		runMigrationCommand(command, os.Args[2:])
	case "create":
		runCreateCommand(os.Args[2:])
	default:
		printUsage()
		os.Exit(1)
	}
}

func runMigrationCommand(command string, args []string) {
	stepsDefault := 1
	if command == "up" {
		stepsDefault = 0
	}

	fs := flag.NewFlagSet(command, flag.ExitOnError)
	dir := fs.String("dir", detectMigrationsDir(), "migration directory")
	steps := fs.Int("steps", stepsDefault, "number of migration steps (0 means all for up)")
	fs.Parse(args)

	if err := config.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	migrator, err := migration.NewMigrator(config.DB, *dir)
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	case "up":
		applied, err := migrator.Up(*steps)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Applied %d migration(s)\n", applied)
	case "down":
		rolledBack, err := migrator.Down(*steps)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Rolled back %d migration(s)\n", rolledBack)
	case "status":
		items, err := migrator.Status()
		if err != nil {
			log.Fatal(err)
		}
		printStatus(items)
	}
}

func runCreateCommand(args []string) {
	fs := flag.NewFlagSet("create", flag.ExitOnError)
	dir := fs.String("dir", detectMigrationsDir(), "migration directory")
	fs.Parse(args)

	if fs.NArg() != 1 {
		log.Fatal("usage: go run ./cmd/migrate create <migration_name>")
	}

	upPath, downPath, err := migration.CreateMigrationFiles(*dir, fs.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created migration files:\n- %s\n- %s\n", upPath, downPath)
}

func printStatus(items []migration.StatusItem) {
	fmt.Printf("%-16s %-30s %-10s %s\n", "Version", "Name", "State", "Applied At")
	for _, item := range items {
		appliedAt := "-"
		if item.AppliedAt != nil {
			appliedAt = item.AppliedAt.Format(time.RFC3339)
		}

		fmt.Printf("%-16d %-30s %-10s %s\n", item.Version, item.Name, item.State, appliedAt)
	}
}

func detectMigrationsDir() string {
	candidates := []string{
		"migrations",
		filepath.Join("backend", "migrations"),
	}

	for _, candidate := range candidates {
		info, err := os.Stat(candidate)
		if err == nil && info.IsDir() {
			return candidate
		}
	}

	return "migrations"
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  go run ./cmd/migrate up [--steps N] [--dir migrations]")
	fmt.Println("  go run ./cmd/migrate down [--steps N] [--dir migrations]")
	fmt.Println("  go run ./cmd/migrate status [--dir migrations]")
	fmt.Println("  go run ./cmd/migrate create <migration_name> [--dir migrations]")
}
