package modules

import (
	"go.uber.org/fx"
	"tasks-api/internal/config"
)

func NewApp() *fx.App {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	return newAppWihConfig(cfg)
}

func newAppWihConfig(cfg config.Configuration) *fx.App {
	options := []fx.Option{
		fx.Provide(func() config.Configuration { return cfg }),
		internalModule,
		tasksModule,
	}

	return fx.New(
		fx.Options(options...),
	)
}
