package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func IsReplyToBot(message *tgbotapi.Message, botSign string) bool {
	isReply := message.ReplyToMessage

	var isReplyToBot bool

	if isReply != nil {
		replyTo := message.ReplyToMessage.From.UserName
		isReplyToBot = replyTo == botSign
	}

	return isReplyToBot
}
