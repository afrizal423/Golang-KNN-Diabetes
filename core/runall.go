package core

import (
	"fmt"
	"fp_datmin_03_19_92/core/helper/plot"
)

// RunAll ...
func RunAll() {
	var (
		k      int
		persen int = 0
	)
	dataset := Buka_file("diabetes.csv")

	// pisah data menjadi data train dan data test
	fmt.Println("\nPersentase data test: (ex: 40; yang sama dengan 0.40)")
	// persentasePembagianData, _, _ := reader.ReadRune()
	// persentasePembagianData, _ := reader.ReadString('\n')
	// persen := 0.25 // 40 persen data train
	// cleaning := core.Hilangkan_0(dataset)
	fmt.Scanf("%d", &persen)
	persentase := float64(float64(persen) / 100)

	dataclean := Hilangkan_0(dataset, dataset)
	testset, trainset := Test_train_data(dataclean, persentase)

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
	var (
		MEmanhattan  []float64
		MEminkowski  []float64
		MEbraycurtis []float64
		MEcanberra   []float64
		MEeuclidean  []float64
		MEL1Dist     []float64
		MEconsine    []float64
	)
	fmt.Println("\nBanyak percobaan berapa kali loop?: (ex: 3; berarti loop k 1 - 3)")
	// persentase_pembagian_data := 0.25 // 40 persen data train
	fmt.Scanf("%d", &k)
	fmt.Println("\nMohon ditunggu, sistem dalam perhitungan")

	fmt.Println("nilai start k=1 sampai k=", k)
	for i := 1; i <= int(k); i++ {
		// manhattan
		mnht, memnht, _ := KNearestNeighbor(i, trainset, testset, "manhattan")
		manhattan = append(manhattan, mnht)
		MEmanhattan = append(MEmanhattan, memnht)

		//minkowski
		mnkw, memnkw, _ := KNearestNeighbor(i, trainset, testset, "minkowski")
		minkowski = append(minkowski, mnkw)
		MEminkowski = append(MEminkowski, memnkw)

		//braycurtis
		bry, mebry, _ := KNearestNeighbor(i, trainset, testset, "braycurtis")
		braycurtis = append(braycurtis, bry)
		MEbraycurtis = append(MEbraycurtis, mebry)

		// canberra
		cnb, mecnb, _ := KNearestNeighbor(i, trainset, testset, "canberra")
		canberra = append(canberra, cnb)
		MEcanberra = append(MEcanberra, mecnb)

		// euclidean
		euc, meeuc, _ := KNearestNeighbor(i, trainset, testset, "euclidean")
		euclidean = append(euclidean, euc)
		MEeuclidean = append(MEeuclidean, meeuc)

		// L1
		l1, mel1, _ := KNearestNeighbor(i, trainset, testset, "l1")
		L1Dist = append(L1Dist, l1)
		MEL1Dist = append(MEL1Dist, mel1)

		//consine
		csn, mecsn, _ := KNearestNeighbor(i, trainset, testset, "consine")
		consine = append(consine, csn)
		MEconsine = append(MEconsine, mecsn)
		jumlah = append(jumlah, float64(i))
		// fmt.Println(coba)
	}

	plot.Hasil(manhattan, minkowski, braycurtis, canberra, euclidean, L1Dist, consine, jumlah)
	plot.MeanError(MEmanhattan, MEminkowski, MEbraycurtis, MEcanberra, MEeuclidean, MEL1Dist, MEconsine, jumlah)
	// fmt.Println(jumlah)
	// var k int = 6
	// core.KNearestNeighbor(k, trainset, testset)
}
