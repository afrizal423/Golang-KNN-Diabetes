package main

import (
	"fmt"
	"fp_datmin_03_19_92/core"
	"fp_datmin_03_19_92/core/helper/plot"
)

func main() {
	dataset := core.Buka_file("diabetes.csv")

	// pisah data menjadi data train dan data test
	persentase_pembagian_data := 0.25 // 40 persen data train
	// cleaning := core.Hilangkan_0(dataset)

	dataclean := core.Hilangkan_0(dataset, dataset)
	testset, trainset := core.Test_train_data(dataclean, persentase_pembagian_data)

	// fmt.Println(core.Hilangkan_0(dataset))
	// fmt.Println("jumlah data clean", len(dataclean))
	fmt.Println("jumlah data", len(dataset))
	fmt.Println("test data", len(testset))
	fmt.Println("train data", len(trainset))
	// plot.Clean_data(dataclean)
	// fmt.Println(trainset)
	var manhattan []float64
	var minkowski []float64
	var braycurtis []float64
	var canberra []float64
	var euclidean []float64
	var jumlah []float64
	var L1Dist []float64
	var consine []float64
	fmt.Println("\nMohon ditunggu, sistem dalam perhitungan")
	for i := 1; i <= 50; i++ {
		// sum += i
		// fmt.Println("data ke ", i)
		// coba := core.KNearestNeighbor(i, trainset, testset, "braycurtis")
		manhattan = append(manhattan, core.KNearestNeighbor(i, trainset, testset, "manhattan"))
		minkowski = append(minkowski, core.KNearestNeighbor(i, trainset, testset, "minkowski"))
		braycurtis = append(braycurtis, core.KNearestNeighbor(i, trainset, testset, "braycurtis"))
		canberra = append(canberra, core.KNearestNeighbor(i, trainset, testset, "canberra"))
		euclidean = append(euclidean, core.KNearestNeighbor(i, trainset, testset, "euclidean"))
		L1Dist = append(L1Dist, core.KNearestNeighbor(i, trainset, testset, "l1"))
		consine = append(consine, core.KNearestNeighbor(i, trainset, testset, "consine"))
		jumlah = append(jumlah, float64(i))
		// fmt.Println(coba)
	}

	plot.Hasil(manhattan, minkowski, braycurtis, canberra, euclidean, L1Dist, consine, jumlah)
	// fmt.Println(jumlah)
	// var k int = 6
	// core.KNearestNeighbor(k, trainset, testset)
}
