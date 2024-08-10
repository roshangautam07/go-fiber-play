package utility

import (
	"log"
	"math"
	"math/rand"

	"github.com/arl/statsviz"
)

func ScatterPlot() statsviz.TimeSeriesPlot {
	// Describe the 'sine' time series.
	sine := statsviz.TimeSeries{
		Name:     "short sin",
		Unitfmt:  "%{y:.4s}B",
		GetValue: UpdateSine,
	}

	// Build a new plot, showing our sine time series
	plot, err := statsviz.TimeSeriesPlotConfig{
		Name:  "sine",
		Title: "Sine",
		Type:  statsviz.Scatter,
		InfoText: `This is an example of a 'scatter' type plot, showing a single time series.<br>
InfoText field (this) accepts any HTML tags like <b>bold</b>, <i>italic</i>, etc.`,
		YAxisTitle: "y unit",
		Series:     []statsviz.TimeSeries{sine},
	}.Build()
	if err != nil {
		log.Fatalf("failed to build timeseries plot: %v", err)
	}

	return plot
}

func BarPlot() statsviz.TimeSeriesPlot {
	// Describe the 'user logins' time series.
	logins := statsviz.TimeSeries{
		Name:     "user logins",
		Unitfmt:  "%{y:.4s}",
		GetValue: Logins,
	}

	// Describe the 'user signins' time series.
	signins := statsviz.TimeSeries{
		Name:     "user signins",
		Unitfmt:  "%{y:.4s}",
		GetValue: Signins,
	}

	// Build a new plot, showing both time series at once.
	plot, err := statsviz.TimeSeriesPlotConfig{
		Name:  "users",
		Title: "Users",
		Type:  statsviz.Bar,
		InfoText: `This is an example of a 'bar' type plot, showing 2 time series.<br>
InfoText field (this) accepts any HTML tags like <b>bold</b>, <i>italic</i>, etc.`,
		YAxisTitle: "users",
		Series:     []statsviz.TimeSeries{logins, signins},
	}.Build()
	if err != nil {
		log.Fatalf("failed to build timeseries plot: %v", err)
	}

	return plot
}

func StackedPlot() statsviz.TimeSeriesPlot {
	// Describe the 'user logins' time series.
	logins := statsviz.TimeSeries{
		Name:     "user logins",
		Unitfmt:  "%{y:.4s}",
		Type:     statsviz.Bar,
		GetValue: Logins,
	}

	// Describe the 'user signins' time series.
	signins := statsviz.TimeSeries{
		Name:     "user signins",
		Unitfmt:  "%{y:.4s}",
		Type:     statsviz.Bar,
		GetValue: Signins,
	}

	// Build a new plot, showing both time series at once.
	plot, err := statsviz.TimeSeriesPlotConfig{
		Name:    "users-stack",
		Title:   "Stacked Users",
		Type:    statsviz.Bar,
		BarMode: statsviz.Stack,
		InfoText: `This is an example of a 'bar' plot showing 2 time series stacked on top of each other with <b>BarMode:Stack</b>.<br>
InfoText field (this) accepts any HTML tags like <b>bold</b>, <i>italic</i>, etc.`,
		YAxisTitle: "users",
		Series:     []statsviz.TimeSeries{logins, signins},
	}.Build()
	if err != nil {
		log.Fatalf("failed to build timeseries plot: %v", err)
	}

	return plot
}

var val = 0.

func UpdateSine() float64 {
	val += 0.5
	return math.Sin(val)
}

func Logins() float64 {
	return (rand.Float64() + 2) * 1000
}

func Signins() float64 {
	return (rand.Float64() + 1.5) * 100
}
