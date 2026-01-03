package main

import (
	"log/slog"
	"strings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() { logger = slog.Default() }

func botStart() {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil { logger.Error("Error when declaring a telegram bot API object:", slog.Any("err", err)) }
	logger.Info("Bot authorized:", slog.Any("username", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates { //bot cycle
		if update.Message != nil {
			command := update.Message.Command() //looks - /msg, we get msg
			args := strings.Fields(update.Message.CommandArguments()) //split string to array 
			logger.Info("Message from:", slog.Any("user", update.Message.From.UserName), slog.Any("command", command), slog.Any("args", args))
			err, output := pluginHandler(&update.Message.From.UserName, &command, &args)
			if err {
				logger.Error("Error executing plugin")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error executing plugin")
				bot.Send(msg)
			} else {
				logger.Info("Plugin complete")
				if len(output) == 0 { //if there is no echo
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Successful executing plugin") 
					bot.Send(msg)
					logger.Info("Report sent")
				} else {
					msg := outputHandler(&output, &update.Message.Chat.ID)
					bot.Send(msg)
					logger.Info("Report sent")
				}
			}
		}
	}
}

func outputHandler(rowOutput *[]byte, chatID *int64) tgbotapi.Chattable {
	str := strings.TrimSpace(string(*rowOutput)) //remove /r/n
	args := strings.Fields(str) //split the string by spaces
	if len(args) > 1 {
		output := strings.Join(args[:len(args)-2], " ") //remove key-word and path
		switch args[len(args)-2] { //check for key word
		case "vid":
			msg := tgbotapi.NewVideo(*chatID, tgbotapi.FilePath(args[len(args)-1]))
			msg.Caption = string(output)
			msg.ParseMode = "HTML"
			return msg
		case "img":
			msg := tgbotapi.NewPhoto(*chatID, tgbotapi.FilePath(args[len(args)-1]))
			msg.Caption = string(output)
			msg.ParseMode = "HTML"
			return msg
		case "doc": //also .exe, .jar, .zip, etc.
			msg := tgbotapi.NewDocument(*chatID, tgbotapi.FilePath(args[len(args)-1]))
			msg.Caption = string(output)
			msg.ParseMode = "HTML"
			return msg
		case "aud":
			msg := tgbotapi.NewAudio(*chatID, tgbotapi.FilePath(args[len(args)-1]))
			msg.Caption = string(output)
			msg.ParseMode = "HTML"
			return msg
		default:
			msg := tgbotapi.NewMessage(*chatID, str)
			return msg
		}
    } else { return tgbotapi.NewMessage(*chatID, str) }
}