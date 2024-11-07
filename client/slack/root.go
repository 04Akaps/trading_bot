package slack

import (
	"context"
	"fmt"
	"github.com/04Akaps/trading_bot.git/config"
	"github.com/slack-go/slack"
	"log"
)

type SlackClient struct {
	client *slack.Client
	id     string
}

func NewSlackClient(cfg config.Slack) SlackClient {
	var client SlackClient

	client.client = slack.New(
		cfg.Token,
		slack.OptionDebug(true),
	)
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
		log.Println("Failed to send slack message", "currencPrice", "err", err)
	}
}

func (c SlackClient) TestMessage(format string) {
	_, _, err := c.client.PostMessageContext(
		context.Background(),
		c.id,
		slack.MsgOptionText(format, false),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		fmt.Println("errrr", err)
	}
}

func (c SlackClient) HealthCheck() {

	_, _, err := c.client.PostMessageContext(
		context.Background(),
		c.id,
		slack.MsgOptionAttachments(defaultAttachment("Healch Check")),
	)

	if err != nil {
		fmt.Println("errrr", err)
		// TODO logger
	}

}

func defaultAttachment(format string) slack.Attachment {

	attachment := slack.Attachment{
		Fields: []slack.AttachmentField{
			{
				Title: "trading_bot_by_golang",
				Value: format,
			},
		},
	}

	return attachment
}
