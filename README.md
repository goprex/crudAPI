# 🛒 Kasir API - Clean Architecture Edition

API manajemen kasir modern yang dibangun dengan bahasa **Go (Golang)**. Proyek ini mendemonstrasikan implementasi **Clean Architecture**, dokumentasi otomatis dengan **Swagger**, serta integrasi **Cloud Database (Supabase)** dan **Cloud Hosting (Railway)**.

Tugas ini merupakan serangkain Program bootcamp Jago Golang - Ariaseta & Umam

---

## 🏗️ Arsitektur & Alur Sistem

Aplikasi ini menggunakan prinsip **Clean Architecture** untuk memisahkan logika kode:
1. **Handler**: Mengelola HTTP request & response.
2. **Service**: Berisi logika bisnis (Business Logic).
3. **Repository**: Berkomunikasi langsung dengan Database (SQL).



---

## 🛠️ Langkah 1: Persiapan Alat (Prasyarat)

Pastikan alat berikut sudah terinstal di komputer Anda:
* **Go (v1.23+)**: [Download di sini](https://go.dev/dl/).
* **Git**: Untuk manajemen kode.
* **Swaggo CLI**: Untuk dokumentasi API otomatis.
  ```bash
  go install [github.com/swaggo/swag/cmd/swag@latest](https://github.com/swaggo/swag/cmd/swag@latest)

---

## 🗄️ Langkah 2: Setup Database (Supabase)

Jalankan script SQL berikut di SQL Editor Supabase Anda untuk membuat tabel dengan relasi:
SQL
* **Tabel Kategori**
    ```bash
    CREATE TABLE categories (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL
    );
    ```

* **Tabel Produk dengan Foreign Key**
```bash
    CREATE TABLE products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        price NUMERIC(15, 2) NOT NULL,
        stock INTEGER NOT NULL,
        category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL
    );
```

* **Masukkan Data Contoh**
```bash
    INSERT INTO categories (name) VALUES ('Elektronik'), ('Makanan');
```

---

## 🚀 Langkah 3: Menjalankan di Lokal

* **Clone & Install:**
```bash
    git clone [https://github.com/goprex/crudAPI.git](https://github.com/goprex/crudAPI.git)
    cd crudAPI
    go mod tidy
```

* **Setup Environment**
Buat file .env di root folder dan isi dengan link URI Supabase Anda:
```bash
    PORT=8080
    DB_CONN=postgres://postgres:[PASSWORD]@db.supabase.co:5432/postgres
```

* **Generate Swagger & Run:**
```bash
    swag init
    go run main.go
```

---

## 📖 Langkah 4: Testing API (CURL)

Gunakan perintah ini di terminal (Arch/Linux/Mac) untuk menguji API:
* **Tambah Produk Baru (POST)**
```bash
curl -X POST http://localhost:8080/api/produk \
     -H "Content-Type: application/json" \
     -d '{
           "name": "Mouse Gaming",
           "price": 250000,
           "stock": 10,
           "category_id": 1
         }'
```

* **Update Produk (PUT)**
```bash
curl -X PUT http://localhost:8080/api/produk/1 \
     -H "Content-Type: application/json" \
     -d '{
           "name": "Mouse Wireless",
           "price": 300000,
           "stock": 5,
           "category_id": 1
         }'
```

* **Hapus Produk (DELETE)**
```bash
curl -X DELETE http://localhost:8080/api/produk/1
```

---

## 📂 Struktur Folder

    handlers/ - Kontroler HTTP.

    services/ - Logika bisnis inti.

    repositories/ - Akses database.

    models/ - Definisi Struct data.

    docs/ - File Swagger (Auto-generated).


---

### Cara Update ke GitHub:
1. Simpan file `README.md` dengan isi di atas.
2. Jalankan perintah ini di terminal:
   ```bash
   git add README.md
   git commit -m "docs: fix formatting in README"
   git push origin main
    ```
