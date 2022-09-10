package modules

import (
	"go.uber.org/fx"
	createTaskEntrypoint "tasks-api/src/tasks/entrypoints/createtask"
	"tasks-api/src/tasks/repository"
	"tasks-api/src/tasks/usecases/createtask"
)

var tasksFactories = fx.Provide(
	// data and infrastructure
	repository.NewWriterMySQL,

	// business layer / use cases
	createtask.NewUseCase,

	// present layer
	createTaskEntrypoint.NewHandler,
)

var tasksTranslations = fx.Provide(
	func(u createtask.TaskCreator) createTaskEntrypoint.UseCase { return u },
	func(w repository.WriterMySQL) createtask.Writer { return w },
)

var tasksEndpoints = fx.Invoke(
	createTaskEntrypoint.RegisterHandler,
)

var tasksModule = fx.Options(
	tasksFactories,
	tasksTranslations,
	tasksEndpoints,
)
