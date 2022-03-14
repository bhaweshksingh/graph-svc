package graphinfo

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"graph-svc/pkg/graphinfo/dto"
	"graph-svc/pkg/graphinfo/model"
	"graph-svc/pkg/repository/mock"
	"testing"
	"time"
)

func TestGormGraphRepository_GetChartData(t *testing.T) {
	ctx := context.Background()
	timeStart := time.Now()
	dataPoints := []model.DataPoint{{Day: timeStart, VisitCount: 1}, {Day: timeStart.Add(time.Hour * 24), VisitCount: 1}}
	startDayString := timeStart.Format(dto.DateLayout)
	lastDayString := timeStart.Add(time.Hour * 24).Format(dto.DateLayout)
	expectedDataPoints := dto.ChartData{
		DataPoints: []dto.Point{
			{Day: startDayString, VisitCount: 1},
			{Day: lastDayString, VisitCount: 1},
		},
		Trendline: &dto.LinearTrendline{
			Start: dto.Point{Day: startDayString, VisitCount: 1},
			End:   dto.Point{Day: lastDayString, VisitCount: 1},
		},
		TrendlinePoints: []dto.Point{
			{Day: startDayString, VisitCount: 1},
			{Day: lastDayString, VisitCount: 1},
		},
	}
	repositoryMock := mock.GraphRepositoryMock{
		GetChartDataFunc: func(ctx context.Context) ([]model.DataPoint, error) {
			return dataPoints, nil
		}}

	service := NewGraphService(&repositoryMock)

	actualDataPoints, err := service.GetChartData(ctx)

	assert.Nil(t, err)
	assert.Equal(t, actualDataPoints, expectedDataPoints)
}

func TestGormGraphRepository_GetChartData_fails(t *testing.T) {
	ctx := context.Background()
	mockError := errors.New("failed to read CSV data")

	repositoryMock := mock.GraphRepositoryMock{
		GetChartDataFunc: func(ctx context.Context) ([]model.DataPoint, error) {
			return nil, mockError
		}}

	service := NewGraphService(&repositoryMock)

	actualChartData, err := service.GetChartData(ctx)

	assert.Equal(t, 0, len(actualChartData.DataPoints))
	assert.Contains(t, err.Error(), mockError.Error())
}

func TestGormGraphRepository_GetChartData_passesForEmptyData(t *testing.T) {
	ctx := context.Background()

	repositoryMock := mock.GraphRepositoryMock{
		GetChartDataFunc: func(ctx context.Context) ([]model.DataPoint, error) {
			return []model.DataPoint{}, nil
		}}

	service := NewGraphService(&repositoryMock)

	actualChartData, err := service.GetChartData(ctx)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(actualChartData.DataPoints))
}
