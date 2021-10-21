package answer

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/education/answer"
)

func (c *AnswerCommander) New(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "")

	args := inputMessage.CommandArguments()
	var parsedData answer.Answer
	err := json.Unmarshal([]byte(args), &parsedData)
	if err != nil {
		msg.Text = fmt.Sprintf("json parse fail, err is %v", err)
		_, err = c.bot.Send(msg)
		return err
	}
	newItemID, err := c.answerService.Create(parsedData)
	if err != nil {
		return err
	}

	msg.Text = fmt.Sprintf("answer created, id: %d", newItemID)
	_, err = c.bot.Send(msg)
	return err
}
