package entities

import "time"

type (
	Transaction struct {
		ID        int64
		RequestID string
		Date      time.Time
		Value     float64
	}
)
