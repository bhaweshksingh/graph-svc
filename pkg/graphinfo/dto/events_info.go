package dto

import (
	"graph-svc/pkg/graphinfo/model"
	"time"
)

type ChartData struct {
	DataPoints []Point
	Trendline  LinearTrendline
}

func (d *ChartData) CalculateTrendline() {
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
	Day        time.Time
	VisitCount int
}

func NewChartDataResponse(dataPoints []model.DataPoint) ChartData {
	chartData := &ChartData{}
	for _, point := range dataPoints {
		chartData.DataPoints = append(chartData.DataPoints, Point{point.Day, point.VisitCount})
	}

	chartData.CalculateTrendline()

	return *chartData
}
