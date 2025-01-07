package main

import (
	"bybjt-trading-bot/internal/transport/rest"
	"bybjt-trading-bot/internal/utils"
)

func main() {
	go rest.StartServer()
	utils.OutputTimer()
}