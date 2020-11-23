package core

import (
	"fp_datmin_03_19_92/core/structs"
)

type mean struct {
	Kolom string
	Hasil float64
}

func Hilangkan_0(data []structs.DiabetesRecord, bantu []structs.DiabetesRecord) []structs.DiabetesRecord {

	hasil := func(data []structs.DiabetesRecord, kolom string) float64 {

		nilai_Glucose := 0.0
		nilai_BloodPressure := 0.0
		nilai_SkinThickness := 0.0
		nilai_BMI := 0.0
		nilai_Insulin := 0.0

		for i := range data {
			// rulenya := mean{}
			if float64(data[i].Glucose) != float64(0) {
				nilai_Glucose += data[i].Glucose
			}
			if float64(data[i].BloodPressure) != float64(0) {
				nilai_BloodPressure += data[i].BloodPressure
			}
			if float64(data[i].SkinThickness) != float64(0) {
				nilai_SkinThickness += data[i].SkinThickness
			}
			if float64(data[i].BMI) != float64(0) {
				nilai_BMI += data[i].BMI
			}
			if float64(data[i].Insulin) != float64(0) {
				nilai_Insulin += data[i].Insulin
			}

		}

		if kolom == "Glucose" {
			return nilai_Glucose / float64(len(data))
		}
		if kolom == "BloodPressure" {
			return nilai_BloodPressure / float64(len(data))
		}
		if kolom == "SkinThickness" {
			return nilai_SkinThickness / float64(len(data))
		}
		if kolom == "BMI" {
			return nilai_BMI / float64(len(data))
		}
		if kolom == "Insulin" {
			return nilai_Insulin / float64(len(data))
		}

		return 0.0
		// fmt.Println(nilai_Glucose / float64(len(data)))
		// fmt.Println(len(data))
	}
	mean_Glucose := hasil(bantu, "Glucose")
	mean_BloodPressure := hasil(bantu, "BloodPressure")
	mean_SkinThickness := hasil(bantu, "SkinThickness")
	mean_BMI := hasil(bantu, "BMI")
	mean_Insulin := hasil(bantu, "Insulin")

	for i := range data {
		// me := mean{}
		if float64(data[i].Glucose) == float64(0) {
			// fmt.Println("a asdasdasd sd", data[i].Insulin)
			// fmt.Println("Sebelum ", data[i].Glucose)
			data[i].Glucose = mean_Glucose
			// fmt.Println("sesudah ", data[i].Glucose)
		}
		if float64(data[i].BloodPressure) == float64(0) {
			// fmt.Println("a asdasdasd sd", data[i].Insulin)
			// count_mean(bantu, "BloodPressure")
			// fmt.Println(count_mean(bantu, "BloodPressure"))
			data[i].BloodPressure = mean_BloodPressure
		}
		if float64(data[i].SkinThickness) == float64(0) {
			// fmt.Println("a asdasdasd sd", data[i].Insulin)
			// count_mean(bantu, "SkinThickness")
			// fmt.Println(count_mean(bantu, "SkinThickness"))
			data[i].SkinThickness = mean_SkinThickness
		}
		if float64(data[i].BMI) == float64(0) {
			// fmt.Println("a asdasdasd sd", data[i].Insulin)
			// count_mean(bantu, "BMI")
			data[i].BMI = mean_BMI
			// fmt.Println(count_mean(bantu, "BMI"))
		}
		if float64(data[i].Insulin) == float64(0) {
			// fmt.Println("a asdasdasd sd", data[i].Insulin)
			// count_mean(bantu, "Insulin")
			// fmt.Println("Sebelum ", data[i].Insulin)
			data[i].Insulin = mean_Insulin
			// fmt.Println("sesudah ", data[i].Insulin)
			// fmt.Println(count_mean(	bantu, "Insulin"))
		}
	}
	// fmt.Println("jumlah data sesudah", data[0].Insulin)
	var dataset []structs.DiabetesRecord
	for i := range data {
		// memasukkan ke array dataset
		dataset = append(dataset, data[i])

	}
	return dataset
	// fmt.Println(hasil(bantu, "Insulin"))
}
