# Deploy Guide: Vercel + Render + Neon

Stack ini cocok kalau kamu belum punya VPS/server sendiri:
- Frontend: Vercel
- Backend: Render
- Database: Neon PostgreSQL

## 1. Push repo ke GitHub
Pastikan repo ini sudah ada di GitHub karena Vercel dan Render akan ambil source dari sana.

## 2. Buat database di Neon
1. Login ke Neon.
2. Buat project/database baru.
3. Copy connection string Postgres dari Neon.

Catatan:
- Backend sekarang sudah support `DATABASE_URL`, jadi kamu tidak perlu pecah manual ke `DB_HOST`, `DB_PORT`, dan seterusnya.
- Untuk Neon, gunakan connection string yang sudah menyertakan SSL.

## 3. Deploy backend ke Render
File [render.yaml](../render.yaml) sudah disiapkan.

### Opsi paling gampang: Blueprint
1. Login ke Render.
2. Klik `New +` -> `Blueprint`.
3. Pilih repository GitHub kamu.
4. Render akan membaca `render.yaml`.
5. Saat diminta, isi environment variable berikut:
   - `DATABASE_URL` = connection string dari Neon
   - `DEFAULT_ADMIN_USERNAME` = username admin kamu
   - `DEFAULT_ADMIN_PASSWORD` = password admin kamu
   - `CORS_ALLOWED_ORIGINS` = isi sementara dengan `http://localhost:5173`
6. Klik `Apply`.

Apa yang sudah otomatis dari config:
- Root project backend diarahkan ke folder `backend/`
- Render build binary server dan migrator
- Saat service start, migration akan dijalankan dulu
- Health check pakai `/api/health`

Setelah deploy sukses, simpan URL backend Render, contoh:
- `https://my-porto-api.onrender.com`

## 4. Deploy frontend ke Vercel
File [vercel.json](../frontend/vercel.json) sudah disiapkan supaya route SPA seperti `/about` dan `/admin/login` tidak 404 saat refresh.

1. Login ke Vercel.
2. Klik `Add New...` -> `Project`.
3. Import repository GitHub yang sama.
4. Saat setup project:
   - `Root Directory` = `frontend`
   - Framework akan terdeteksi sebagai `Vite`
   - Build Command = `npm run build`
   - Output Directory = `dist`
5. Tambahkan environment variable:
   - `VITE_API_BASE_URL` = `https://my-porto-api.onrender.com/api`
6. Deploy.

Setelah deploy selesai, simpan URL frontend Vercel, contoh:
- `https://my-porto.vercel.app`

## 5. Balik ke Render dan update CORS
Supaya frontend production bisa call API:
1. Buka service backend di Render.
2. Masuk ke `Environment`.
3. Ubah `CORS_ALLOWED_ORIGINS` menjadi:

```env
http://localhost:5173,https://my-porto.vercel.app
```

4. Save changes lalu redeploy.

Kalau nanti kamu pakai custom domain, ganti domain Vercel di atas dengan domain final kamu.

## 6. Test production
Checklist cepat:
- Buka frontend Vercel
- Cek halaman `/projects`
- Test form `/contact`
- Login ke `/admin/login`
- Pastikan CRUD project berjalan

## Env yang dipakai

### Backend Render
Minimal:

```env
APP_TLS_ENABLED=false
APP_AUTO_MIGRATE=false
ADMIN_SESSION_TTL_HOURS=24
DATABASE_URL=postgresql://...
DEFAULT_ADMIN_USERNAME=admin
DEFAULT_ADMIN_PASSWORD=your-strong-password
CORS_ALLOWED_ORIGINS=https://your-frontend-domain.vercel.app
```

### Frontend Vercel

```env
VITE_API_BASE_URL=https://your-backend-domain.onrender.com/api
```

## Kalau mau deploy manual tanpa Blueprint
Di Render, kamu bisa bikin `Web Service` biasa dengan setting:
- Runtime: `Go`
- Root Directory: `backend`
- Build Command: `go build -o bin/myporto-api ./cmd/server && go build -o bin/migrate ./cmd/migrate`
- Start Command: `/bin/sh -c "./bin/migrate up && ./bin/myporto-api"`
- Health Check Path: `/api/health`

## Catatan penting
- Free web service di Render bisa sleep saat idle, jadi request pertama bisa lambat.
- Kalau nanti app makin serius, pertimbangkan upgrade backend plan atau pindah ke Railway/Fly.io.
- Untuk custom domain, set domain frontend di Vercel dan backend di Render, lalu update `CORS_ALLOWED_ORIGINS` serta `VITE_API_BASE_URL`.
