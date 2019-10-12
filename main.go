package main

import (
	"encoding/json"
	"fmt"
	"github.com/nanont/feinschmecker/commands"
	"github.com/nanont/feinschmecker/config"
	"github.com/nanont/feinschmecker/reply"
	"github.com/nanont/feinschmecker/sessions"
	"io/ioutil"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CommandMap map[string]func (*config.Config, *sessions.Session) *reply.Reply

func main() {
	fmt.Println("Started Feinschmecker!")

	configRaw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	conf := config.Config{}
	err = json.Unmarshal(configRaw, &conf)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure workdir is present
	err = os.MkdirAll(conf.Workdir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	sessionMap := sessions.Init(conf.Workdir)

	bot, err := tgbotapi.NewBotAPI(conf.Telegram.Token)
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

		text := strings.TrimSpace(update.Message.Text)

		commandMap := CommandMap{
			"/start": commands.Start,
			"/now": commands.Now,
			"/next": commands.Next,
			"/about": commands.About,
			"/en": commands.En,
			"/de": commands.De,
		}

		var rep *reply.Reply
		commandFunc, ok := commandMap[text]
		if ok {
			rep = commandFunc(&conf, session)
		} else {
			// This is the default reply
			rep = commands.Default(&conf, session)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, rep.Translation(session.Language))
		msg.ParseMode = "Markdown"
		// msg.ReplyToMessageID = update.Message.MessageID

		_, err = bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
