package lineAPI

import "github.com/line/line-bot-sdk-go/v7/linebot"

func NewLineBot(channelSecret, channelToken string) *linebot.Client {
	if bot, err := linebot.New(
		channelSecret,
		channelToken,
	); err != nil {
		return nil
	} else {
		return bot
	}
}
