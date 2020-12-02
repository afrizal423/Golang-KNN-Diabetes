package distance_metrics

import (
	"fp_datmin_03_19_92/core/structs"
	"math"
)

func EuclidianDistance(instanceSatu structs.DiabetesRecord, instanceDua structs.DiabetesRecord) int64 {
	var distance float64

	distance += math.Pow((instanceSatu.Pregnancies - instanceDua.Pregnancies), 2)
	distance += math.Pow((instanceSatu.Glucose - instanceDua.Glucose), 2)
	distance += math.Pow((instanceSatu.BloodPressure - instanceDua.BloodPressure), 2)
	distance += math.Pow((instanceSatu.SkinThickness - instanceDua.SkinThickness), 2)
	distance += math.Pow((instanceSatu.Insulin - instanceDua.Insulin), 2)
	distance += math.Pow((instanceSatu.BMI - instanceDua.BMI), 2)
	distance += math.Pow((instanceSatu.DiabetesPedigreeFunction - instanceDua.DiabetesPedigreeFunction), 2)
	distance += math.Pow((instanceSatu.Age - instanceDua.Age), 2)

	return int64(math.Sqrt(distance))
}
