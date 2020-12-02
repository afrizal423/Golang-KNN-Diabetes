package core

import "fmt"

func ConfusionMatrix(array []Hasil) {
	// fmt.Println(array)
	var (
		tp float64 = 0
		tn float64 = 0
		fp float64 = 0
		fn float64 = 0
	)
	for x := range array {
		if array[x].aktual == 1 && array[x].prediksi == 1 {
			tp += 1
		}
		if array[x].aktual == 0 && array[x].prediksi == 0 {
			tn += 1
		}
		if array[x].aktual == 0 && array[x].prediksi == 1 {
			fp += 1
		}
		if array[x].aktual == 1 && array[x].prediksi == 0 {
			fn += 1
		}
	}
	var (
		recall     = float64((tp) / (tp + fn))
		precission = float64((tp) / (tp + fp))
	)
	fmt.Println("hasil tp ", tp)
	fmt.Println("hasil tn ", tn)
	fmt.Println("hasil fp ", fp)
	fmt.Println("hasil fn ", fn)
	fmt.Println("AKURASI = ", float64((tp+tn)/(tp+fp+fn+tn)))
	fmt.Println("PRECISSION = ", float64((tp)/(tp+fp)))
	fmt.Println("RECALL = ", float64((tp)/(tp+fn)))
	// fmt.Println("SPECIFICITY = ", float64((tn)/(tn+fp)))
	fmt.Println("F1 SCORE = ", float64(2*(recall*precission)/(recall+precission)))
}
