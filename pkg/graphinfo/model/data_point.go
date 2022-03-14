package model

import (
	"time"
)

type DataPoint struct {
	Day time.Time
	VisitCount int
}
