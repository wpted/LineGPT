package httpHandlers

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
)

func (h *Handler) LineHandler(w http.ResponseWriter, r *http.Request) {
	events, err := h.Controller.Bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				answer, err := h.Controller.GPTClient.GetAnswers(message.Text)
				if err != nil {
					log.Fatalln(err)
				} else {
					if _, err = h.Controller.Bot.ReplyMessage(
						// this times out fast
						event.ReplyToken,
						linebot.NewTextMessage(answer)).Do(); err != nil {
						log.Fatalln(err)
					}
				}
			}
		}
	}
}
