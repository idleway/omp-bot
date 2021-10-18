package answer

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *AnswerCommander) Default(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("you wrote: %s", inputMessage.Text))
	_, err := c.bot.Send(msg)
	return err
}
