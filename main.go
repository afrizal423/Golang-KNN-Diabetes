package main

import (
	"fmt"
	"fp_datmin_03_19_92/core"
)

func main() {
	dataset := core.Buka_file("diabetes.csv")

	// pisah data menjadi data train dan data test
	persentase_pembagian_data := 0.4 // 40 persen data train
	testset, trainset := core.Test_train_data(dataset, persentase_pembagian_data)

	fmt.Println("jumlah data", len(dataset))
	fmt.Println("test data", len(testset))
	fmt.Println("train data", len(trainset))
}
