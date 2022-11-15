Pada task akhir VIX Full Stack Developer ini kalian diarahkan untuk membentuk API berdasarkan kasus yang telah diberikan. Pada kasus ini, kalian diinstruksikan untuk membuat API untuk mengupload dan menghapus gambar. API yang kalian bentuk adalah POST, GET, PUT, dan DELETE.



## Ketentuan API :
Pada bagian User Endpoint :

1. POST : /users/register, dan gunakan atribut berikut ini
2. ID (primary key, required)
3. Username (required)
4. Email (unique & required) 
5. Password (required & minlength 6)
6. Relasi dengan model Photo (Gunakan constraint cascade)
7. Created At (timestamp)
8. Updated At (timestamp)
9. GET: /users/login
10. Using email & password (required)
11. PUT : /users/:userId (Update User)
12. DELETE : /users/:userId (Delete User)

### Photos Endpoint

1. POST : /photos 
2. ID
3. Title
4. Caption
5. PhotoUrl
6. UserID
7. Relasi dengan model User
8. GET : /photos
9. PUT : /photoId
10. DELETE : /:photoId


### Requirement :
1. Authorization dapat menggunakan tool Go JWT (https://github.com/dgrijalva/jwt-go) 
2. Pastikan hanya user yang membuat foto yang dapat menghapus / mengubah foto


### Environment
Struktur dokumen / environment dari GoLang yang akan dibentuk kurang lebih sebagai berikut :

1. app :Menampung pembuatan struct dalam kasus ini menggunakan struct user untuk keperluan data dan authentication
2. controllers : Berisi antara logic database yaitu models dan query
3. database: Berisi konfigurasi database serta digunakan untuk menjalankan koneksi database dan migration
4. helpers : Berisi fungsi-fungsi yang dapat digunakan di setiap tempat dalam hal ini jwt, bcrypt, headerValue
5. middlewares :Berisi fungsi yang digunakan untuk proses otentikasi jwt yang digunakan untuk proteksi api
6. models : Berisi models yang digunakan untuk relasi database 
7. router : Berisi konfigurasi routing / endpoint yang akan digunakan untuk mengakses api
8. go mod : Yang digunakan untuk manajemen package / dependency berupa library


### Tools :
Tools yang digunakan : 

1. Gin Gonic Framework : https://github.com/gin-gonic/gin 
2. Gorm : https://gorm.io/index.html 
3. JWT Go : https://github.com/dgrijalva/jwt-go 
4. Go Validator : http://github.com/asaskevich/govalidator 


## Untuk database, menggunakan server SQL MySQL.
