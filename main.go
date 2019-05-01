package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"median-finding/algorithms"
	"median-finding/util"
	"time"
)

const(
	// Labels + File name of the plot generated
	PlotTitle = "Laufzeitvergleich"
	PlotXAxisLabel = "n"
	PlotYAxisLabel = "Laufzeit in Sekunden"
	LabelLinearSelection = "Repeated Linear Selection"
	LabelSortFirst = "Sort First"
	LabelQuickSelect = "QuickSelect"
	LabelMedianOfMedians = "Median of Medians"

	// File the plot will be saved to
	PlotFile = "comparison.png"

	// Step for each point in the x axis of the plot
	PlotStep = 1000

	// Total Length of the random list
	ListLength = 30000
)

func main() {
	list := util.GenerateRandomList(ListLength)
	makePlot(list)
}

func makePlot(list []int) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = PlotTitle
	p.X.Label.Text = PlotXAxisLabel
	p.Y.Label.Text = PlotYAxisLabel

	err = plotutil.AddLinePoints(p,
		LabelLinearSelection, getPoints(algorithms.RepeatedLinearSelection, PlotStep, list),
		LabelSortFirst, getPoints(algorithms.SortFirst, PlotStep, list),
		LabelQuickSelect, getPoints(algorithms.QuickSelect, PlotStep, list),
		LabelMedianOfMedians, getPoints(algorithms.MedianOfMedians, PlotStep, list))

	if err != nil {
		panic(err)
	}

	if err := p.Save(6*vg.Inch, 6*vg.Inch, PlotFile); err != nil {
		panic(err)
	}
}

func getPoints(algorithm algorithms.Algorithm, step int, list []int) plotter.XYs {
	fmt.Print("STARTING COLLECTING PLOTS FOR ALGORITHM ", algorithm)
	points := make(plotter.XYs, len(list) / step)

	for i := 1; i < len(points); i++ {
		// Debug the percentage
		fmt.Println((step*i*100)/len(list), "%")

		// Measure the time for the sub list
		slice := list[:i*step]
		startTime := time.Now()
		algorithms.FindMedian(algorithm, slice)
		elapsedTime := time.Since(startTime)

		// Add the measurements to the plot
		points[i].X = float64(i*step)
		points[i].Y = float64(elapsedTime.Seconds())
	}
	return points
}
