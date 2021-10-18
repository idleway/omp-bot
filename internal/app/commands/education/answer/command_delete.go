package answer

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (c *AnswerCommander) Delete(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "")

	args := inputMessage.CommandArguments()
	ID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		msg.Text = "wrong args, need answer id"
		_, err = c.bot.Send(msg)
		return err
	}

	removeOK, err := c.answerService.Remove(ID)
	if err != nil {
		msg.Text = "fail to delete answer by id"
	} else {
		if removeOK {
			msg.Text = "answer successfully deleted"
		} else {
			msg.Text = "answer already deleted"
		}
	}

	_, err = c.bot.Send(msg)
	return err
}
