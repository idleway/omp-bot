package answer

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const paginationLimit = 3

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
}

func (c *AnswerCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) error {
	var parsedData CallbackListData
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "")
	answer, err := c.answerService.List(parsedData.Cursor, paginationLimit)
	if err != nil {
		msg.Text = "fail to get answer list"
	} else {
		for i := range answer {
			msg.Text += fmt.Sprintf("%s\n", answer[i].String())
		}
		if msg.Text == "" {
			msg.Text = "answer list is empty"
		}
	}

	//prepare keyboard
	var newData CallbackListData
	if parsedData.Cursor+paginationLimit >= c.answerService.Count() {
		newData = CallbackListData{Cursor: parsedData.Cursor}
	} else {
		newData = CallbackListData{Cursor: parsedData.Cursor + paginationLimit}
	}
	callbackDataButtonNext, _ := json.Marshal(&newData)

	//unsigned type => check overflow
	if parsedData.Cursor < paginationLimit {
		newData = CallbackListData{Cursor: 0}
	} else {
		newData = CallbackListData{Cursor: parsedData.Cursor - paginationLimit}
	}
	callbackDataButtonPrev, _ := json.Marshal(&newData)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Prev page", path.CallbackPath{
				Domain:       "education",
				Subdomain:    "answer",
				CallbackName: "list",
				CallbackData: string(callbackDataButtonPrev),
			}.String()),
			tgbotapi.NewInlineKeyboardButtonData("Next page", path.CallbackPath{
				Domain:       "education",
				Subdomain:    "answer",
				CallbackName: "list",
				CallbackData: string(callbackDataButtonNext),
			}.String()),
		),
	)

	_, err = c.bot.Send(msg)
	return err
}
