package core

import (
	"fmt"
	"fp_datmin_03_19_92/core/helper/distance_metrics"
	"fp_datmin_03_19_92/core/structs"
	"sort"
)

type Hasil struct {
	prediksi int64
	aktual   int64
}

func KNearestNeighbor(k int, dataTrain []structs.DiabetesRecord, dataTest []structs.DiabetesRecord, distance_matrics string) float64 {
	var predictions []int64
	var simpan_hasil []Hasil
	for x := range dataTest {
		tmp := Hasil{}
		// fmt.Println(dataTest[x])
		// fmt.Println("\ntes datanya ", dataTest[x])
		neighbors := getNeighbors(dataTrain, dataTest[x], k, distance_matrics)
		result := getResponse(neighbors)
		// fmt.Println(result[0].key)
		predictions = append(predictions, result[0].key)
		tmp.prediksi = result[0].key
		tmp.aktual = dataTest[x].Outcome
		simpan_hasil = append(simpan_hasil, tmp)
		// fmt.Println("Predicted: ", result[0].key, ", Actual: ", dataTest[x].Outcome)
	}
	// fmt.Println(predictions)
	// fmt.Println(getAccuracy(dataTest, predictions))
	return ConfusionMatrix(simpan_hasil, k, distance_matrics)
}

type distancePair struct {
	record   structs.DiabetesRecord
	distance float64
}

type distancePairs []distancePair

func (slice distancePairs) Len() int           { return len(slice) }
func (slice distancePairs) Less(i, j int) bool { return slice[i].distance < slice[j].distance }
func (slice distancePairs) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

func getNeighbors(trainingSet []structs.DiabetesRecord, testRecord structs.DiabetesRecord, k int, distance_matrics string) []structs.DiabetesRecord {
	var distances distancePairs
	// fmt.Println(len(trainingSet))
	for i := range trainingSet {
		if distance_matrics == "manhattan" {
			dist := distance_metrics.Manhattan(testRecord, trainingSet[i])

			distances = append(distances, distancePair{trainingSet[i], dist})
		}
		if distance_matrics == "minkowski" {
			dist := distance_metrics.MinkowskiDistance(testRecord, trainingSet[i], 3)

			distances = append(distances, distancePair{trainingSet[i], dist})
		}
		if distance_matrics == "braycurtis" {
			dist := distance_metrics.BrayCurtisDistance(testRecord, trainingSet[i])

			distances = append(distances, distancePair{trainingSet[i], dist})
		}
		if distance_matrics == "canberra" {
			dist := distance_metrics.CanberraDistance(testRecord, trainingSet[i])

			distances = append(distances, distancePair{trainingSet[i], dist})
		}
		if distance_matrics == "euclidean" {
			dist := distance_metrics.EuclidianDistance(testRecord, trainingSet[i])

			distances = append(distances, distancePair{trainingSet[i], dist})
		}
		if distance_matrics == "l1" {
			dist := distance_metrics.L1Distance(testRecord, trainingSet[i])

			distances = append(distances, distancePair{trainingSet[i], dist})
		}
		if distance_matrics == "consine" {
			dist := distance_metrics.ConsineSimilarity(testRecord, trainingSet[i])

			distances = append(distances, distancePair{trainingSet[i], dist})
		} else {
			dist := distance_metrics.EuclidianDistance(testRecord, trainingSet[i])

			distances = append(distances, distancePair{trainingSet[i], dist})
		}

		// fmt.Println(trainingSet[i], dist)
	}
	sort.Sort(distances)
	// fmt.Println(distances)
	var neighbors []structs.DiabetesRecord

	// fmt.Println(distances)
	// for i := range trainingSet {
	// 	fmt.Println(distances[i])
	// }

	for x := 0; x < k; x++ {
		neighbors = append(neighbors, distances[x].record)
	}
	// fmt.Println(neighbors)
	return neighbors
}

type classVote struct {
	key   int64
	value int
}

type sortedClassVotes []classVote

func (scv sortedClassVotes) Len() int           { return len(scv) }
func (scv sortedClassVotes) Less(i, j int) bool { return scv[i].value < scv[j].value }
func (scv sortedClassVotes) Swap(i, j int)      { scv[i], scv[j] = scv[j], scv[i] }

func getResponse(neighbors []structs.DiabetesRecord) sortedClassVotes {
	classVotes := make(map[int64]int)

	for x := range neighbors {
		response := neighbors[x].Outcome
		if contains(classVotes, response) {
			classVotes[response] += 1
		} else {
			classVotes[response] = 1
		}
	}

	scv := make(sortedClassVotes, len(classVotes))
	i := 0
	for k, v := range classVotes {
		scv[i] = classVote{k, v}
		i++
	}

	sort.Sort(sort.Reverse(scv))
	// fmt.Println(scv)
	return scv
}

func contains(votesMap map[int64]int, name int64) bool {
	for s, _ := range votesMap {
		if s == name {
			return true
		}
	}

	return false
}

func getAccuracy(testSet []structs.DiabetesRecord, predictions []int64) float64 {
	correct := 0
	incorrect := 0

	for x := range testSet {
		if testSet[x].Outcome == predictions[x] {
			correct += 1
		}
		incorrect += 1
	}
	fmt.Println("incorrect", incorrect)
	fmt.Println("correct", correct)
	pr := float64(correct) / (float64(correct) + float64(incorrect))
	fmt.Printf("precision %v \n", pr)
	return (float64(correct) / float64(len(testSet))) * 100.00
}
