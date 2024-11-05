package response

import (
	"time"

	"github.com/yskikuchi/go-nextjs-app/model"
)

type BookingSummary struct {
	Car       model.Car `json:"car"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
