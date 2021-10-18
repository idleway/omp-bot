package answer

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *AnswerCommander) List(inputMessage *tgbotapi.Message) error {
	serializedData, _ := json.Marshal(CallbackListData{Cursor: 0})
	callbackPath := path.CallbackPath{
		Domain:       "education",
		Subdomain:    "answer",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}
	return c.CallbackList(&tgbotapi.CallbackQuery{Message: inputMessage}, callbackPath)
}
