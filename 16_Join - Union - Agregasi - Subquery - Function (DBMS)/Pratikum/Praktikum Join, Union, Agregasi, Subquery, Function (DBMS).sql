--INSERT--


-- Insert 5 operators pada table operators.
insert into operator values
('1','GUCCI', now(), now()),
('2','RICK OWENS', now(), now()),
('3','SUPREME', now(), now()),
('4','UNIQLO', now(), now()),
('5','PULL&BEAR', now(), now());

-- Insert 3 product type.
insert into product_types values 
('1','Baju', now(), now()),
('2','Celana', now(), now()),
('3','Outwear', now(), now());

-- Insert 2 product dengan product type id = 1, dan operators id = 3.
insert into product(product_type_id, operator_id,code,nama,status,create_at,update_at) values
('1','3','MM01','Kemeja SLim-fit','1',now(),now()),
('1','3','MM02','Kemeja Oversized','1',now(),now());

-- Insert 3 product dengan product type id = 2, dan operators id = 1.
insert into product(product_type_id, operator_id,code,nama,status,create_at,update_at) value
('2','1','MM03','Celana Jeans','1',now(),now()),
('2','1','MM04','Celana Pendek','1',now(),now()),
('2','1','MM05','Celana Cargo','1',now(),now());

-- Insert 3 product dengan product type id = 3, dan operators id = 4.
insert into product(product_type_id, operator_id,code,nama,status,create_at,update_at) value
('3','4','MM06','Hoodie','1',now(),now()),
('3','4','MM07','Jaket Kulit','1',now(),now()),
('3','4','MM08','Vest','1',now(),now());

insert into product_description (description, product_id,create_at,update_at) values 
('Kemeja slim-fit yang memperlihatkan lekukan badan anda','1',now(),now()),
('Kemeja overzied, kemeja kekinian yang sangat trendy','2',now(),now()),
('Celana jeans anti kucel','3',now(),now()),
('Celana pendek untuk bersantai','4',now(),now()),
('Celana cargo membuat anda terlihat lebih wadaw','5',now(),now()),
('Hoodie melindungi anda dari debt-collector','6',now(),now()),
('Jaket kulit membuat anda terlihat seperti bapak','7',now(),now()),
('Vest tukang parkir','8',now(),now());

-- Insert 3 payment methods.
insert into payment_methods (nama, status,create_at,update_at) values
('Mandiri','1',now(),now()),
('QRIS','2',now(),now()),
('CASH','3',now(),now());

insert into users (nama,status_user,dob,gender,create_at,update_at) values
('Mega','1','2002-03-29','W',now(),now()),
('Wati','1','2002-03-01','W',now(),now()),
('Ferdy','1','2002-03-02','M',now(),now()),
('Sambo','1','2002-03-06','M',now(),now()),
('Nixon','1','2002-03-19','M',now(),now());


insert into transaction (user_id,payments_method_id,status,total_qty,total_price,create_at,update_at) values
('1','1','Success','1','199000',now(),now()),
('2','2','Pending','2','599000',now(),now()),
('3','3','Failed','3','799000',now(),now()),
('1','1','Success','2','199000',now(),now()),
('1','1','Success','3','199000',now(),now()),
('2','2','Pending','5','399000',now(),now()),
('2','3','Failed','1','199000',now(),now()),
('3','1','Success','1','999000',now(),now()),
('3','2','Cancelled','1','699000',now(),now()),
('1','1','Success','1','199000',now(),now()),
('2','2','Pending','2','599000',now(),now()),
('4','3','Failed','3','799000',now(),now()),
('4','1','Succsess','3','799000',now(),now()),
('4','2','Pending','3','799000',now(),now()),
('5','3','Failed','3','799000',now(),now()),
('5','2','Pending','3','799000',now(),now()),
('5','3','Failed','3','799000',now(),now());


