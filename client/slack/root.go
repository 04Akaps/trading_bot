package slack

import (
	"context"
	"fmt"
	"github.com/04Akaps/trading_bot.git/config"
	"github.com/04Akaps/trading_bot.git/types"
	"github.com/slack-go/slack"
	"log"
	"strings"
)

type SlackClient struct {
	client *slack.Client
	id     string
}

func NewSlackClient(cfg config.Slack) SlackClient {
	var client SlackClient

	client.client = slack.New(cfg.Token)
	client.id = cfg.ChannelID

	return client
}

func (c SlackClient) CurrentPriceMessage(mapping map[string]map[string]string) {

	attachment := slack.Attachment{}
	fields := make([]slack.AttachmentField, len(mapping)+1)
	index := 0

	for key, info := range mapping {
		att := slack.AttachmentField{
			Title: key,
		}

		var message string

		for s, p := range info {
			message += fmt.Sprintf("%s -> %s", s, p)
			message += "\n"
		}

		att.Value = message

		fields[index] = att
		index++
	}

	fields[index] = slack.AttachmentField{
		Title: "▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄\n█░░░░░░░░▀█▄▀▄▀██████░▀█▄▀▄▀██████░\n░░░░░░░░░░░▀█▄█▄███▀░░░ ▀██▄█▄███▀░\n",
	}

	attachment.Fields = fields

	_, _, err := c.client.PostMessageContext(
		context.Background(),
		c.id,
		slack.MsgOptionAsUser(true),
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		log.Println("Failed to send slack message", "currentPrice", "err", err)
	}
}

func (c SlackClient) VolumeMessage(mapping map[string]map[string]types.VolumeTrend) {

	attachment := slack.Attachment{}
	fields := make([]slack.AttachmentField, len(mapping)+1)
	index := 0

	for key, info := range mapping {
		att := slack.AttachmentField{
			Title: key,
		}

		var message strings.Builder

		for s, p := range info {
			att.Title = fmt.Sprintf("🚀🚀 *%s* 🚀🚀", s)
			message.WriteString(fmt.Sprintf("%s -> %s \n", "priceChange", p.PriceChange))
			message.WriteString(fmt.Sprintf("%s -> %s \n", "priceChangePercent", p.PriceChangePercent))
			message.WriteString(fmt.Sprintf("%s -> %s \n", "highPrice", p.HighPrice))
			message.WriteString(fmt.Sprintf("%s -> %s \n", "lowPrice", p.LowPrice))
			message.WriteString(fmt.Sprintf("%s -> %s \n", "openPrice", p.OpenPrice))
			message.WriteString(fmt.Sprintf("%s -> %s \n", "quoteVolume", p.QuoteVolume))
			message.WriteString(fmt.Sprintf("%s -> %s \n", "volume", p.Volume))
		}

		att.Value = message.String()

		fields[index] = att
		index++
	}

	attachment.Fields = fields

	_, _, err := c.client.PostMessageContext(
		context.Background(),
		c.id,
		slack.MsgOptionAsUser(true),
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		log.Println("Failed to send slack message", "volumeMessage", "err", err)
	}

}

func (c SlackClient) VolumeTracker(symbol string, avg, current, diff float64) {
	var message strings.Builder
	message.WriteString(fmt.Sprintf("*%s* -> %.2f \n", "TotalAvgVolume", avg))
	message.WriteString(fmt.Sprintf("*%s* -> %.2f \n", "CurrentVolume", current))
	message.WriteString(fmt.Sprintf("*%s* -> %.2f \n", "VolumeDiff", diff))

	att := slack.AttachmentField{
		Title: fmt.Sprintf("🚀🚀 *%s* 🚀🚀", symbol),
		Value: message.String(),
	}

	attachment := slack.Attachment{
		Fields: []slack.AttachmentField{att},
	}

	_, _, err := c.client.PostMessageContext(
		context.Background(),
		c.id,
		slack.MsgOptionAsUser(true),
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		log.Println("Failed to send slack message", "VolumeTracker", "err", err)
	}
}
