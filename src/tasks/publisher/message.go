package publisher

import (
	"tasks-api/src/tasks"
	"time"
)

type Operation string

const (
	CreateOperation Operation = "create"
	UpdateOperation Operation = "update"
	DeleteOperation Operation = "delete"
)

type Message struct {
	Entity    Entity    `json:"entity"`
	Operation Operation `json:"operation"`
}

type Entity struct {
	ID          uint      `json:"id"`
	Summary     string    `json:"summary"`
	PerformedAt time.Time `json:"performed_at"`
	CreatedBy   uint      `json:"created_by"`
}

func newCreateMessage(task tasks.Task) Message {
	return Message{
		Entity:    newEntity(task),
		Operation: CreateOperation,
	}
}

func newUpdateMessage(task tasks.Task) Message {
	return Message{
		Entity:    newEntity(task),
		Operation: UpdateOperation,
	}
}

func newDeleteMessage(task tasks.Task) Message {
	return Message{
		Entity:    newEntity(task),
		Operation: DeleteOperation,
	}
}

func newEntity(task tasks.Task) Entity {
	return Entity{
		ID:          task.ID,
		Summary:     task.Summary,
		PerformedAt: task.PerformedAt,
		CreatedBy:   task.CreatedBy,
	}
}
