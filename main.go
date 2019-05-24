package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/nanont/feinschmecker/menu"
)

func main() {
	fmt.Println("Started Feinschmecker!")

	configRaw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	type Config struct {
		Telegram struct {
			Token string `json:"token"`
		} `json:"telegram"`
	}

	config := Config{}
	err = json.Unmarshal(configRaw, &config)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		text := update.Message.Text

		reply := "[Reply not specified]"
		// TODO: re-work this into some kind of map / ...
		if strings.HasPrefix(text, "/now") {
			reply = menu.Show(menu.Now)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		msg.ParseMode = "Markdown"
		// msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
