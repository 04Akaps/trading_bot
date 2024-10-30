package job

import (
	"context"
	"github.com/04Akaps/trading_bot.git/client/cryptoCurrency"
	"github.com/04Akaps/trading_bot.git/client/slack"
	_cryptoCurrency "github.com/04Akaps/trading_bot.git/types/cryptoCurrency"
	"github.com/robfig/cron"
)

type Job struct {
	c *cron.Cron

	slackClient slack.SlackClient
	exchanger   cryptoCurrency.CryptoCurrency
}

func NewJob(
	slackClient slack.SlackClient,
	exchanger cryptoCurrency.CryptoCurrency,
) *Job {
	j := &Job{
		c:           cron.New(),
		slackClient: slackClient,
		exchanger:   exchanger,
	}

	return j
}

func (j *Job) Run(ctx context.Context) error {

	j.c.AddFunc("0 * * * *", func() {
		j.slackClient.HealthCheck()
	})

	j.exchanger.GetTokenPrice(_cryptoCurrency.Binance, "BTC")

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
