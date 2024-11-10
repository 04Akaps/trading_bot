package job

import (
	"context"
	"github.com/04Akaps/trading_bot.git/client/cryptoCurrency"
	"github.com/04Akaps/trading_bot.git/client/slack"
	"github.com/04Akaps/trading_bot.git/config"
	"github.com/04Akaps/trading_bot.git/repository/mongoDB"
	"github.com/robfig/cron"
	"sync/atomic"
)

type Job struct {
	c   *cron.Cron
	cfg config.Config

	mongoDB     mongoDB.MongoDB
	slackClient slack.SlackClient
	exchanger   cryptoCurrency.CryptoCurrency

	volumeTraceInit     atomic.Bool
	volumeUpdateChannel chan string
}

func NewJob(
	slackClient slack.SlackClient,
	exchanger cryptoCurrency.CryptoCurrency,
	mongoDB mongoDB.MongoDB,
	cfg config.Config,
) *Job {
	j := &Job{
		c:                   cron.New(),
		cfg:                 cfg,
		slackClient:         slackClient,
		exchanger:           exchanger,
		mongoDB:             mongoDB,
		volumeUpdateChannel: make(chan string),
	}

	j.volumeTraceInit.Store(cfg.Info.VolumeTraceInit)

	go func() {
		go j.volumeTraceDiffChecker()
	}()

	return j
}

func (j *Job) Run(ctx context.Context) error {

	if j.volumeTraceInit.Load() {
		// 만약 초기에 init 설정이라면, true로 설정 후 false로 변경
		j.volumeTrace(context.WithCancel(ctx))
	}

	j.c.AddFunc("*/5 * * * *", func() {
		j.volumeTrace(context.WithCancel(ctx))
	})

	j.c.AddFunc("*/15 * * * *", func() {
		j.volumeTrend(context.WithCancel(ctx))
	})

	go j.c.Run()

	return nil
}
