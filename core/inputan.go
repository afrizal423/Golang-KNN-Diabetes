package core

import (
	"fmt"
	"fp_datmin_03_19_92/core/structs"
)

// InputanUserProses ...
func InputanUserProses(dataTest []structs.DiabetesRecord) {
	var (
		k int
	)
	dataset := Buka_file("diabetes.csv")

	// pisah data menjadi data train dan data test

	dataclean := Hilangkan_0(dataset, dataset)
	fmt.Println("\nTentukan nilai K: (ex: 3)")
	// persentase_pembagian_data := 0.25 // 40 persen data train
	fmt.Scanf("%d", &k)
	// fmt.Println("\nMohon ditunggu, sistem dalam perhitungan")
	// fmt.Println("nilai k", k)
	// var (
	// 	prediksi [...]
	// 	// euclidean   []float64
	// )
	// for i := 1; i <= int(k); i++ {
	// euclidean
	_, _, hasil := KNearestNeighbor(k, dataclean, dataTest, "euclidean")
	// prediksi = append(prediksi, hasil)
	// MEeuclidean = append(MEeuclidean, meeuc)
	// }
	// fmt.Println(hasil)
	if hasil[0].prediksi == 1 {
		fmt.Println("\nHasil Prediksi sesuai inputan maka pasien dinyatakan POSITIF mengalami penyakit Diabetes")
	} else {
		fmt.Println("\nHasil Prediksi sesuai inputan maka pasien dinyatakan NEGATIF dari penyakit Diabetes")
	}
	// fmt.Println("Predicted: ", hasil[0].prediksi)

}
