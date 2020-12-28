<p align="right">
بِسْــــــــــــــمِ اللَّهِ الرَّحْمَنِ الرَّحِيم 
</p>

# K Nearest Neighbor using Go

Program sederhana untuk klasifikasi Diabetes menggunakan bahasa pemrograman Go. Memakai dataset dari [Kaggle](https://www.kaggle.com/johndasilva/diabetes). Preprocessing data yang dilakukan yaitu dengan normalisasi atribut class Glucose, BloodPressure, SkinThickness, BMI, dan Insulin yang semula nilai 0 diganti menjadi nilai mean dari kolom masing-masing. 
Adapun nilai K yang cocok untuk klasifikasi menggunakan program ini adalah K=3 dengan datatest=40% datatrain=60%, maka didapatkan hasil sebagai berikut:
- AKURASI =  0.9244823386114495
- MEAN ERROR =  0.07551766138855055

# How To Install
- Download Go terlebih dahulu [disini](https://golang.org/dl/)
- Download repo [ini](https://github.com/afrizal423/Golang-KNN-Diabetes/archive/master.zip)
- Masuk kedalam foldernya
- Jalankan perintah ```./install.sh``` untuk membuat folder output plot
- Jalankan perintah ```go run main.go``` untuk menjalankan program

# Author
- Afrizal Muhammad Yasin
- Dina Puspitaningrum
- Nurlisa Aulia Setyaningrum