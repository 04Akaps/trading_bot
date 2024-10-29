package slack

import (
	"context"
	"github.com/04Akaps/trading_bot.git/config"
	"github.com/slack-go/slack"
	"log"
	"os"
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
		slack.OptionLog(log.New(os.Stdout, "slack:api: ", log.Lshortfile|log.LstdFlags)),
	)
	client.id = cfg.ChannelID

	return client
}

//
//attachment := slack.Attachment{
//Pretext: "some pretext",
//Text:    "some text",
//// Uncomment the following part to send a field too
///*
//	Fields: []slack.AttachmentField{
//		slack.AttachmentField{
//			Title: "a",
//			Value: "no",
//		},
//	},
//*/
//}

//slack.MsgOptionAttachments(attachment),

func (c SlackClient) TestMessage(format string) {
	_, _, err := c.client.PostMessageContext(
		context.Background(),
		c.id,
		slack.MsgOptionText(format, false),

		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		// TODO logger
	}
}
