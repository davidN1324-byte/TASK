package cmd

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gopkg.in/telebot.v3"
)

var rootCmd = &cobra.Command{
	Use:   "bot-telegram",
	Short: "Telegram bot Go lang",
	Run: func(cmd *cobra.Command, args []string) {
		runBot()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func runBot() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Download error.env file")
	}

	token := os.Getenv("TELE_TOKEN")
	if token == "" {
		log.Fatal("TELE_TOKEN not found")
	}

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Send("You write: " + c.Text())
	})

	log.Println("Bot connected...")
	bot.Start()
}
