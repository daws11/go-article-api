# Microservice API Artikel (Go)

Repository Microservice API Artikel ini dibuat oleh Muhammad Abdurrahman Firdaus sebagai kandidat Golang Developer di PT Sharing Vision Indonesia.

Repository ini berisikan kode untuk sebuah microservice API CRUD (Create, Read, Update, Delete) sederhana untuk mengelola artikel, dibangun menggunakan Go, Gin, dan MySQL. Proyek ini merupakan contoh implementasi berdasarkan kebutuhan dari dokumen "Test Backend Sharing Vision 2023".

## Technology Stack

* **Bahasa:** Go (Golang)
* **Framework Web:** Gin (`github.com/gin-gonic/gin`)
* **Database:** MySQL
* **Driver Database:** `github.com/go-sql-driver/mysql`
* **Migrasi Database:** `github.com/golang-migrate/migrate`
* **Validasi:** `github.com/go-playground/validator/v10`
* **Variabel Lingkungan:** `github.com/joho/godotenv` 

## Fitur

* Operasi CRUD (Buat, Baca, Ubah, Hapus) untuk artikel.
* Endpoint terpaginasi (paginated) untuk mengambil daftar artikel.
* Validasi input saat membuat dan mengubah artikel.
* Manajemen skema database menggunakan migrasi.

## Prasyarat

* Go (versi 1.18)
* Server MySQL
* Alat CLI `golang-migrate/migrate` (Instal: `go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`)
* Git

## Pengaturan & Instalasi

1.  **Clone repository:**
    ```bash
    git clone [https://github.com/daws11/go-article-api.git](https://github.com/daws11/go-article-api.git) 
    cd go-article-api
    ```

2.  **Buat File Lingkungan (.env):**
    Salin file contoh lingkungan dan perbarui dengan kredensial database Anda:
    ```bash
    cp .env.example .env
    ```
    Edit file `.env` dan isi `DB_USER`, `DB_PASSWORD`, `DB_HOST`, `DB_PORT`, dan `DB_NAME` Anda. Pastikan database yang ditentukan dalam `DB_NAME` (misalnya, `article`) sudah ada di server MySQL Anda. Anda juga bisa mengatur `PORT` untuk aplikasi (standarnya 8080).

3.  **Instal Dependensi:**
    ```bash
    go mod tidy
    ```
    atau
    ```bash
    go get ./...
    ```

4.  **Jalankan Migrasi Database:**
    Pastikan server MySQL Anda berjalan dan database sudah ada dan sudah membut database dengan nama 'article'. Kemudian jalankan:
    ```bash
    # Ganti DSN dengan koneksi Anda jika berbeda dari .env
    # Format DSN: mysql://user:password@tcp(host:port)/dbname
    migrate -database 'mysql://root:secret@tcp(localhost:3306)/article' -path ./migrations up
    ```
    *(Ganti DSN pada perintah di atas dengan string koneksi database Anda yang sebenarnya jika perlu)*

5.  **Jalankan Aplikasi:**
    ```bash
    go run cmd/server/main.go
    ```
    Server akan dimulai, biasanya pada port 8080 (atau port yang ditentukan di `.env`).

## Variabel Lingkungan (Environment Variables)

Aplikasi menggunakan variabel lingkungan berikut (didefinisikan dalam `.env` untuk pengembangan lokal):

| Variabel      | Deskripsi                           | Standar     |
| :------------ | :---------------------------------- | :---------- |
| `DB_USER`     | Nama pengguna database MySQL        | `root`      |
| `DB_PASSWORD` | Kata sandi database MySQL           | `secret`    |
| `DB_HOST`     | Host database MySQL                 | `localhost` |
| `DB_PORT`     | Port database MySQL                 | `3306`      |
| `DB_NAME`     | Nama database MySQL                 | `article`   |
| `PORT`        | Port tempat server API akan berjalan | `8080`      |

## Dokumentasi API

**URL Dasar (Base URL):** `http://localhost:8080` (atau port yang dikonfigurasi)

---

### 1. Buat Artikel (Create Article)

* **Metode:** `POST`
* **Path:** `/article/`
* **Deskripsi:** Membuat artikel baru.
* **Request Body:** (`application/json`)
    ```json
    {
      "title": "Judul Valid",
      "content": "Konten artikel harus cukup panjang. Field ini menyimpan isi utama artikel. Pastikan Anda memberikan teks yang cukup di sini jika tidak, API akan mengembalikan kesalahan validasi berdasarkan persyaratan panjang minimum.",
      "category": "Teknologi",
      "status": "Draft"
    }
    ```
