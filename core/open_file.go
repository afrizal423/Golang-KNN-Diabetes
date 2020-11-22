package core

import (
	"encoding/csv"
	"fp_datmin_03_19_92/core/helper/error_handle"
	"fp_datmin_03_19_92/core/structs"
	"io"
	"os"
	"strconv"
)

func Buka_file(path string) []structs.DiabetesRecord {
	diabetdata, err := os.Open(path)
	error_handle.ErrHandle(err)
	defer diabetdata.Close()

	reader := csv.NewReader(diabetdata)
	reader.Comma = ','

	var dataset []structs.DiabetesRecord

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		error_handle.ErrHandle(err)

		dataset = append(dataset, ParsingDataDiabet(record))
	}
	return dataset
}

func ParsingDataDiabet(data []string) structs.DiabetesRecord {
	var diabet structs.DiabetesRecord

	diabet.Pregnancies, _ = strconv.ParseFloat(data[0], 64)
	diabet.Glucose, _ = strconv.ParseFloat(data[1], 64)
	diabet.BloodPressure, _ = strconv.ParseFloat(data[2], 64)
	diabet.SkinThickness, _ = strconv.ParseFloat(data[3], 64)
	diabet.Insulin, _ = strconv.ParseFloat(data[4], 64)
	diabet.BMI, _ = strconv.ParseFloat(data[5], 64)
	diabet.DiabetesPedigreeFunction, _ = strconv.ParseFloat(data[6], 64)
	diabet.Age, _ = strconv.ParseFloat(data[7], 64)
	diabet.Outcome, _ = strconv.ParseInt(data[8], 10, 64)

	return diabet

}
