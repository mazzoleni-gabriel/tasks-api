package tasks

import "time"

type Task struct {
	ID          uint
	Summary     string
	PerformedAt time.Time
	CreatedBy   uint
}

type SearchFilters struct {
	CreatedBy *uint
}
