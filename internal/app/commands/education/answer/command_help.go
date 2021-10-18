package answer

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *AnswerCommander) Help(inputMessage *tgbotapi.Message) error {
	message := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__education__answer - help\n"+
			"/list__education__answer - answer list\n"+
			"/get__education__answer {id} - get answer by id\n"+
			"/delete__education__answer {id} - delete answer by id\n",
	)
	_, err := c.bot.Send(message)
	return err
}
