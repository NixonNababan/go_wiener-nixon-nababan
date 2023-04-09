# (22) Middleware
[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

Middleware dalam bahasa Go adalah sebuah konsep yang digunakan untuk menambahkan fungsionalitas tambahan pada aplikasi web. Middleware adalah lapisan perantara antara permintaan dan respons dalam aplikasi web.

Middleware dapat digunakan untuk melakukan berbagai tugas seperti otentikasi, otorisasi, penanganan kesalahan, logging, caching, dan banyak lagi. Dalam Go, middleware biasanya dibuat sebagai sebuah fungsi yang mengambil parameter http.Handler dan mengembalikan http.Handler.

Fungsi middleware dapat digunakan dengan memanggil http.HandlerFunc dan menjalankan fungsinya di dalam http.ListenAndServe atau http.ServeHTTP. Middleware juga dapat digunakan dengan menggunakan framework web Go seperti Gorilla Mux atau Gin.