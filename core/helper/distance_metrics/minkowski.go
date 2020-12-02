package distance_metrics

import (
	"fp_datmin_03_19_92/core/structs"
	"math"
)

func MinkowskiDistance(instanceSatu structs.DiabetesRecord, instanceDua structs.DiabetesRecord, p float64) float64 {
	var distance float64

	distance += math.Pow(math.Abs(instanceSatu.Pregnancies-instanceDua.Pregnancies), p)
	distance += math.Pow(math.Abs(instanceSatu.Glucose-instanceDua.Glucose), p)
	distance += math.Pow(math.Abs(instanceSatu.BloodPressure-instanceDua.BloodPressure), p)
	distance += math.Pow(math.Abs(instanceSatu.SkinThickness-instanceDua.SkinThickness), p)
	distance += math.Pow(math.Abs(instanceSatu.Insulin-instanceDua.Insulin), p)
	distance += math.Pow(math.Abs(instanceSatu.BMI-instanceDua.BMI), p)
	distance += math.Pow(math.Abs(instanceSatu.DiabetesPedigreeFunction-instanceDua.DiabetesPedigreeFunction), p)
	distance += math.Pow(math.Abs(instanceSatu.Age-instanceDua.Age), p)

	return math.Pow(distance, 1/p)
}
