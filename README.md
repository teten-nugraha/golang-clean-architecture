## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)

## General info
Repository adalah hasil percobaan penulis dalam belajar Golang yang mana menggunakan Clean Architecture, yang memisahkan setiap layer aplikasi. Pada umumnya dibedakan menjadi layer :

1. Repository
2. Service
3. Handler

![clean](images/clean.jpg "clean")

## Technologies
Project ini menggunakan beberapa teknologi yaitu :
* Golang Versi 1.15
* Echo Labstack 
* GORM
* Godotenv
* Logrush
* MySQL
## Setup
Sebelum itu pastikan MySQL sudah running dan buat sebuah database dengan nama **golang_crud** dan sesuaikan dengan yang ada di **file .env**

Berikut langkah-langkah untuk running repo ini

```
$ git clone https://github.com/teten-nugraha/golang-crud.git
$ go mod download
$ go run main.go
```