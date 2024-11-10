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

	volumeTraceInit atomic.Bool
}

func NewJob(
	slackClient slack.SlackClient,
	exchanger cryptoCurrency.CryptoCurrency,
	mongoDB mongoDB.MongoDB,
	cfg config.Config,
) *Job {
	j := &Job{
		c:           cron.New(),
		cfg:         cfg,
		slackClient: slackClient,
		exchanger:   exchanger,
		mongoDB:     mongoDB,
	}

	j.volumeTraceInit.Store(cfg.Info.VolumeTraceInit)

	return j
}

func (j *Job) Run(ctx context.Context) error {

	//j.c.AddFunc("0 * * * *", func() {
	//	j.slackClient.HealthCheck()
	//})

	if j.volumeTraceInit.Load() {
		// 만약 초기에 init 설정이라면, true로 설정 후 false로 변경
		j.volumeTrace(context.WithCancel(ctx))
	}
	//
	//j.volumeTrend(context.WithCancel(ctx))
	//j.currentPrice(context.WithCancel(ctx))

	//j.exchanger.GetTokenPrice(_cryptoCurrency.Binance, "BTC")

	//j.c.AddFunc("0 0/2 * 1/1 * ? *", func() {
	//	j.alertSignificantPriceChange(context.WithCancel(ctx))
	//})
	//j.c.AddFunc("*/3 * * * *", func() {
	//	j.fetchAndSendPositionData(context.WithCancel(ctx))
	//})
	//j.c.AddFunc("*/5 * * * *", func() {
	//	j.updatePriceTrend(context.WithCancel(ctx))
	//})
	//j.c.AddFunc("*/5 * * * *", func() {
	//	j.CurrentPrice(context.WithCancel(ctx))
	//})

	go j.c.Run()

	return nil
}
