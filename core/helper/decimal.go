package helper

import "math"

// Semua fungsi untuk meng-generate angka float.
// maksudnya ingin menampilkan berapa angka dibelakang koma

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