insert into transaction_details (transaction_id,product_id,status,qty,price,create_at,update_at) values
('1','1','OK','1','199000',now(),now()),
('1','2','OK','2','199000',now(),now()),
('1','3','OK','3','199000',now(),now()),
('2','4','OK','1','199000',now(),now()),
('2','5','OK','2','199000',now(),now()),
('2','6','OK','3','199000',now(),now()),
('3','7','OK','1','199000',now(),now()),
('3','8','OK','2','199000',now(),now()),
('3','1','OK','3','599000',now(),now()),
('4','2','OK','1','599000',now(),now()),
('4','3','OK','2','599000',now(),now()),
('4','4','OK','3','599000',now(),now()),
('5','5','OK','1','599000',now(),now()),
('5','6','OK','1','599000',now(),now()),
('5','7','OK','1','599000',now(),now()),
('6','8','OK','1','20000',now(),now()),
('6','1','OK','2','199000',now(),now()),
('6','2','OK','3','199000',now(),now()),
('7','3','OK','1','199000',now(),now()),
('7','4','OK','2','199000',now(),now()),
('7','5','OK','3','599000',now(),now()),
('8','6','OK','1','599000',now(),now()),
('8','7','OK','2','599000',now(),now()),
('8','8','OK','3','599000',now(),now()),
('9','1','OK','1','599000',now(),now()),
('9','2','OK','2','299000',now(),now()),
('9','3','OK','3','299000',now(),now()),
('10','4','OK','1','999000',now(),now()),
('10','5','OK','2','999000',now(),now()),
('10','6','OK','3','999000',now(),now()),
('11','7','OK','1','999000',now(),now()),
('11','8','OK','2','999000',now(),now()),
('11','1','OK','3','699000',now(),now()),
('12','2','OK','1','699000',now(),now()),
('12','3','OK','2','699000',now(),now()),
('12','4','OK','3','699000',now(),now()),
('13','5','OK','1','699000',now(),now()),
('13','6','OK','2','699000',now(),now()),
('13','7','OK','3','599000',now(),now()),
('14','8','OK','1','599000',now(),now()),
('14','1','OK','2','599000',now(),now()),
('14','2','OK','3','599000',now(),now()),
('15','3','OK','1','599000',now(),now()),
('15','4','OK','2','899000',now(),now()),
('15','5','OK','3','8990000',now(),now());

-- Tampilkan nama user / pelanggan dengan gender Laki-laki / M.--
select * from users where gender='M';

-- Tampilkan product dengan id = 3. --
select * from product where id='3';

-- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’. --
select * from users where create_at > date_sub(now(), interval 1 week) AND nama LIKE '%a%';

-- Hitung jumlah user / pelanggan dengan status gender Perempuan. --
 select sum(gender='W') from users;
 
 -- Tampilkan data pelanggan dengan urutan sesuai nama abjad --
 select * from users order by nama ASC;

-- Tampilkan 5 data pada data product -- 
select * from product limit 5;

-- Update data product id 1 dengan nama ‘product dummy’. --
update product set nama='product dummy' where id='1'; 

-- Update qty = 3 pada transaction detail dengan product id = 1. --
update transaction_details set qty='3' where product_id='1'; 

-- Delete data pada tabel product dengan id = 1. --
DELETE FROM product WHERE id = 1;

-- Delete pada pada tabel product dengan product type id 1. -- 
DELETE FROM product WHERE product_type_id = 1;


----------------------------------------
-- Join, Union, Sub query, Function--
----------------------------------------

-- Gabungkan data transaksi dari user id 1 dan user id 2. --
SELECT * FROM transaction WHERE user_id = 1 OR user_id = 2;
 
-- Tampilkan jumlah harga transaksi user id 1.
SELECT SUM(total_price) FROM transaction WHERE user_id = 1;

-- Tampilkan semua field table product dan field name table product type yang saling berhubungan. -- 
select count(*) from transaction t
join transaction_details td on t.id = td.transaction_id
join product p on p.id = td.product_id
where p.product_type_id = 2;

-- Tampilkan semua field table transaction, field name table product dan field name table user. --
select t.*, p.nama as product, u.nama as users from transaction_details td
join transaction t on t.id = td.transaction_id
join users u on u.id = t.users_id
join product p on p.id = td.product_id
where p.deleted_at is null


 -- Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud. --
 delimiter $$
 create trigger delete_transaction_details
 before delete on transaction for each row
 begin
 declare trans_id int;
 set trans_id = old.id;
 delete from transaction_details where transaction_id = trans_id;
 end $$
 
 delete from transaction where id = 4;
 
 -- Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus. --
 delimiter $$
 create trigger update_transaction_id
 after delete on transaction_details for each row
 begin
 declare trans_id int;
 set trans_id = old.transaction_id;
 update transaction set total_quanty = (select sum(quanty) from transaction_details where transaction_id = trans_id)
 where id = trans_id;
 end $$
 
 select * from transaction;
 select * from transaction_details where transaction_id = 15;
 delete from transaction_details where transaction_id = 15 AND product_id = 4;
 
 -- Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query. -- 
 select * from product where id not in(
 select product_id from transaction_details
 );
 
select * from product where id in (
select p.id from product p
left outer join transaction_details td on p.id = td.product_id
where td.transaction_id is null
);


