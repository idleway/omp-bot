package answer

import (
	"fmt"
)

var allEntities = []Answer{
	{ID: 1, QuestionID: 1, TextValue: "one", CorrectScoreValue: 0.54},
	{ID: 2, QuestionID: 1, TextValue: "two", CorrectScoreValue: 0.64},
	{ID: 3, QuestionID: 1, TextValue: "three", CorrectScoreValue: 0.73},
	{ID: 4, QuestionID: 2, TextValue: "four", CorrectScoreValue: 0.36},
	{ID: 5, QuestionID: 2, TextValue: "five", CorrectScoreValue: 0.1},
	{ID: 6, QuestionID: 2, TextValue: "six", CorrectScoreValue: 0.87},
	{ID: 7, QuestionID: 3, TextValue: "seven", CorrectScoreValue: 0.33},
	{ID: 8, QuestionID: 3, TextValue: "eight", CorrectScoreValue: 0.55},
	{ID: 9, QuestionID: 4, TextValue: "nine", CorrectScoreValue: 0.76},
	{ID: 10, QuestionID: 4, TextValue: "ten", CorrectScoreValue: 0.22},
}

type Answer struct {
	ID                uint64
	QuestionID        uint64
	TextValue         string
	CorrectScoreValue float64
}

func (el Answer) String() string {
	return fmt.Sprintf("ID: %d, QuestionID: %d, TextValue: %s, CorrectScoreValue: %.2f", el.ID, el.QuestionID, el.TextValue, el.CorrectScoreValue)
}
