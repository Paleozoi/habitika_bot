package main

import tgbotapi "github.com/Syfaro/telegram-bot-api"

type User struct{
	Username string
	Password string
}

func CreateButtons() tgbotapi.InlineKeyboardMarkup {
	button1 := tgbotapi.NewInlineKeyboardButtonData("Start", "HI")
	button2 := tgbotapi.NewInlineKeyboardButtonData("Test2", "Hi2")

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(button1, button2),
	)

	return keyboard
}
