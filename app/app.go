package app

import (
	"context"
	"github.com/04Akaps/trading_bot.git/job"
	"go.uber.org/fx"
)

type App struct {
}

func NewTracing(
	lc fx.Lifecycle,
	job *job.Job,
) App {

	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			return job.Run()
		},
	})
	return App{}
}
