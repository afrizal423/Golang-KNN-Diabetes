package distance_metrics

import (
	"fp_datmin_03_19_92/core/structs"
	"math"
)

func Manhattan(instanceSatu structs.DiabetesRecord, instanceDua structs.DiabetesRecord) float64 {

	var distance float64

	distance += math.Abs(instanceSatu.Pregnancies - instanceDua.Pregnancies)
	distance += math.Abs(instanceSatu.Glucose - instanceDua.Glucose)
	distance += math.Abs(instanceSatu.BloodPressure - instanceDua.BloodPressure)
	distance += math.Abs(instanceSatu.SkinThickness - instanceDua.SkinThickness)
	distance += math.Abs(instanceSatu.Insulin - instanceDua.Insulin)
	distance += math.Abs(instanceSatu.BMI - instanceDua.BMI)
	distance += math.Abs(instanceSatu.DiabetesPedigreeFunction - instanceDua.DiabetesPedigreeFunction)
	distance += math.Abs(instanceSatu.Age - instanceDua.Age)
	// fmt.Println(int64(distance))
	return float64(distance)
}
