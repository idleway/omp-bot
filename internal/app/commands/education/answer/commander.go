package answer

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education/answer"
	"log"
)

type AnswerCommander struct {
	bot           *tgbotapi.BotAPI
	answerService *answer.DummyAnswerService
}

func NewAnswerCommander(bot *tgbotapi.BotAPI) *AnswerCommander {
	return &AnswerCommander{bot: bot, answerService: answer.NewDummyAnswerService()}
}

func (c *AnswerCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var err error
	switch callbackPath.CallbackName {
	case "list":
		err = c.CallbackList(callback, callbackPath)
	default:
		err = errors.New(fmt.Sprintf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName))
	}
	if err != nil {
		log.Printf("callback %s exec fail, err is %v", callbackPath.CallbackName, err)
	}
}

func (c *AnswerCommander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	var err error
	switch commandPath.CommandName {
	case "help":
		err = c.Help(message)
	case "list":
		err = c.List(message)
	case "get":
		err = c.Get(message)
	case "delete":
		err = c.Delete(message)
	case "new":
		err = c.New(message)
	case "edit":
		err = c.Edit(message)
	default:
		err = c.Default(message)
	}
	if err != nil {
		log.Printf("command %s exec fail, err is %v", commandPath.CommandName, err)
	}
}
