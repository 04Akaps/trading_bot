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
	slackClient *slack.SlackClient
	exchanger   cryptoCurrency.CryptoCurrency

	volumeTraceInit     atomic.Bool
	volumeUpdateChannel chan string
}

func NewJob(
	slackClient *slack.SlackClient,
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

func (j *Job) Run() {

	j.c.Start()

	j.tracker()

	go j.c.Run()

}

func (j *Job) tracker() {

	if j.volumeTraceInit.Load() {
		j.volumeTrace()
	}

	j.c.AddFunc("0 */15 * * * *", func() {
		j.volumeTrace()
		j.volumeTrend()
	})

}

func (j *Job) exchangerTrend(ctx context.Context) {

}
