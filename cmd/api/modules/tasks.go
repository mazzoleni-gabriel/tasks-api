package modules

import (
	"go.uber.org/fx"
	createTaskEntrypoint "tasks-api/src/tasks/entrypoints/createtask"
	deleteTaskEntrypoint "tasks-api/src/tasks/entrypoints/deletetask"
	searchTasksEntrypoint "tasks-api/src/tasks/entrypoints/searchtasks"
	updateTaskEntrypoint "tasks-api/src/tasks/entrypoints/updatetask"
	"tasks-api/src/tasks/repository"
	"tasks-api/src/tasks/usecases/createtask"
	"tasks-api/src/tasks/usecases/deletetask"
	"tasks-api/src/tasks/usecases/searchtasks"
	"tasks-api/src/tasks/usecases/updatetask"
)

var tasksFactories = fx.Provide(
	// data and infrastructure
	repository.NewWriterMySQL,
	repository.NewReaderMySQL,

	// business layer / use cases
	createtask.NewUseCase,
	searchtasks.NewUseCase,
	deletetask.NewUseCase,
	updatetask.NewUseCase,

	// present layer
	createTaskEntrypoint.NewHandler,
	searchTasksEntrypoint.NewHandler,
	deleteTaskEntrypoint.NewHandler,
	updateTaskEntrypoint.NewHandler,
)

var tasksTranslations = fx.Provide(
	func(u createtask.TaskCreator) createTaskEntrypoint.UseCase { return u },
	func(w repository.WriterMySQL) createtask.Writer { return w },

	func(u searchtasks.TaskSearcher) searchTasksEntrypoint.UseCase { return u },
	func(r repository.ReaderMySQL) searchtasks.Reader { return r },

	func(u deletetask.TaskDeleter) deleteTaskEntrypoint.UseCase { return u },
	func(w repository.WriterMySQL) deletetask.Writer { return w },

	func(u updatetask.TaskUpdater) updateTaskEntrypoint.UseCase { return u },
	func(w repository.WriterMySQL) updatetask.Writer { return w },
	func(r repository.ReaderMySQL) updatetask.Reader { return r },
)

var tasksEndpoints = fx.Invoke(
	createTaskEntrypoint.RegisterHandler,
	searchTasksEntrypoint.RegisterHandler,
	deleteTaskEntrypoint.RegisterHandler,
	updateTaskEntrypoint.RegisterHandler,
)

var tasksModule = fx.Options(
	tasksFactories,
	tasksTranslations,
	tasksEndpoints,
)
