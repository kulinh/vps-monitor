package main

import (
	"fmt"
	"net"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	botToken      = "YOUR_BOT_TOKEN"
	chatID        = YOUR_CHAT_ID
	checkInterval = 5 * time.Minute
)

type Server struct {
	Name string
	IP   string
}

var servers = []Server{
	{"SGP-01", "209.17.118.64"},
	{"GCE", "34.168.229.209"},
	{"EC2", "35.160.230.166"},
}

func main() {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}

	for {
		checkServers(bot)
		time.Sleep(checkInterval)
	}
}

func checkServers(bot *tgbotapi.BotAPI) {
	for _, server := range servers {
		if !isServerReachable(server.IP) {
			message := fmt.Sprintf("%s (%s) đang không thể truy cập.", server.Name, server.IP)
			sendTelegramMessage(bot, message)
		}
	}
}

func isServerReachable(ip string) bool {
	conn, err := net.DialTimeout("tcp", ip+":80", 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func sendTelegramMessage(bot *tgbotapi.BotAPI, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := bot.Send(msg)
	if err != nil {
		fmt.Println("Lỗi khi gửi tin nhắn Telegram:", err)
	}
}
