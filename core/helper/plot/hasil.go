package plot

import (
	"bytes"
	"log"
	"os"

	chart "github.com/wcharczuk/go-chart/v2"
)

func Hasil(manhattan []float64, minkowski []float64, braycurtis []float64, canberra []float64, euclidean []float64) {
	// fmt.Println(braycurtis, manhattan)

	ts1 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(7),
		},
		Name:    "Manhattan",
		XValues: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		YValues: manhattan,
	}

	ts2 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
		Name:    "MinkowskiDistance",
		XValues: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		YValues: minkowski,
	}
	ts3 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(1),
		},
		Name:    "BrayCurtisDistance",
		XValues: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		YValues: braycurtis,
	}
	ts4 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(2),
		},
		Name:    "CanberraDistance",
		XValues: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		YValues: canberra,
	}
	ts5 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(9),
		},
		Name:    "Euclidean",
		XValues: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		YValues: euclidean,
	}

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 260,
			},
		},

		XAxis: chart.XAxis{
			Name: "Nilai K",
		},

		YAxis: chart.YAxis{
			Name: "Akurasi",
		},

		Series: []chart.Series{
			ts1,
			ts2,
			ts3,
			ts4,
			ts5,
		},
	}
	//note we have to do this as a separate step because we need a reference to graph
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Fatal(err)
	}

	fo, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}

	if _, err := fo.Write(buffer.Bytes()); err != nil {
		panic(err)
	}
}
