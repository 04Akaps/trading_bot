package slack

import (
	"context"
	"fmt"
	"github.com/04Akaps/trading_bot.git/config"
	"github.com/slack-go/slack"
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

func (c SlackClient) TestMessage(format string) {
	_, _, err := c.client.PostMessageContext(
		context.Background(),
		c.id,
		slack.MsgOptionText(format, false),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		fmt.Println("errrr", err)
		// TODO logger
	}
}

func (c SlackClient) HealthCheck(format string) {

	_, _, err := c.client.PostMessageContext(
		context.Background(),
		c.id,
		slack.MsgOptionAttachments(defaultAttachment(format)),
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
