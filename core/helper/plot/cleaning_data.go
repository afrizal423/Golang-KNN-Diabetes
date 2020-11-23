package plot

import (
	"fp_datmin_03_19_92/core/structs"
	"os"

	chart "github.com/wcharczuk/go-chart/v2"
)

func Clean_data(data []structs.DiabetesRecord) {

	/*
	   The below will draw the same chart as the `basic` example, except with both the x and y axes turned on.
	   In this case, both the x and y axis ticks are generated automatically, the x and y ranges are established automatically, the canvas "box" is adjusted to fit the space the axes occupy so as not to clip.
	*/
	// jumlah := len(data)
	// var asd [767]float64
	var s []float64
	var jumlah []float64

	for i := range data {
		// fmt.Println(data[i].Insulin)
		jumlah = append(jumlah, float64(i))
		s = append(s, data[i].Insulin)
	}
	// fmt.Println(s)
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "The XAxis",
		},
		YAxis: chart.YAxis{
			Name: "The YAxis",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: jumlah,
				YValues: s,
			},
		},
	}

	f, _ := os.Create("output2.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
