package education

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/education/answer"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type EducationCommander struct {
	bot             *tgbotapi.BotAPI
	answerCommander Commander
}

func NewEducationCommander(bot *tgbotapi.BotAPI) EducationCommander {
	return EducationCommander{bot: bot, answerCommander: answer.NewAnswerCommander(bot)}
}

func (c *EducationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "answer":
		c.answerCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("AnswerCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *EducationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "answer":
		c.answerCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("AnswerCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
