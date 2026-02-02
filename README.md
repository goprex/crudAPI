# 🛒 Kasir API - Clean Architecture Edition

API manajemen kasir modern yang dibangun dengan bahasa **Go (Golang)**. Proyek ini mendemonstrasikan implementasi **Clean Architecture**, dokumentasi otomatis dengan **Swagger**, serta integrasi **Cloud Database (Supabase)** dan **Cloud Hosting (Railway)**.

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
