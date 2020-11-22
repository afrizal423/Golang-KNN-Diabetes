package core

import (
	"fp_datmin_03_19_92/core/structs"
	"math/rand"
)

func Test_train_data(data []structs.DiabetesRecord, persentase_data float64) ([]structs.DiabetesRecord, []structs.DiabetesRecord) {
	var testSet []structs.DiabetesRecord
	var trainSet []structs.DiabetesRecord

	for i := range data {
		// jika random integer kurang persentase_data,
		// maka akan memasukkan ke array trainSet
		if rand.Float64() < persentase_data {
			trainSet = append(trainSet, data[i])
		} else {
			// jika tidak maka akan masuk ke array testSet
			testSet = append(testSet, data[i])
		}
	}
	return testSet, trainSet
}
