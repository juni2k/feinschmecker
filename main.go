package main

import (
	"encoding/json"
	"fmt"
	"github.com/nanont/feinschmecker/lang"
	"github.com/nanont/feinschmecker/sessions"
	"io/ioutil"
	"log"
	"os"
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
		Workdir  string `json:"workdir"`
		Telegram struct {
			Token string `json:"token"`
		} `json:"telegram"`
	}

	config := Config{}
	err = json.Unmarshal(configRaw, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure workdir is present
	err = os.MkdirAll(config.Workdir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	sessionMap := sessions.Init(config.Workdir)

	bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		session := sessions.GetOrNew(sessionMap, update.Message.Chat.ID)
		_ = session

		text := update.Message.Text

		reply := "[Reply not specified]"
		// TODO: re-work this into some kind of map / ...
		if strings.HasPrefix(text, "/now") {
			reply = menu.Show(menu.Now, session.Language)
		} else if strings.HasPrefix(text, "/next") {
			reply = menu.Show(menu.Next, session.Language)
		} else if strings.HasPrefix(text, "/en") {
			session.Language = lang.En
			session.Save()
			reply = "Excellent!"
		} else if strings.HasPrefix(text, "/de") {
			session.Language = lang.De
			session.Save()
			reply = "Sehr gut!"
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		msg.ParseMode = "Markdown"
		// msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
