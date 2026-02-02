# 🛒 Kasir API - Clean Architecture Edition

API manajemen kasir modern yang dibangun dengan bahasa **Go (Golang)**. Proyek ini mendemonstrasikan implementasi **Clean Architecture**, dokumentasi otomatis dengan **Swagger**, serta integrasi **Cloud Database (Supabase)** dan **Cloud Hosting (Railway)**.

Ini merupakan bootcamp series yang diinisialisai mas Ariaseta & Umam (SUMOPOD)

---

## 🏗️ Arsitektur & Alur Sistem

Aplikasi ini menggunakan prinsip **Clean Architecture** untuk memastikan kode mudah diuji dan dikembangkan. Alur data bergerak dari luar ke dalam:
1. **Handler**: Mengelola HTTP request & response.
2. **Service**: Berisi logika bisnis (Business Logic).
3. **Repository**: Berkomunikasi langsung dengan Database (SQL).
4. **Models**: Definisi struktur data (Schema).



---

## 🛠️ Langkah 1: Persiapan Alat (Prasyarat)

Pastikan alat berikut sudah terinstal di komputer Anda:
* **Go (v1.23+)**: [Download di sini](https://go.dev/dl/).
* **Git**: Untuk manajemen kode.
* **Swaggo CLI**: Untuk dokumentasi API. Instal dengan:
  ```bash
  go install [github.com/swaggo/swag/cmd/swag@latest](https://github.com/swaggo/swag/cmd/swag@latest)

---

## 🗄️ Langkah 2: Setup Database (Supabase)

Aplikasi ini menggunakan relasi One-to-Many antara Kategori dan Produk. Jalankan script SQL berikut di SQL Editor Supabase Anda:

1. Tabel Kategori
  ```bash
    CREATE TABLE categories (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL
    );

2. Tabel Produk dengan Foreign Key ke Kategori
  ```bash
    CREATE TABLE products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        price NUMERIC(15, 2) NOT NULL,
        stock INTEGER NOT NULL,
        category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL
    );

3. Masukkan Data Contoh
```bash
    INSERT INTO categories (name) VALUES ('Elektronik'), ('Makanan');



---

## 🚀 Langkah 3: Menjalankan di Lokal

1. Clone & Install

```bash
    git clone [https://github.com/goprex/crudAPI.git](https://github.com/goprex/crudAPI.git)
    cd crudAPI
    go mod tidy

2. Setup Environment: Buat file .env di root folder
```bash
    PORT=8080
    DB_CONN=postgres://postgres:[PASSWORD_MU]@db.supabase.co:5432/postgres

3. Generate Swagger & Run
```bash
    swag init
    go run main.go

Akses dokumentasi di: http://localhost:8080/docs/index.html

---
## ☁️ Langkah 4: Deployment ke Railway (Online)

1. Push kode terbaru Anda ke GitHub.
2. Di Railway, pilih New Project > Deploy from GitHub.
3. Pilih repository crudAPI.
4 . Atur Variables (PENTING): Masuk ke tab Variables di dashboard Railway, lalu tambahkan:
```bash
    DB_CONN: (Isi dengan Link URI Supabase Anda)
    PORT: 8080

--- 
## 5. 📖 Testing API (CURL)

1. Tambah Produk Baru (POST)
```bash
    curl -X POST http://localhost:8080/api/produk \
         -H "Content-Type: application/json" \
         -d '{
               "name": "Mouse Gaming",
               "price": 250000,
               "stock": 10,
               "category_id": 1
             }'

2. Update Produk (PUT)
```bash
    curl -X PUT http://localhost:8080/api/produk/1 \
         -H "Content-Type: application/json" \
         -d '{
               "name": "Mouse Gaming Wireless",
               "price": 300000,
               "stock": 5,
               "category_id": 1
             }'

3. Hapus Produk (DELETE)
```bash
    curl -X DELETE http://localhost:8080/api/produk/1

4. Ambil Semua Produk (GET)
```bash
    curl -X GET http://localhost:8080/api/produk

