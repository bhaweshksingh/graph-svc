package dto

import (
	"graph-svc/pkg/graphinfo/model"
	"math"
	"time"
)

type ChartData struct {
	DataPoints      []Point
	Trendline       *LinearTrendline
	TrendlinePoints []Point
}

const (
	DateLayout = "02 Jan 2006"
)

type TrendlinePoint struct {
	DayNum int
	Count  int
}

func (d *ChartData) CalculateTrendline(dataPoints []model.DataPoint) {

	d.Trendline = &LinearTrendline{
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

	sumDaysSinceStart := 0.0
	sumVisits := 0.0
	sumDaysSinceStartSquared := 0.0
	sumVisitsSquared := 0.0
	sumVisitsIntoDaysSinceStart := 0.0

	for _, point := range trendlinePoints {
		sumDaysSinceStart += float64(point.DayNum)
		sumVisits += float64(point.Count)
		sumDaysSinceStartSquared += math.Pow(float64(point.DayNum), 2)
		sumVisitsSquared += math.Pow(float64(point.Count), 2)

		sumVisitsIntoDaysSinceStart += float64(point.DayNum * point.Count)
	}

	b := (float64(len(trendlinePoints))*sumVisitsIntoDaysSinceStart - sumDaysSinceStart*sumVisits) / (float64(len(trendlinePoints))*sumDaysSinceStartSquared - math.Pow(sumDaysSinceStart, 2))

	a := sumVisits/float64(len(trendlinePoints)) - sumDaysSinceStart*b/float64(len(trendlinePoints))

	trendlineCalculated := make([]Point, len(dataPoints))

	for i, point := range trendlinePoints {
		visitCountTrend := a + b*float64(point.DayNum)
		trendlineCalculated[i] = Point{Day: dataPoints[i].Day.Format(DateLayout), VisitCount: visitCountTrend}
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
	return roundedTime.Format(DateLayout)
}
