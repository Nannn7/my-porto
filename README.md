# my-porto

Monorepo portfolio dengan:
- Frontend: Vue 3 + Vite + Pinia (`frontend/`)
- Backend: Go + Gin + GORM + PostgreSQL (`backend/`)

## Fitur Utama
- Public portfolio pages (`/`, `/about`, `/projects`, `/contact`)
- Admin auth berbasis session token (Bearer)
- Admin API CRUD project dengan validasi payload
- Admin panel frontend:
  - `/admin/login`
  - `/admin/projects` (list, create, edit, delete)

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
- `PORT` (diprioritaskan di platform seperti Render/Railway)
- `APP_TLS_ENABLED` (`false` untuk mode reverse proxy)
- `APP_AUTO_MIGRATE` (`false` direkomendasikan; aktifkan `true` hanya jika butuh AutoMigrate GORM)
- `DATABASE_URL` (direkomendasikan untuk managed Postgres seperti Neon)
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

### Jalankan semua migration pending (`up`)
```bash
go run ./cmd/migrate up
```

### Lihat status migration
```bash
go run ./cmd/migrate status
```

Catatan:
- Disarankan `APP_AUTO_MIGRATE=false` dan pakai `cmd/migrate` untuk schema changes.
- Pastikan env database sudah benar sebelum menjalankan command migrate.

## Akun Admin Default
Akun admin default dibuat otomatis jika tabel admin masih kosong.
Nilainya diambil dari env:
- `DEFAULT_ADMIN_USERNAME`
- `DEFAULT_ADMIN_PASSWORD`

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