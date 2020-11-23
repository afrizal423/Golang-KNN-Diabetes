package main

import (
	"fmt"
	"fp_datmin_03_19_92/core"
)

func main() {
	dataset := core.Buka_file("diabetes.csv")

	// pisah data menjadi data train dan data test
	persentase_pembagian_data := 0.7 // 40 persen data train
	// cleaning := core.Hilangkan_0(dataset)

	dataclean := core.Hilangkan_0(dataset, dataset)
	testset, trainset := core.Test_train_data(dataclean, persentase_pembagian_data)
	// fmt.Println(core.Hilangkan_0(dataset))
	fmt.Println("jumlah data clean", len(dataclean))
	fmt.Println("jumlah data", len(dataset))
	fmt.Println("test data", len(testset))
	fmt.Println("train data", len(trainset))
	// fmt.Println(trainset)
}