* **Aturan Validasi:** Lihat bagian [Aturan Validasi](#aturan-validasi) di bawah.
* **Respons Sukses (Success Response):**
    * **Kode:** `201 Created`
    * **Body:** (`application/json`) Mengembalikan objek artikel yang baru dibuat.
        ```json
        {
          "id": 1,
          "title": "Judul Valid Minimal 20 Karakter",
          "content": "Konten artikel harus cukup panjang, setidaknya dua ratus karakter agar lolos validasi...",
          "category": "Teknologi",
          "created_date": "2025-04-11T18:00:00Z", // Contoh timestamp
          "updated_date": "2025-04-11T18:00:00Z", // Contoh timestamp
          "status": "Draft"
        }
        ```
* **Respons Gagal (Error Responses):**
    * **Kode:** `400 Bad Request` (JSON tidak valid atau Kesalahan Validasi)
        ```json
        {
          "error": "Invalid request body: json: cannot unmarshal string into Go struct field ArticleInput.title of type int" // Contoh error JSON
        }
        ```
        ```json
        {
          "error": "Failed to create article: validation error: Key: 'ArticleInput.Content' Error:Field validation for 'Content' failed on the 'min' tag" // Contoh error Validasi
        }
        ```
    * **Kode:** `500 Internal Server Error` (Kesalahan database, dll.)
        ```json
        {
          "error": "Failed to create article: error creating article: ..." // Contoh error server
        }
        ```

---

### 2. Dapatkan Daftar Artikel (Get Articles - Paginated)

* **Metode:** `GET`
* **Path:** `/article`
* **Deskripsi:** Mengambil daftar artikel dengan paginasi.
* **Parameter Kueri (Query Parameters):**
    * `limit` (integer, opsional, standar: 10): Jumlah maksimum artikel yang akan dikembalikan.
    * `offset` (integer, opsional, standar: 0): Jumlah artikel yang akan dilewati dari awal.
* **Contoh Request:** `GET http://localhost:8080/article?limit=5&offset=10`
* **Respons Sukses:**
    * **Kode:** `200 OK`
    * **Body:** (`application/json`) Mengembalikan sebuah array (daftar) objek artikel.
        ```json
        [
          {
            "id": 11,
            "title": "Judul Artikel 11",
            "content": "Konten untuk artikel 11...",
            "category": "Berita",
            "created_date": "2025-04-10T10:00:00Z",
            "updated_date": "2025-04-10T11:00:00Z",
            "status": "Publish"
          },
          {
            "id": 12,
            "title": "Judul Artikel 12",
            "content": "Konten untuk artikel 12...",
            "category": "Olahraga",
            "created_date": "2025-04-10T12:00:00Z",
            "updated_date": "2025-04-10T12:00:00Z",
            "status": "Draft"
          }
          
        ]
        ```
        *Mengembalikan array kosong `[]` jika tidak ada artikel yang cocok dengan offset/limit.*
* **Respons Gagal:**
    * **Kode:** `500 Internal Server Error` (Kesalahan database, dll.)
        ```json
        {
          "error": "Failed to retrieve articles: error querying articles: ..."
        }
        ```

---

### 3. Dapatkan Artikel berdasarkan ID (Get Article by ID)

* **Metode:** `GET`
* **Path:** `/article/:id`
* **Deskripsi:** Mengambil artikel spesifik berdasarkan ID-nya.
* **Parameter Path (Path Parameters):**
    * `id` (integer, wajib): ID artikel yang ingin diambil.
* **Contoh Request:** `GET http://localhost:8080/article/5`
* **Respons Sukses:**
    * **Kode:** `200 OK`
    * **Body:** (`application/json`) Mengembalikan objek artikel.
        ```json
        {
          "id": 5,
          "title": "Judul Artikel Spesifik",
          "content": "Konten untuk artikel 5...",
          "category": "Gaya Hidup",
          "created_date": "2025-04-09T15:30:00Z",
          "updated_date": "2025-04-09T16:00:00Z",
          "status": "Publish"
        }
        ```
* **Respons Gagal:**
    * **Kode:** `400 Bad Request` (Format ID tidak valid)
        ```json
        {
          "error": "Invalid article ID format"
        }
        ```
    * **Kode:** `404 Not Found` (Artikel dengan ID yang diberikan tidak ada)
        ```json
        {
          "error": "Article with ID 999 not found"
        }
        ```
    * **Kode:** `500 Internal Server Error` (Kesalahan database, dll.)
        ```json
        {
          "error": "Failed to retrieve article: ..."
        }
        ```

---

### 4. Ubah Artikel (Update Article)

* **Metode:** `PUT` atau `PATCH` (Implementasi saat ini menggunakan `PUT`)
* **Path:** `/article/:id`
* **Deskripsi:** Memperbarui artikel yang sudah ada berdasarkan ID-nya. `PUT` biasanya menggantikan seluruh sumber daya.
* **Parameter Path:**
    * `id` (integer, wajib): ID artikel yang ingin diperbarui.
* **Request Body:** (`application/json`) - Struktur dan aturan validasi sama seperti Buat Artikel.
    ```json
    {
      "title": "Judul Valid Terupdate (Min 20 Karakter)",
      "content": "Konten terupdate yang juga harus memiliki panjang minimal dua ratus karakter. Ini memastikan konsistensi selama pembaruan. Ingat bahwa PUT biasanya menyiratkan penggantian seluruh sumber daya, jadi semua field yang dapat divalidasi harus disediakan sesuai aturan.",
      "category": "Kategori Terupdate",
      "status": "Publish"
    }
    ```
* **Aturan Validasi:** Lihat bagian [Aturan Validasi](#aturan-validasi) di bawah.
* **Respons Sukses:**
    * **Kode:** `200 OK`
    * **Body:** (`application/json`) Mengembalikan objek artikel yang sudah diperbarui.
        ```json
        {
          "id": 5,
          "title": "Judul Valid Terupdate (Min 20 Karakter)",
          "content": "Konten terupdate yang juga harus memiliki panjang minimal dua ratus karakter...",
          "category": "Kategori Terupdate",
          "created_date": "2025-04-09T15:30:00Z", // Tanggal dibuat biasanya tidak berubah
          "updated_date": "2025-04-11T18:30:00Z", // Timestamp terupdate
          "status": "Publish"
        }
        ```
* **Respons Gagal:**
    * **Kode:** `400 Bad Request` (Format ID tidak valid, JSON tidak valid, atau Kesalahan Validasi)
    * **Kode:** `404 Not Found` (Artikel dengan ID yang diberikan tidak ada)
    * **Kode:** `500 Internal Server Error` (Kesalahan database, dll.)

---

### 5. Hapus Artikel (Delete Article)

* **Metode:** `DELETE`
* **Path:** `/article/:id`
* **Deskripsi:** Menghapus artikel berdasarkan ID-nya.
* **Parameter Path:**
    * `id` (integer, wajib): ID artikel yang ingin dihapus.
* **Contoh Request:** `DELETE http://localhost:8080/article/5`
* **Respons Sukses:**
    * **Kode:** `200 OK`
    * **Body:** (`application/json`)
        ```json
        {
          "message": "Article with ID 5 deleted successfully"
        }
        ```
    *(Alternatifnya, respons `204 No Content` dengan body kosong juga umum digunakan)*
* **Respons Gagal:**
    * **Kode:** `400 Bad Request` (Format ID tidak valid)
    * **Kode:** `404 Not Found` (Artikel dengan ID yang diberikan tidak ada)
    * **Kode:** `500 Internal Server Error` (Kesalahan database, dll.)

---

## Aturan Validasi

Aturan validasi berikut berlaku saat membuat (`POST /article/`) atau memperbarui (`PUT /article/:id`) artikel:

* **`title`**: Wajib diisi (Required), 
* **`content`**: Wajib diisi (Required),
* **`category`**: Wajib diisi (Required), Panjang minimal: 3 karakter.
* **`status`**: Wajib diisi (Required), Harus salah satu dari: `"Publish"`, `"Draft"`, `"Thrash"`.

## Cara Pengujian (How to Test)

Anda dapat menggunakan alat seperti [Postman](https://www.postman.com/), Insomnia, atau `curl` untuk berinteraksi dengan endpoint API.

1.  Pastikan aplikasi berjalan (`go run cmd/server/main.go`).
2.  Kirim permintaan ke endpoint yang sesuai yang tercantum di bagian [Dokumentasi API](#dokumentasi-api) menggunakan URL dasar `http://localhost:8080` (atau port yang Anda konfigurasikan).
3.  Ingatlah untuk mengatur header `Content-Type` ke `application/json` untuk permintaan dengan body (POST, PUT).
