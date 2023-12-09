package main

import (
	"flag"
	"log"
	"os"
	"sync"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var (
	telegramBotToken string
	mainMenu = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Start"),
			tgbotapi.NewKeyboardButton("Start"),
			tgbotapi.NewKeyboardButton("Start"),
		),
	)
	)

func init() {
	flag.StringVar(&telegramBotToken, "telegrambottoken", "6890656666:AAFo2uXLPSW0wwDQUMTEHBmU77mB2zR35ls", "Telegram Bot Token")
	flag.Parse()

	if telegramBotToken == "" {
		log.Fatalf("telegram bot token is required")
		os.Exit(1)
	}
}


func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	bot := start_bot(telegramBotToken)
	go login_habitika("NoPantsuMan", "P1zd4l1z")
	go get_updates(bot)
	wg.Wait()
}

func start_bot(token string) tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized with acc %s", bot.Self.UserName)
	return *bot
}
// TODO: нужно сделать так что бы апдейт кончался, что бы  не надо было каждый раз тыкать /menu для обновления кнопок
// TODO: сделать функционал для кнопок, что бы кнопка логин принимала инфу от пользака, юзернейм и пасс
func get_updates(bot tgbotapi.BotAPI){
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	
	for update := range updates {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = "Напиши /start что бы выбрать действие"
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message.IsCommand(){
			switch update.Message.Command() {
			case "start":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Что хочешь сделать?")
				msg.ReplyMarkup = CreateButtons()
			case "menu":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Главное меню")
				msg.ReplyMarkup = mainMenu
			default:
				msg.Text = "Напиши /start что бы выбрать действие"
			}
		}else{
			msg.Text = "Напиши /start что бы выбрать действие"
		}
		bot.Send(msg)
	}
}
