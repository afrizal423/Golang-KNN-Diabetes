package distance_metrics

import (
	"fp_datmin_03_19_92/core/structs"
	"math"
)

func CanberraDistance(instanceSatu structs.DiabetesRecord, instanceDua structs.DiabetesRecord) float64 {
	var distance float64
	distance += (math.Abs(instanceSatu.Pregnancies-instanceDua.Pregnancies) / (math.Abs(instanceSatu.Pregnancies) + math.Abs(instanceDua.Pregnancies)))
	distance += (math.Abs(instanceSatu.Glucose-instanceDua.Glucose) / (math.Abs(instanceSatu.Glucose) + math.Abs(instanceDua.Glucose)))
	distance += (math.Abs(instanceSatu.BloodPressure-instanceDua.BloodPressure) / (math.Abs(instanceSatu.BloodPressure) + math.Abs(instanceDua.BloodPressure)))
	distance += (math.Abs(instanceSatu.SkinThickness-instanceDua.SkinThickness) / (math.Abs(instanceSatu.SkinThickness) + math.Abs(instanceDua.SkinThickness)))
	distance += (math.Abs(instanceSatu.Insulin-instanceDua.Insulin) / (math.Abs(instanceSatu.Insulin) + math.Abs(instanceDua.Insulin)))
	distance += (math.Abs(instanceSatu.BMI-instanceDua.BMI) / (math.Abs(instanceSatu.BMI) + math.Abs(instanceDua.BMI)))
	distance += (math.Abs(instanceSatu.DiabetesPedigreeFunction-instanceDua.DiabetesPedigreeFunction) / (math.Abs(instanceSatu.DiabetesPedigreeFunction) + math.Abs(instanceDua.DiabetesPedigreeFunction)))
	distance += (math.Abs(instanceSatu.Age-instanceDua.Age) / (math.Abs(instanceSatu.Age) + math.Abs(instanceDua.Age)))
	return distance
}
