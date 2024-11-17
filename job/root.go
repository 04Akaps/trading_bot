package job

import (
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

	symbols     map[string]bool
	scanSymbols map[string]bool
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
		symbols:             make(map[string]bool),
		scanSymbols:         make(map[string]bool),
	}

	j.volumeTraceInit.Store(cfg.Info.VolumeTraceInit)
	j.symbols = j.binanceAllSymbols()
	j.scanSymbols = j.getScanSymbols()

	go func() {
		go j.volumeTraceDiffChecker()
	}()

	return j
}

func (j *Job) Run() error {

	j.c.Start()

	if err := j.util(); err != nil {
		return err
	} else if err = j.trend(); err != nil {
		return err
	} else if err = j.tracker(); err != nil {
		return err
	}

	go j.c.Run()

	return nil
}

func (j *Job) tracker() error {

	if j.volumeTraceInit.Load() {
		j.volumeTrace()
	}

	err := j.c.AddFunc("0 */5 * * * *", func() {
		j.volumeTrace()
		j.volumeTrend()
	})

	return err
}

func (j *Job) trend() error {
	err := j.c.AddFunc("0 */15 * * * *", func() {
		j.exchangerTrend()
	})

	return err
}

func (j *Job) util() error {
	err := j.c.AddFunc("0 * * * * *", func() {
		j.symbols = j.binanceAllSymbols()
		j.scanSymbols = j.getScanSymbols()
	})

	return err
}
