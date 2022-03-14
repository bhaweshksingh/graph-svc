package dto

import (
	"graph-svc/pkg/graphinfo/model"
	"time"
)

type ChartData struct {
	DataPoints []Point
	Trendline  LinearTrendline
}

const (
	dayLayout = "02 Jan 2006"
)

func (d *ChartData) CalculateTrendline([]model.DataPoint) {
	d.Trendline = LinearTrendline{
		Start: Point{},
		End:   Point{},
	}
}

type LinearTrendline struct {
	Start Point
	End   Point
}

type Point struct {
	Day        string
	VisitCount int
}

func NewChartDataResponse(dataPoints []model.DataPoint) ChartData {
	chartData := &ChartData{}
	for _, point := range dataPoints {
		chartData.DataPoints = append(chartData.DataPoints, Point{getDateString(point.Day), point.VisitCount})
	}

	chartData.CalculateTrendline(dataPoints)

	return *chartData
}

func getDateString(roundedTime time.Time) string {
	return roundedTime.Format(dayLayout)
}
