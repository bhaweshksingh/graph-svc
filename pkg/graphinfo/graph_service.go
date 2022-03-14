package graphinfo

import (
	"context"
	"fmt"
	"graph-svc/pkg/graphinfo/dto"
	"graph-svc/pkg/repository"
)

type Service interface {
	GetChartData(ctx context.Context) (dto.ChartData, error)
}

type GraphService struct {
	repository repository.GraphRepository
}


func (es *GraphService) GetChartData(ctx context.Context) (dto.ChartData, error) {
	dataPoints, err := es.repository.GetDataPoints(ctx)
	if err != nil {
		return dto.ChartData{}, fmt.Errorf("Service.GetDataPoints: %+v", err)
	}

	return dto.NewChartDataResponse(dataPoints), nil
}

func NewGraphService(repository repository.GraphRepository) Service {
	return &GraphService{
		repository: repository,
	}
}
