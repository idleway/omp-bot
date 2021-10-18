package answer

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (c *AnswerCommander) Get(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "")

	args := inputMessage.CommandArguments()
	ID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		msg.Text = "wrong args, need answer id"
		_, err = c.bot.Send(msg)
		return err
	}

	answer, err := c.answerService.Describe(ID)
	if err != nil {
		msg.Text = "fail to get answer by id"
	} else {
		msg.Text = answer.String()
	}

	_, err = c.bot.Send(msg)
	return err
}
