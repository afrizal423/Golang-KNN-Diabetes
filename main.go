package main

import (
	"bufio"
	"fmt"
	"fp_datmin_03_19_92/core"
	"fp_datmin_03_19_92/core/helper"
	"fp_datmin_03_19_92/core/structs"
	"math"
	"os"
)

func main() {
	// core.RunAll()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("========================================")
	fmt.Println("Final Project Data Mining 2020")
	fmt.Println("========================================")
	fmt.Println("Judul: Deteksi Diabetes Menggunakan Metode KNearestNeighbor")
	fmt.Println("By:")
	fmt.Println("Dina Puspitaningrum (17081010003)\nNurlisa Aulia Seetyaningrum (17081010019)\nAfrizal Muhammad Yasin (17081010092)")
	fmt.Println("========================================")
	status := true

	for status {
		fmt.Println("\nPilihan Menu")
		fmt.Println("1. Klasifikasi dengan inputan user")
		fmt.Println("2. Lihat Akurasi dataset")
		fmt.Println("3. Keluar Program")
		fmt.Println("\nMasukkan Pilihan: ")
		pil, _, _ := reader.ReadRune()

		// fmt.Print("\033[H\033[2J")
		if pil == '1' {
			fmt.Print("\033[H\033[2J")
			reader := bufio.NewReader(os.Stdin)
			// fmt.Println(inputuser(reader))
			core.InputanUserProses(inputuser(reader))
		}
		if pil == '2' {
			fmt.Print("\033[H\033[2J")
			core.RunAll()
		}
		if pil == '3' {
			status = stop()
		}
	}

}

func stop() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\napakah ingin menghentikan aplikasi?")
	stop, _, _ := reader.ReadRune()
	var status bool
	if stop == 'y' || stop == '1' {
		status = false
	}
	return status
}

func inputuser(reader *bufio.Reader) []structs.DiabetesRecord {
	atribut := make([]structs.DiabetesRecord, 0)
	for i := 0; i < 1; i++ {
		var (
			berat     float64
			tinggi    float64
			perkiraan string
		)
		tmp := structs.DiabetesRecord{}
		fmt.Println("Jenis Kelamin Pasien? (L/P)")
		jk, _, _ := reader.ReadRune()
		if jk == 'p' || jk == 'P' {
			fmt.Println("Berapakah kali Pasien hamil?")
			fmt.Fscan(reader, &tmp.Pregnancies)
		}
		fmt.Println("Berapakah umur Pasien saat ini?")
		fmt.Fscan(reader, &tmp.Age)
		fmt.Println("Berapakah Konsentrasi glukosa plasma 2 jam saat ini pada Pasien?")
		fmt.Fscan(reader, &tmp.Glucose)
		fmt.Println("Berapakah Tekanan darah diastolik (mmHg) saat ini pada Pasien?")
		fmt.Fscan(reader, &tmp.BloodPressure)
		fmt.Println("Berapakah Ketebalan lipatan kulit trisep (mm) pada Pasien?")
		fmt.Fscan(reader, &tmp.SkinThickness)
		fmt.Println("Berapakah Insulin serum 2 jam (mu U / ml) pada Pasien?")
		fmt.Fscan(reader, &tmp.Insulin)

		fmt.Println("Berapakah Berat Badan pada Pasien (kg)?")
		fmt.Fscan(reader, &berat)
		fmt.Println("Berapakah Tinggi Badan pada Pasien? (cm)?")
		fmt.Fscan(reader, &tinggi)
		// fmt.Println(berat / math.Pow(float64(tinggi/100), 2))
		tmp.BMI = helper.ToFixed(berat/math.Pow(float64(tinggi/100), 2), 1)
		fmt.Println("Diabetes pedigree function pada Pasien?")
		fmt.Fscan(reader, &tmp.DiabetesPedigreeFunction)
		fmt.Println("Apakah pasien terlihat seperti terkena penyakit diabetes menurut perkiraan anda? (y/n)")
		fmt.Fscan(reader, &perkiraan)
		if perkiraan == "y" || perkiraan == "Y" {
			tmp.Outcome = 1
		}
		if perkiraan == "n" || perkiraan == "N" {
			tmp.Outcome = 0
		}
		atribut = append(atribut, tmp)
	}
	return atribut
}
