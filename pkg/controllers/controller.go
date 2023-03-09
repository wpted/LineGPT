package controllers

import (
	"LineGPT/pkg/services/gptAPI"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Controller struct {
	Bot       *linebot.Client
	GPTClient *gptAPI.GPTClient
}

// NewController creates a
func NewController(bot *linebot.Client, gptClient *gptAPI.GPTClient) *Controller {
	return &Controller{
		Bot:       bot,
		GPTClient: gptClient,
	}
}
