package main

import (
	"bufio"
	"fmt"
	"fp_datmin_03_19_92/core"
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
			fmt.Print("Coming Soon")
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
