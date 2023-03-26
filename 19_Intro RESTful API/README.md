# (19) Intro RESTful API
[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

### RESTful API 
## Apa itu RESTful API?
RESTful API merupakan sebuah standar arsitektur yang digunakan dalam pembangunan aplikasi web yang berbasis pada protokol HTTP. RESTful API menyediakan sebuah cara untuk mengakses sumber daya atau data melalui antarmuka yang bersifat stateless. Dalam arsitektur RESTful API, setiap data atau sumber daya diidentifikasi dengan sebuah URI (Uniform Resource Identifier), dan pengguna dapat mengakses data tersebut dengan menggunakan metode HTTP seperti GET, POST, PUT, dan DELETE.

## Komponen dari RESTful API
Ada beberapa komponen utama yang terdapat dalam arsitektur RESTful API, antara lain:

Resource: Setiap data atau sumber daya dalam aplikasi web diidentifikasi dengan sebuah URI. URI tersebut dapat digunakan oleh pengguna untuk mengakses data atau sumber daya yang diinginkan.

HTTP Methods: Metode HTTP digunakan untuk mengakses atau memanipulasi data atau sumber daya. Beberapa metode HTTP yang sering digunakan dalam RESTful API antara lain GET, POST, PUT, dan DELETE.

Representation: Representasi data atau sumber daya dapat berupa format yang berbeda-beda seperti JSON, XML, atau HTML.

## Keuntungan dari RESTful API
Beberapa keuntungan dari menggunakan RESTful API dalam pembangunan aplikasi web antara lain:

Scalability: RESTful API memungkinkan untuk memisahkan antara server dan klien secara jelas, sehingga memudahkan dalam melakukan scaling pada sistem.

Flexibility: RESTful API dapat digunakan dalam berbagai bahasa pemrograman dan dapat diakses dari berbagai platform.

Caching: RESTful API mendukung caching, sehingga memungkinkan untuk mengoptimalkan performa aplikasi web.

Simplicity: RESTful API menggunakan standar HTTP dan URI yang sederhana dan mudah dipahami oleh pengguna.

## Cara menggunakan RESTful API
Untuk menggunakan RESTful API, pengguna harus melakukan request ke server menggunakan metode HTTP yang sesuai dengan sumber daya atau data yang ingin diakses atau dimanipulasi. Setelah itu, server akan memberikan respon yang sesuai dengan request yang diberikan.

Berikut adalah contoh request yang dilakukan untuk mengakses data menggunakan metode HTTP GET:

```
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type User struct {
    ID   string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
}

func main() {
    resp, err := http.Get("http://localhost:8000/users")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    var users []User
    err = json.NewDecoder(resp.Body).Decode(&users)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(users)
}

```

