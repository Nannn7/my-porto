# my-porto

Monorepo untuk aplikasi portfolio yang terdiri dari:

- **Frontend**: Vue 3 + Vite + Pinia (folder `frontend/`)
- **Backend**: Go + Gin + GORM + PostgreSQL (folder `backend/`)

---

## 1) Arsitektur Monorepo

Struktur level atas:

```txt
my-porto/
├── frontend/   # SPA Vue (UI portfolio)
└── backend/    # REST API Go (data project)
```

### Alur data singkat

1. User membuka frontend (Vite dev server / static build).
2. Halaman `Projects` memanggil endpoint backend `/api/projects`.
3. Backend membaca data dari PostgreSQL melalui GORM.
4. Backend mengembalikan JSON ke frontend.

Catatan implementasi saat ini:

- Backend menjalankan **HTTPS langsung** via `RunTLS` di port `8080` dengan sertifikat lokal `backend/cert/cert.pem` + `backend/cert/key.pem`.
- CORS backend saat ini masih longgar (`Access-Control-Allow-Origin: *`) agar frontend lokal dapat akses API.

---

## 2) Prasyarat

### Global

- Git
- Node.js sesuai `frontend/package.json` (disarankan `^20.19.0` atau `>=22.12.0`)
- npm
- Go `1.25.x`
- PostgreSQL (lokal)

### Database default backend

Backend saat ini menggunakan DSN hardcoded:

- Host: `127.0.0.1`
- Port: `5432`
- User: `postgres`
- Password: `postgres`
- DB name: `myporto`
- SSL mode: `prefer`

Pastikan database `myporto` sudah dibuat sebelum backend dijalankan.

---

## 3) Install Dependency

### Frontend

```bash
cd frontend
npm install
```

### Backend

```bash
cd backend
go mod download
```

---

## 4) Menjalankan Project (Development)

Gunakan 2 terminal terpisah.

### Terminal A — Backend (HTTPS API)

```bash
cd backend
go run ./cmd/server
```

Server API aktif di `https://localhost:8080`.

### Terminal B — Frontend (Vite)

```bash
cd frontend
npm run dev
```

Frontend default di `http://localhost:5173`.

Jika backend menggunakan self-signed certificate, browser bisa menampilkan warning keamanan saat request API ke `https://localhost:8080`.

---

## 5) Menjalankan Mode Production-like

## Frontend

Build + preview lokal:

```bash
cd frontend
npm run build
npm run preview
```

Preview default di `http://localhost:4173`.

## Backend

Tidak ada script release khusus saat ini, jadi pendekatan production-like paling dekat:

```bash
cd backend
go run ./cmd/server
```

Untuk environment production nyata, disarankan compile binary terlebih dahulu:

```bash
cd backend
go build -o myporto-api ./cmd/server
./myporto-api
```

Tetap membutuhkan file sertifikat TLS (`cert/cert.pem` dan `cert/key.pem`) kecuali arsitektur TLS dipindah ke reverse proxy.

---

## 6) Environment Variable

## Frontend

| Variable | Wajib | Default | Keterangan |
|---|---|---|---|
| `VITE_API_BASE_URL` | Tidak | `https://localhost:8080/api` | Base URL API untuk mengambil data project di store frontend. |
| `BASE_URL` | Otomatis (Vite) | `/` | Dipakai router history (`import.meta.env.BASE_URL`). Biasanya tidak perlu diset manual. |

Contoh `.env` frontend:

```env
VITE_API_BASE_URL=https://localhost:8080/api
```

## Backend

Saat ini **belum ada** konfigurasi environment variable yang dipakai di source code backend. Seluruh DSN database masih hardcoded di `backend/config/database.go`.

Konsekuensi:

- Untuk ganti host/port/user/password DB, saat ini harus ubah source code.
- Untuk production, sangat disarankan migrasi ke ENV (`DB_HOST`, `DB_PORT`, dll) agar aman dan fleksibel.

---

## 7) API Reference

Base path: `/api`

## Public endpoint (sudah implementasi)

### `GET /api/projects`

Mengambil daftar project dari database.

Contoh request:

```bash
curl -k https://localhost:8080/api/projects
```

Contoh response sukses (`200`):

```json
{
  "message": "projects fetched",
  "data": [
    {
      "id": 1,
      "title": "Portfolio Website",
      "description": "Website portfolio pribadi"
    }
  ]
}
```

Contoh response gagal (`500`):

```json
{
  "message": "failed to fetch projects",
  "data": null
}
```

## Admin endpoint (belum implementasi)

Belum ada route `/api/admin/*` pada kode saat ini.

Roadmap endpoint admin (proposal kontrak API):

1. `POST /api/admin/login`
2. `GET /api/admin/projects`
3. `POST /api/admin/projects`
4. `PUT /api/admin/projects/:id`
5. `DELETE /api/admin/projects/:id`

Contoh payload login (rencana):

```json
{
  "email": "admin@my-porto.local",
  "password": "********"
}
```

Contoh payload create project (rencana):

```json
{
  "title": "Dashboard Monitoring",
  "description": "Admin dashboard untuk observability"
}
```

Contoh payload update project (rencana):

```json
{
  "title": "Dashboard Monitoring v2",
  "description": "Menambah fitur alerting"
}
```

---

## 8) Deployment Notes

## A. TLS certificate

Kondisi saat ini:

- Backend menjalankan HTTPS langsung via `RunTLS`.
- Sertifikat dibaca dari `backend/cert/cert.pem` dan `backend/cert/key.pem`.

Rekomendasi production:

- Jangan commit private key production ke repo.
- Gunakan sertifikat valid dari CA tepercaya (mis. Let's Encrypt).
- Atau terminasi TLS di reverse proxy, backend cukup jalan HTTP internal.

## B. Reverse proxy

Rekomendasi arsitektur deploy:

- Nginx/Caddy/Traefik sebagai entrypoint publik.
- Frontend static build dilayani proxy/web server.
- Request `/api` diteruskan ke backend service.

Contoh alur:

- `https://my-porto.com/` → frontend
- `https://my-porto.com/api/*` → backend (upstream internal)

## C. CORS

Kondisi saat ini backend:

- `Access-Control-Allow-Origin: *`

Rekomendasi production:

- Batasi origin spesifik (mis. `https://my-porto.com`)
- Pertahankan method/header seperlunya
- Pastikan preflight (`OPTIONS`) tetap ditangani dengan benar

---

## 9) Roadmap Fitur Bonus & Status

| Fitur Bonus | Status | Keterangan |
|---|---|---|
| Admin panel UI | ⏳ Belum implementasi | Belum ada halaman admin di frontend. |
| Admin login/auth | ⏳ Belum implementasi | Belum ada endpoint login / middleware auth. |
| CRUD project (admin) | ⏳ Belum implementasi | Backend saat ini baru read endpoint (`GET /api/projects`). |
| Public list project | ✅ Sudah implementasi | Endpoint publik tersedia dan terhubung ke DB. |

### Prioritas implementasi yang disarankan

1. Tambahkan auth admin (JWT/session) + endpoint login.
2. Tambahkan CRUD project di backend + validasi payload.
3. Buat admin panel frontend (login, list, create, edit, delete).
4. Pindahkan konfigurasi sensitif ke environment variable.
5. Perketat CORS + setup reverse proxy production.

---

## 10) Quick Start

```bash
# 1) Backend
cd backend
go run ./cmd/server

# 2) Frontend (terminal baru)
cd frontend
npm install
npm run dev
```

Buka `http://localhost:5173`, lalu pastikan endpoint API dapat diakses di `https://localhost:8080/api/projects`.
