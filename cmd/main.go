package main

import (
	"LineGPT/configs"
	"LineGPT/pkg/controllers"
	"LineGPT/pkg/handlers/httpHandlers"
	"LineGPT/pkg/services/gptAPI"
	"LineGPT/pkg/services/lineAPI"
	"log"
	"net/http"
)

func main() {
	configurationFile := "LINEGPT.yaml"
	cfg, err := configs.ReadEnv(configurationFile)
	if err != nil {
		log.Fatal(err)
	}

	bot := lineAPI.NewLineBot(cfg.Services.LineService.ChannelSecret, cfg.Services.LineService.ChannelToken)
	gptClient := gptAPI.NewGPTClient(cfg.Services.ChatGPTService.Key)
	controller := controllers.NewController(bot, gptClient)
	handler := httpHandlers.NewHandler(controller)
	//
	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", handler.LineHandler)
	if err := http.ListenAndServe(cfg.Server.Port, nil); err != nil {
		log.Fatalln(err)
	}
}
