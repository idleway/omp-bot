package answer

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/education/answer"
)

func (c *AnswerCommander) Edit(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "")

	args := inputMessage.CommandArguments()
	var parsedData answer.Answer
	err := json.Unmarshal([]byte(args), &parsedData)
	if err != nil {
		msg.Text = fmt.Sprintf("json parse fail, err is %v", err)
		_, err = c.bot.Send(msg)
		return err
	}
	err = c.answerService.Update(parsedData)
	if err != nil {
		msg.Text = err.Error()
	} else {
		msg.Text = "edit successful"
	}

	_, err = c.bot.Send(msg)
	return err
}
