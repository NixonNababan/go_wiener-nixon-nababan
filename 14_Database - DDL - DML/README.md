# (14)Database Schema, DDL, dan DML
[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

## Database Schema, DDL, dan DML

### Database Schema
Database schema adalah struktur logika dari suatu database yang berisi definisi objek-objek dalam database seperti tabel, kolom, kunci primer, kunci asing, dan lain sebagainya. Database schema memberikan panduan untuk mengakses data pada database.

### Data Definition Language (DDL)
Data Definition Language (DDL) adalah bahasa yang digunakan untuk membuat, mengubah, dan menghapus struktur database. Beberapa perintah DDL yang sering digunakan antara lain:

CREATE: digunakan untuk membuat objek database baru seperti tabel, view, dan indeks.
ALTER: digunakan untuk mengubah struktur objek database yang telah ada.
DROP: digunakan untuk menghapus objek database yang tidak dibutuhkan lagi.
Berikut adalah contoh kode untuk membuat tabel mahasiswa menggunakan perintah DDL:

```
  CREATE TABLE mahasiswa (
  id INT PRIMARY KEY,
  nama VARCHAR(50),
  jurusan VARCHAR(50),
  angkatan INT
);
```
Kode di atas akan membuat tabel mahasiswa dengan empat kolom yaitu id, nama, jurusan, dan angkatan.

### Data Manipulation Language (DML)
Data Manipulation Language (DML) adalah bahasa yang digunakan untuk memanipulasi data pada database. Beberapa perintah DML yang sering digunakan antara lain:

INSERT: digunakan untuk memasukkan data ke dalam tabel.
SELECT: digunakan untuk mengambil data dari tabel.
UPDATE: digunakan untuk mengubah data pada tabel.
DELETE: digunakan untuk menghapus data dari tabel.
Berikut adalah contoh kode untuk memasukkan data ke dalam tabel mahasiswa menggunakan perintah DML:

```
INSERT INTO mahasiswa (id, nama, jurusan, angkatan) VALUES
(1, 'John Doe', 'Teknik Informatika', 2020),
(2, 'Jane Doe', 'Sistem Informasi', 2021),
(3, 'Foo Bar', 'Teknik Elektro', 2019);
```
Kode di atas akan memasukkan tiga baris data ke dalam tabel mahasiswa.

Demikianlah penjelasan singkat mengenai database schema, DDL, dan DML. Dengan memahami konsep-konsep ini, Anda dapat membuat dan mengelola database dengan lebih efektif dan efisien.