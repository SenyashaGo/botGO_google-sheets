package main

import "botTG/internal/services"

func main() {

	bot, err := services.NewBot()
	if err != nil {
		panic(err)
	}
	bot.Polling()
}
