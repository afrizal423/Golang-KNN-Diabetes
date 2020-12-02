package distance_metrics

import (
	"fp_datmin_03_19_92/core/structs"
	"math"
)

func BrayCurtisDistance(instanceSatu structs.DiabetesRecord, instanceDua structs.DiabetesRecord) float64 {
	var numerator, denominator float64
	numerator += math.Abs(instanceSatu.Pregnancies - instanceDua.Pregnancies)
	numerator += math.Abs(instanceSatu.Glucose - instanceDua.Glucose)
	numerator += math.Abs(instanceSatu.BloodPressure - instanceDua.BloodPressure)
	numerator += math.Abs(instanceSatu.SkinThickness - instanceDua.SkinThickness)
	numerator += math.Abs(instanceSatu.Insulin - instanceDua.Insulin)
	numerator += math.Abs(instanceSatu.BMI - instanceDua.BMI)
	numerator += math.Abs(instanceSatu.DiabetesPedigreeFunction - instanceDua.DiabetesPedigreeFunction)
	numerator += math.Abs(instanceSatu.Age - instanceDua.Age)

	denominator += math.Abs(instanceSatu.Pregnancies + instanceDua.Pregnancies)
	denominator += math.Abs(instanceSatu.Glucose + instanceDua.Glucose)
	denominator += math.Abs(instanceSatu.BloodPressure + instanceDua.BloodPressure)
	denominator += math.Abs(instanceSatu.SkinThickness + instanceDua.SkinThickness)
	denominator += math.Abs(instanceSatu.Insulin + instanceDua.Insulin)
	denominator += math.Abs(instanceSatu.BMI + instanceDua.BMI)
	denominator += math.Abs(instanceSatu.DiabetesPedigreeFunction + instanceDua.DiabetesPedigreeFunction)
	denominator += math.Abs(instanceSatu.Age + instanceDua.Age)

	return numerator / denominator
}
