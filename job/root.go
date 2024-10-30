package job

import (
	"context"
	"github.com/04Akaps/trading_bot.git/client/slack"
	"github.com/robfig/cron"
)

type Job struct {
	c *cron.Cron

	slackClient slack.SlackClient
}

func NewJob(
	slackClient slack.SlackClient,
) *Job {
	j := &Job{
		c:           cron.New(),
		slackClient: slackClient,
	}

	return j
}

func (j *Job) Run(ctx context.Context) error {

	j.c.AddFunc("0 * * * *", func() {
		j.slackClient.HealthCheck("Health check")
	})

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
