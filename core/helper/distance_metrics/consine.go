package distance_metrics

import (
	"fp_datmin_03_19_92/core/structs"
	"math"
)

// ConsineSimilarity ...
func ConsineSimilarity(instanceSatu structs.DiabetesRecord, instanceDua structs.DiabetesRecord) float64 {
	var se, sNorm, dNorm, consine float64

	// hitung instancesatu * instancedua
	se += instanceSatu.Pregnancies * instanceDua.Pregnancies
	se += instanceSatu.Glucose * instanceDua.Glucose
	se += instanceSatu.BloodPressure * instanceDua.BloodPressure
	se += instanceSatu.SkinThickness * instanceDua.SkinThickness
	se += instanceSatu.Insulin * instanceDua.Insulin
	se += instanceSatu.BMI * instanceDua.BMI
	se += instanceSatu.DiabetesPedigreeFunction * instanceDua.DiabetesPedigreeFunction
	se += instanceSatu.Age * instanceDua.Age

	// hitung instancesatu norm
	sNorm += math.Pow(instanceSatu.Pregnancies, 2)
	sNorm += math.Pow(instanceSatu.Glucose, 2)
	sNorm += math.Pow(instanceSatu.BloodPressure, 2)
	sNorm += math.Pow(instanceSatu.SkinThickness, 2)
	sNorm += math.Pow(instanceSatu.Insulin, 2)
	sNorm += math.Pow(instanceSatu.BMI, 2)
	sNorm += math.Pow(instanceSatu.DiabetesPedigreeFunction, 2)
	sNorm += math.Pow(instanceSatu.Age, 2)
	sNorm = float64(math.Sqrt(sNorm))

	// hitung instancedua norm
	dNorm += math.Pow(instanceDua.Pregnancies, 2)
	dNorm += math.Pow(instanceDua.Glucose, 2)
	dNorm += math.Pow(instanceDua.BloodPressure, 2)
	dNorm += math.Pow(instanceDua.SkinThickness, 2)
	dNorm += math.Pow(instanceDua.Insulin, 2)
	dNorm += math.Pow(instanceDua.BMI, 2)
	dNorm += math.Pow(instanceDua.DiabetesPedigreeFunction, 2)
	dNorm += math.Pow(instanceDua.Age, 2)
	dNorm = float64(math.Sqrt(dNorm))

	// get similarity
	if se < 0 {
		se = 0 - se
	}
	consine = se / (sNorm * dNorm)

	// fmt.Println(int64(math.Sqrt(distance)))
	return float64(1 / consine)
}
