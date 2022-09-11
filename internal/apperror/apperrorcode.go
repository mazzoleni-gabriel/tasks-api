package apperror

type ErrCode string

const (
	ErrOtherUserTask ErrCode = "other_user_task"
	ErrTaskNotFound  ErrCode = "task_not_found"
)
