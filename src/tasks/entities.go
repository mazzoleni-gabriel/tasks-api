package tasks

import "time"

type Task struct {
	Summary     string
	PerformedAt time.Time
	CreatedBy   uint
}
