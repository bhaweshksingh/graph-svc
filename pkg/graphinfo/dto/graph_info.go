package dto

import (
	"graph-svc/pkg/graphinfo/model"
	"math"
	"time"
)

type ChartData struct {
	DataPoints      []Point
	Trendline       LinearTrendline
	TrendlinePoints []Point
}

const (
	dateLayout = "02 Jan 2006"
)

type TrendlinePoint struct {
	DayNum int
	Count  int
}

func (d *ChartData) CalculateTrendline(dataPoints []model.DataPoint) {

	d.Trendline = LinearTrendline{
		Start: Point{},
		End:   Point{},
	}

	if len(dataPoints) == 0 {
		return
	}

	trendlinePoints := make([]TrendlinePoint, len(dataPoints))

	firstDay := dataPoints[0].Day
	trendlinePoints[0] = TrendlinePoint{DayNum: 0, Count: dataPoints[0].VisitCount}

	for i := 1; i < len(dataPoints); i++ {
		daySinceFirst := dataPoints[i].Day.Sub(firstDay)
		dayNum := int(daySinceFirst.Hours() / 24)
		trendlinePoints[i] = TrendlinePoint{DayNum: dayNum, Count: dataPoints[i].VisitCount}
	}

	sumX := 0.0
	sumY := 0.0
	sumXSquared := 0.0
	sumYSquared := 0.0
	sumXY := 0.0

	for _, point := range trendlinePoints {
		sumX += float64(point.DayNum)
		sumY += float64(point.Count)
		sumXSquared += math.Pow(float64(point.DayNum), 2)
		sumYSquared += math.Pow(float64(point.Count), 2)

		sumXY += float64(point.DayNum * point.Count)
	}

	b := (float64(len(trendlinePoints))*sumXY - sumX*sumY) / (float64(len(trendlinePoints))*sumXSquared - math.Pow(sumX, 2))

	a := sumY/float64(len(trendlinePoints)) - sumX*b/float64(len(trendlinePoints))

	trendlineCalculated := make([]Point, len(dataPoints))

	for i, point := range trendlinePoints {
		visitCountTrend := a + b*float64(point.DayNum)
		trendlineCalculated[i] = Point{Day: dataPoints[i].Day.Format(dateLayout), VisitCount: visitCountTrend}
	}

	d.Trendline.Start = trendlineCalculated[0]
	d.Trendline.End = trendlineCalculated[len(trendlineCalculated)-1]

	d.TrendlinePoints = trendlineCalculated

}

type LinearTrendline struct {
	Start Point
	End   Point
}

type Point struct {
	Day        string
	VisitCount float64
}

func NewChartDataResponse(dataPoints []model.DataPoint) ChartData {
	chartData := &ChartData{}
	for _, point := range dataPoints {
		chartData.DataPoints = append(chartData.DataPoints, Point{getDateString(point.Day), float64(point.VisitCount)})
	}

	chartData.CalculateTrendline(dataPoints)

	return *chartData
}

func getDateString(roundedTime time.Time) string {
	return roundedTime.Format(dateLayout)
}
