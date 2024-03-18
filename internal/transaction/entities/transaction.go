package entities

import "time"

type Transaction struct {
	ID    int64
	Date  time.Time
	Value float64
}
