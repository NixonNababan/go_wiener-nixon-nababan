# (16) Join - Union - Agregasi - Subquery - Function (DBMS)
[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

### Join
Join adalah operasi penggabungan tabel dalam database yang digunakan untuk mengkombinasikan dua atau lebih tabel berdasarkan kolom yang sama atau memiliki keterkaitan.

Contoh:

```
SELECT orders.order_id, customers.customer_name
FROM orders
JOIN customers ON orders.customer_id = customers.customer_id;
```

### Union
Union adalah operasi penggabungan hasil dari dua atau lebih pernyataan SELECT. Hasilnya adalah gabungan baris-baris dari dua kueri yang berbeda.

Contoh:

```
SELECT customer_id, order_date FROM orders
UNION
SELECT customer_id, payment_date FROM payments;
```

### Agregasi
Agregasi adalah proses pengelompokan dan perhitungan pada data dalam database. Contoh fungsi agregasi adalah SUM, AVG, COUNT, MAX, dan MIN.

Contoh:

```
SELECT COUNT(order_id) AS total_order FROM orders;
```

### Subquery
Subquery adalah pernyataan SELECT yang diletakkan di dalam pernyataan SELECT, INSERT, UPDATE, atau DELETE lainnya. Subquery digunakan untuk mendapatkan data yang dibutuhkan dalam kondisi tertentu.

Contoh:
```
SELECT customer_name
FROM customers
WHERE customer_id IN (
    SELECT customer_id
    FROM orders
    WHERE order_date = '2022-01-01'
);
```

### Function
Fungsi dalam DBMS adalah prosedur yang didefinisikan sebelumnya dan dapat dipanggil kembali di kemudian hari. Fungsi dapat mengembalikan nilai atau hanya menjalankan tugas tertentu.

Contoh:
```
CREATE FUNCTION total_price(price INT, qty INT)
RETURNS INT
BEGIN
    DECLARE total INT;
    SET total = price * qty;
    RETURN total;
END;
```
Itulah beberapa materi penting dalam DBMS seperti Join, Union, Agregasi, Subquery, dan Function. Diharapkan dapat membantu Anda memahami dasar-dasar pengolahan data dalam database.