package dependency

import (
	"flag"
	"github.com/04Akaps/trading_bot.git/client/cryptoCurrency"
	"github.com/04Akaps/trading_bot.git/client/slack"
	"github.com/04Akaps/trading_bot.git/config"
	"github.com/04Akaps/trading_bot.git/job"
	"go.uber.org/fx"
)

var envPath = flag.String("config", "./config.toml", "config path")

func init() {
	flag.Parse()
}

var Cfg = fx.Module(
	"config",
	fx.Provide(func() config.Config {
		return config.NewCfg(*envPath)
	}),
)

var Slack = fx.Module(
	"slack",
	fx.Provide(func(cfg config.Config) slack.SlackClient {
		return slack.NewSlackClient(cfg.Slack)
	}),
)

var CryptoClient = fx.Module(
	"crypto_client",
	fx.Provide(func(cfg config.Config) cryptoCurrency.CryptoCurrency {
		return cryptoCurrency.NewCryptoCurrency(cfg.CryptoCurrency)
	}),
)

var Job = fx.Module(
	"job",
	fx.Provide(func(cfg config.Config, slack slack.SlackClient) *job.Job {
		return job.NewJob(slack)
	}),
)