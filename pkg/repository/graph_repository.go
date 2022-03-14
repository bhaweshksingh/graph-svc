package repository

import (
	"context"
	"fmt"
	"github.com/gocarina/gocsv"
	"graph-svc/pkg/graphinfo/model"
	"os"
	"sort"
	"time"
)

//go:generate moq -out mock/GraphRepository.go -pkg mock . GraphRepository
type GraphRepository interface {
	GetDataPoints(ctx context.Context) ([]model.DataPoint, error)
}

type gormGraphRepository struct {
	csvPath string
}

type Visit struct {
	Id         string    `csv:"id"`
	DateTime   time.Time `csv:"timestamp"`
	ResidentId string    `csv:"residentId"`
	Location   string    `csv:"location"`
}

func (gbr *gormGraphRepository) GetDataPoints(ctx context.Context) ([]model.DataPoint, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	in, err := os.Open(gbr.csvPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file from path: %q", gbr.csvPath)
	}
	defer in.Close()

	var visits []Visit

	if err := gocsv.UnmarshalFile(in, &visits); err != nil {
		return nil, fmt.Errorf("failed to parse visit data from CSV")
	}

	visitsByDay := make(map[string]int)

	for _, visit := range visits {
		toRound := visit.DateTime
		roundedTime := time.Date(toRound.Year(), toRound.Month(), toRound.Day(), 0, 0, 0, 0, toRound.Location())

		formattedTime := roundedTime.Format(time.RFC3339)
		if count, ok := visitsByDay[formattedTime]; ok {
			visitsByDay[formattedTime] = count + 1
		} else {
			visitsByDay[formattedTime] = 1
		}
	}

	dataPoints := make([]model.DataPoint, 0)

	for visitDay, visitCount := range visitsByDay {
		parsedTime, _ := time.Parse(time.RFC3339, visitDay)
		dataPoints = append(dataPoints, model.DataPoint{Day: parsedTime, VisitCount: visitCount})
	}

	sort.Slice(dataPoints[:], func(i, j int) bool {
		return dataPoints[i].Day.Before(dataPoints[j].Day)
	})
	return dataPoints, nil
}

func NewGraphRepository(csvPath string) GraphRepository {
	return &gormGraphRepository{
		csvPath: csvPath,
	}
}
