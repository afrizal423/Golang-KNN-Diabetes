package distance_metrics

import (
	"fp_datmin_03_19_92/core/structs"
	"math"
)

// L1Distance ...
func L1Distance(instanceSatu structs.DiabetesRecord, instanceDua structs.DiabetesRecord) float64 {
	var numerator float64
	numerator += math.Abs(instanceSatu.Pregnancies - instanceDua.Pregnancies)
	numerator += math.Abs(instanceSatu.Glucose - instanceDua.Glucose)
	numerator += math.Abs(instanceSatu.BloodPressure - instanceDua.BloodPressure)
	numerator += math.Abs(instanceSatu.SkinThickness - instanceDua.SkinThickness)
	numerator += math.Abs(instanceSatu.Insulin - instanceDua.Insulin)
	numerator += math.Abs(instanceSatu.BMI - instanceDua.BMI)
	numerator += math.Abs(instanceSatu.DiabetesPedigreeFunction - instanceDua.DiabetesPedigreeFunction)
	numerator += math.Abs(instanceSatu.Age - instanceDua.Age)
	// fmt.Println(int64(math.Sqrt(distance)))
	return float64(numerator)
}
