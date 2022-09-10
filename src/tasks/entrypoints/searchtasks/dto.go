package searchtasks

import (
	"net/http"
	"strconv"
	"tasks-api/src/tasks"
	"time"
)

const createdByKey = "created_by"

type SearchTasksResponse struct {
	List []Task `json:"list"`
}

type Task struct {
	ID          uint      `json:"id"`
	Summary     string    `json:"summary"`
	PerformedAt time.Time `json:"performed_at"`
	CreatedBy   uint      `json:"created_by"`
}

func newResponse(tasks []tasks.Task) SearchTasksResponse {
	response := SearchTasksResponse{}
	for _, t := range tasks {
		response.List = append(response.List, newTask(t))
	}
	return response
}

func newTask(task tasks.Task) Task {
	return Task{
		ID:          task.ID,
		Summary:     task.Summary,
		PerformedAt: task.PerformedAt,
		CreatedBy:   task.CreatedBy,
	}
}

func newFiltersFromRequest(r *http.Request) (tasks.SearchFilters, error) {
	filters := tasks.SearchFilters{}

	userID, err := getUserID(r)
	if err != nil {
		return tasks.SearchFilters{}, err
	}

	if userID != 0 {
		filters.CreatedBy = &userID
	}

	return filters, nil
}

func getUserID(r *http.Request) (uint, error) {
	strUserID := r.URL.Query().Get(createdByKey)
	if strUserID == "" {
		return 0, nil
	}
	userID, err := strconv.ParseUint(strUserID, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(userID), nil
}
