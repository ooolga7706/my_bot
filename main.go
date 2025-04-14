package main

import (
    "log"
    "os"
    "strings"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
    botToken := os.Getenv("TELEGRAM_TOKEN")
    bot, err := tgbotapi.NewBotAPI(botToken)

    if err != nil {
        log.Panic("Ошибка создания бота:", err)
    }

    bot.Debug = true

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)
    if err != nil {
        log.Panic("Ошибка получения обновлений:", err)
    }

    for update := range updates {
        if update.Message == nil {
            continue
        }

        log.Printf("[%s | %d] %s", update.Message.From.UserName, update.Message.Chat.ID, update.Message.Text)

        if strings.ToLower(update.Message.Text) == "/start" {
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я простой Go-бот.")
            _, err := bot.Send(msg)
            if err != nil {
                log.Println("Ошибка при отправке:", err)
            }
        }
    }
}
