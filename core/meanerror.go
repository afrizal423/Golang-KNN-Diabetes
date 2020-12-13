package core

// MeanError ...
func MeanError(array []Hasil, k int) float64 {
	var (
		// mean_error float64 = 0
		number int64 = 0
	)
	for x := range array {
		if array[x].aktual != array[x].prediksi {
			number++
		}
	}
	// fmt.Println(float64(number) / float64(len(array)))
	// fmt.Println(float64(len(array)))
	return float64(number) / float64(len(array))
}
