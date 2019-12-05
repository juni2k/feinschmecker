package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/nanont/feinschmecker/cache"

	"github.com/nanont/feinschmecker/commands"
	"github.com/nanont/feinschmecker/config"
	"github.com/nanont/feinschmecker/reply"
	"github.com/nanont/feinschmecker/sessions"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var commandMap = commands.Map{
	"/start": commands.Start,
	"/now":   commands.Now,
	"/next":  commands.Next,
	"/about": commands.About,
	"/en":    commands.En,
	"/de":    commands.De,
}

func main() {
	fmt.Println("Started Feinschmecker!")

	configPath := flag.String("c", "./config.json", "Path to configuration file")
	flag.Parse()

	configRaw, err := ioutil.ReadFile(*configPath)
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

		text := strings.TrimSpace(update.Message.Text)

		// Telegram does not remove mentions its own
		text = strings.Replace(text, "@"+bot.Self.UserName, "", -1)

		msg := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			cache.GetOrSet(fmt.Sprintf("%d %s", session.Language, text), func() string {
				var rep *reply.Reply

				commandFunc, ok := commandMap[text]
				if ok {
					rep = &reply.Reply{commandFunc(&conf, session)}
				} else {
					// This is the default reply
					rep = &reply.Reply{commands.Default(&conf, session)}
				}

				return rep.Translation(session.Language)
			}))
		msg.ParseMode = "Markdown"

		_, err = bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
