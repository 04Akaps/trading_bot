package main

import (
	"github.com/04Akaps/trading_bot.git/app"
	"github.com/04Akaps/trading_bot.git/app/dependency"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		dependency.Cfg,
		dependency.Slack,
		dependency.CryptoClient,
		dependency.Job,
		fx.Provide(app.NewTracing),
		fx.Invoke(func(app.App) {}),
	).Run()
}
