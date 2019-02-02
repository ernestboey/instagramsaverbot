package telegrambot

import (
	"fmt"

	"github.com/yanyi/instagramsaverbot/internal/app/scraper"
	telebot "gopkg.in/tucnak/telebot.v2"
)

func sendStartMessage(bot *telebot.Bot, m *telebot.Message) {
	msg := `
	Hello 👋! To utilize me, send me a message like: /save https://instagram.com/p/url. I will then and send images back to you for your consumption.
	`
	bot.Send(m.Sender, msg)
	logger.Log(
		"event", "Welcomed user",
		"username", m.Sender.Username,
		"reply", msg,
	)
}

func sendHelloWorld(bot *telebot.Bot, m *telebot.Message) {
	reply := fmt.Sprintf("Hello, %s 👋", m.Sender.FirstName)
	bot.Send(m.Sender, reply)
	logger.Log(
		"event", "Replied user",
		"username", m.Sender.Username,
		"reply", reply,
	)
}

func sendInstagramImage(bot *telebot.Bot, m *telebot.Message) {
	inputURL := m.Payload
	urls := []string{}
	err := scraper.Scrape(inputURL, &urls)
	if err != nil {
		errMsg := fmt.Sprintf("%s. Please try sending an Instagram link that contains a photo. We currently do not support videos 🙇‍♂️.", err.Error())
		bot.Send(m.Sender, errMsg)
		logger.Log(
			"event", "Can't scrape link",
			"error", err,
			"payload", m.Payload,
		)
		return
	}

	for _, url := range urls {
		bot.Send(m.Sender, url, telebot.Silent, telebot.NoPreview)
		logger.Log(
			"event", "Sent Instagram image URL",
			"username", m.Sender.Username,
			"reply", url,
		)
	}

	bot.Send(m.Sender, "Enjoy your links.")
}
