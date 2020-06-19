package main

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	BotToken   = "1283737315:AAHtMv_A071gsgBbfP5Tsiy8QA-GawbjOxs"
	WebhookURL = "https://2908ba66a1b4.ngrok.io"
)

var rss = map[string]string{
	"Habr":    "https://habrahabr.ru/rss/best/",
	"Tproger": "https://tproger.ru/feed/",
}

type RSS struct {
	Items []Item `xml:"channel>item"`
}

type Item struct {
	URL   string `xml:"guid"`
	Title string `xml:"title"`
}

func main() {

	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))

	updates := bot.ListenForWebhook("/")

	go http.ListenAndServe(":8080", nil)
	fmt.Println("Server is listening")

	for update := range updates {
		if url, ok := rss[update.Message.Text]; ok {
			rss, err := getNews(url)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "sorry, error happened"))
			}
			for _, item := range rss.Items {
				if strings.Contains(item.Title, "IT") {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, item.URL+"\n"+item.Title))
				}
			}
		} else {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "there is only Habr feed availible"))
		}
	}
}

func getNews(url string) (*RSS, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	rss := new(RSS)
	err = xml.Unmarshal(body, rss)
	if err != nil {
		return nil, err
	}
	return rss, nil
}
