package plot

import (
	"bytes"
	"log"
	"os"
	"time"

	chart "github.com/wcharczuk/go-chart/v2"
)

// Hasil ...
func Hasil(manhattan []float64, minkowski []float64, braycurtis []float64, canberra []float64, euclidean []float64, L1Dist []float64, consine []float64, jumlah []float64) {
	// fmt.Println(braycurtis, manhattan)

	ts1 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(3),
		},
		Name:    "Manhattan",
		XValues: jumlah,
		YValues: manhattan,
	}

	ts2 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
		Name:    "MinkowskiDistance",
		XValues: jumlah,
		YValues: minkowski,
	}
	ts3 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(1),
		},
		Name:    "BrayCurtisDistance",
		XValues: jumlah,
		YValues: braycurtis,
	}
	ts4 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(2),
		},
		Name:    "CanberraDistance",
		XValues: jumlah,
		YValues: canberra,
	}
	ts5 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(9),
		},
		Name:    "Euclidean",
		XValues: jumlah,
		YValues: euclidean,
	}
	ts6 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetAlternateColor(2),
		},
		Name:    "L1 Distance",
		XValues: jumlah,
		YValues: L1Dist,
	}
	ts7 := chart.ContinuousSeries{ //TimeSeries{
		Style: chart.Style{
			StrokeColor: chart.GetAlternateColor(3),
		},
		Name:    "Consine Distance",
		XValues: jumlah,
		YValues: consine,
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
			ts6,
			ts7,
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
	tm := time.Now()
	ass := tm.Format("01-02-2006 15:04:05")
	fo, err := os.Create("plot_hasil/output " + ass + ".png")
	if err != nil {
		panic(err)
	}

	if _, err := fo.Write(buffer.Bytes()); err != nil {
		panic(err)
	}
}
