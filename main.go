package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/telegram-bot-api.v4"
)

func initFlags() {
	pflag.StringP("token", "t", "", "Telegram bot token.")
	viper.BindEnv("TOKEN")

	pflag.Int64("chatid", 0, "Telegram chatID for your bot + user.")
	viper.BindEnv("CHATID")

	pflag.StringP("file", "f", "", "path for file to send")
	viper.BindEnv("FILE")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	initFlags()

	token := viper.GetString("token")
	if token == "" {
		fmt.Printf("Error: please provide a token set as an enviroment variable `TOKEN`\n")
		fmt.Printf("You can create a bot by talking to https://telegram.me/BotFather\n")
		os.Exit(1)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	b, err := bot.GetMe()
	if err != nil {
		fmt.Printf("Fatal error: unexpected issue when trying to get bot info. Is the token correct?")
		fmt.Printf(err.Error())
		os.Exit(2)
	}
	fmt.Printf("Bot created with username: %s\n", b.UserName)

	chatID := viper.GetInt64("chatid")
	if chatID == 0 {
		fmt.Printf("No `CHATID` env var or `--chatid` flag provided. Starting up in echo mode.\n")
		fmt.Printf("Please start a chat with https://telegram.me/%s and make note of the chatID, then restart this bot.\n", b.UserName)
		startEcho(bot)
		os.Exit(3)
	}

	filePath := viper.GetString("file")
	if filePath == "" {
		fmt.Println("Error: Please provide a file path with the '-f' flag or via the `FILE` env var")
		os.Exit(4)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("Error: File '%s' does not exist\n", filePath)
		os.Exit(5)
	}

	fmt.Printf("Sending %s...\n", filePath)

	msg := tgbotapi.NewDocumentUpload(chatID, filePath)
	_, err = bot.Send(msg)
	if err != nil {
		fmt.Printf("Error sending message: %s", err)
	} else {
		fmt.Println("Success!")
	}
}

// startEcho is a very simple bot that simply prints out the username and chatID of all messages that arrive.
// This is used to get the `CHATID` value for the main bot.
func startEcho(bot *tgbotapi.BotAPI) {
	updates, err := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 60,
		Limit:   0,
	})
	if err != nil {
		fmt.Printf("error opening updates channel: %s\n", err.Error())
	}

	fmt.Printf("Waiting for messages...\n")
	for update := range updates {
		resp := fmt.Sprintf("  - User: %s, chatID: %d", update.Message.From.UserName, update.Message.Chat.ID)
		fmt.Println(resp)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, resp)
		bot.Send(msg)
	}
}
