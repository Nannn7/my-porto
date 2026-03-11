# my-porto

Monorepo portfolio dengan:
- Frontend: Vue 3 + Vite + Pinia (`frontend/`)
- Backend: Go + Gin + GORM + PostgreSQL (`backend/`)

## Fitur Utama
- Public portfolio pages (`/`, `/about`, `/projects`, `/contact`)
- Public API: `GET /api/projects`
- Admin auth berbasis session token (Bearer)
- Admin API CRUD project dengan validasi payload
- Admin panel frontend:
  - `/admin/login`
  - `/admin/projects` (list, create, edit, delete)
- Konfigurasi sensitif dipindah ke environment variable
- CORS diperketat via whitelist origin
- Template reverse proxy production (Nginx)

## Struktur Folder
```txt
my-porto/
+-- backend/
+-- frontend/
+-- deploy/
    +-- nginx/
        +-- my-porto.conf
```

## Prasyarat
- Node.js `^20.19.0 || >=22.12.0`
- npm
- Go `1.25.x`
- PostgreSQL

## Setup Environment

### Backend
1. Copy contoh env:
```bash
cd backend
cp .env.example .env
```
2. Isi nilai `.env` sesuai environment.

Variable penting backend:
- `APP_PORT` (default `8080`)
- `APP_TLS_ENABLED` (`false` untuk mode reverse proxy)
- `APP_AUTO_MIGRATE` (`false` direkomendasikan; aktifkan `true` hanya jika butuh AutoMigrate GORM)
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME`, `DB_SSLMODE`
- `DEFAULT_ADMIN_USERNAME`, `DEFAULT_ADMIN_PASSWORD`
- `ADMIN_SESSION_TTL_HOURS`
- `CORS_ALLOWED_ORIGINS` (comma-separated)

### Frontend
1. Copy contoh env:
```bash
cd frontend
cp .env.example .env
```
2. Atur:
- `VITE_API_BASE_URL` (contoh local: `http://localhost:8080/api`)

## Menjalankan Development
Buka 2 terminal.

### Terminal 1: Backend
```bash
cd backend
go run ./cmd/server
```

### Terminal 2: Frontend
```bash
cd frontend
npm install
npm run dev
```

Frontend default: `http://localhost:5173`

## Migration (ala artisan migrate)
Command dijalankan dari folder `backend/`.

### Create migration file baru
```bash
go run ./cmd/migrate create add_project_slug
```

### Jalankan semua migration pending (`up`)
```bash
go run ./cmd/migrate up
```

### Jalankan migration per step
```bash
go run ./cmd/migrate up --steps 1
```

### Lihat status migration
```bash
go run ./cmd/migrate status
```

### Rollback migration (`down`)
```bash
go run ./cmd/migrate down --steps 1
```

Catatan:
- Disarankan `APP_AUTO_MIGRATE=false` dan pakai `cmd/migrate` untuk schema changes.
- Pastikan env database sudah benar sebelum menjalankan command migrate.

## Akun Admin Default
Akun admin default dibuat otomatis jika tabel admin masih kosong.
Nilainya diambil dari env:
- `DEFAULT_ADMIN_USERNAME`
- `DEFAULT_ADMIN_PASSWORD`

## API Ringkas
Base path: `/api`

### Public
- `GET /projects`
- `GET /skills`
- `POST /contact`

### Admin
- `POST /admin/login`
- `GET /admin/projects` (auth)
- `POST /admin/projects` (auth)
- `PUT /admin/projects/:id` (auth)
- `DELETE /admin/projects/:id` (auth)

### Header Auth Admin
```http
Authorization: Bearer <session_token>
```

## Validasi Payload Project
- `title`: required, min 3, max 120
- `description`: required, min 10, max 2000
- `tech_stack`: optional, max 300
- `image_url`: optional, URL valid
- `project_url`: optional, URL valid

## Production Reverse Proxy
Template Nginx tersedia di:
- `deploy/nginx/my-porto.conf`

Skema yang disarankan:
- Nginx handle TLS publik (`443`)
- Static frontend dari `frontend/dist`
- Path `/api/*` diproxy ke backend (`127.0.0.1:8080`)
- Backend jalan HTTP internal (`APP_TLS_ENABLED=false`)

## Build
### Frontend
```bash
cd frontend
npm run build
```

### Backend
```bash
cd backend
go build -o myporto-api ./cmd/server
```
