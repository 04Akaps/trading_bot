package slack

import (
	"context"
	"fmt"
	"github.com/04Akaps/trading_bot.git/config"
	"github.com/04Akaps/trading_bot.git/types"
	"github.com/slack-go/slack"
	"log"
	"strconv"
	"strings"
)

type SlackClient struct {
	client *slack.Client
	id     string

	volumeTracker map[string]types.VolumeTrend
}

func NewSlackClient(cfg config.Slack) *SlackClient {
	return &SlackClient{
		client:        slack.New(cfg.Token),
		id:            cfg.ChannelID,
		volumeTracker: make(map[string]types.VolumeTrend),
	}
}

func (c *SlackClient) CurrentPriceMessage(mapping map[string]map[string]string) {

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

func (c *SlackClient) VolumeMessage(mapping map[string]map[string]types.VolumeTrend) {
	var blocks []slack.Block

	for _, info := range mapping {

		for s, p := range info {
			beforeVolumeDiff := "0"
			beforeQVolumeDiff := "0"
			beforePercentDiff := "0"

			v, ok := c.volumeTracker[s]

			if ok {
				if beforeValue, err := strconv.ParseFloat(v.Volume, 64); err == nil {
					currentValue, _ := strconv.ParseFloat(p.Volume, 64)
					beforeVolumeDiff = fmt.Sprintf("%.2f", currentValue-beforeValue)
				}

				if beforeQVolume, err := strconv.ParseFloat(v.QuoteVolume, 64); err == nil {
					currentQVolume, _ := strconv.ParseFloat(p.QuoteVolume, 64)
					beforeQVolumeDiff = fmt.Sprintf("%.2f", currentQVolume-beforeQVolume)
				}

				if beforePercent, err := strconv.ParseFloat(v.PriceChangePercent, 64); err == nil {
					currentPercent, _ := strconv.ParseFloat(p.PriceChangePercent, 64)
					beforePercentDiff = fmt.Sprintf("%.2f", currentPercent-beforePercent)
				}
			}

			v.Volume = p.Volume
			v.QuoteVolume = p.QuoteVolume
			v.PriceChangePercent = p.PriceChangePercent

			c.volumeTracker[s] = v

			message := fmt.Sprintf(
				"> %s \n"+
					"> *Change*   `%s`       \n"+
					"> *CPercent* `%s` | diff `%s`      \n"+
					"> *HPrice*   `%s`       \n"+
					"> *LPrice*   `%s`       \n"+
					"> *QVolume*  `%s` | diff `%s`     \n"+
					"> *Volume*   `%s` | diff `%s` ",
				s,
				formatFloat(p.PriceChange),
				formatFloat(p.PriceChangePercent),
				formatFloat(beforePercentDiff),
				formatFloat(p.HighPrice),
				formatFloat(p.LowPrice),
				formatFloat(p.QuoteVolume),
				formatFloat(beforeQVolumeDiff),
				formatFloat(p.Volume),
				formatFloat(beforeVolumeDiff),
			)

			blocks = append(blocks, slack.NewSectionBlock(
				slack.NewTextBlockObject("mrkdwn", message, false, false),
				nil, nil,
			))

			blocks = append(blocks, slack.NewDividerBlock())
		}

	}

	_, _, err := c.client.PostMessageContext(
		context.Background(),
		c.id,
		slack.MsgOptionAsUser(true),
		slack.MsgOptionBlocks(blocks...), // 블록들을 슬랙에 전송
	)

	if err != nil {
		log.Println("Failed to send slack message", "volumeMessage", "err", err)
	}

}

func (c *SlackClient) VolumeTracker(symbol string, avg, current, diff float64) {
	var message strings.Builder

	message.WriteString(fmt.Sprintf("> *Total Avg Volume*  : `%.2f`\n", avg))
	message.WriteString(fmt.Sprintf("> *Current Volume*    : `%.2f`\n", current))
	message.WriteString(fmt.Sprintf("> *Volume Diff*       : `%.2f`\n", diff))

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

func (c *SlackClient) Top5VolumeDiffTrend(result []types.Top5VolumeDiff) {

	fields := make([]slack.AttachmentField, len(result))

	for i, info := range result {
		var message strings.Builder

		message.WriteString(fmt.Sprintf("> *Today Volume*     : `%.2f`\n", info.CurrentVolume))
		message.WriteString(fmt.Sprintf("> *Before Volume*    : `%.2f`\n", info.BeforeVolume))
		message.WriteString(fmt.Sprintf("> *Diff*             : `%.2f`\n", info.Diff))

		att := slack.AttachmentField{
			Title: fmt.Sprintf("*%s*", info.Symbol),
			Value: message.String(),
		}

		fields[i] = att
	}

	attachment := slack.Attachment{
		Title:  "----------- Current Volume Diff Top5 ----------- ",
		Fields: fields,
	}

	_, _, err := c.client.PostMessageContext(
		context.Background(),
		c.id,
		slack.MsgOptionAsUser(true),
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		log.Println("Failed to send slack message", "Top5VolumeDiffTrend", "err", err)
	}
}

func (c *SlackClient) Top5VolumeDiffStarter() {
	var message strings.Builder
	message.WriteString("> 어제, 오늘을 비교하여 거래량 차이가 큰 Top5를 집계 시작합니다. \n")
	message.WriteString("데이터가 많기 떄문에 기다려주세요.")

	att := slack.AttachmentField{Title: message.String()}

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
		log.Println("Failed to send slack message", "Top5VolumeDiffStarter", "err", err)
	}
}

func formatFloat(value string) string {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return value
	}

	return fmt.Sprintf("%.2f", floatValue)
}
