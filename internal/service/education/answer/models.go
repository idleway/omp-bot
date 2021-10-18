package answer

import "fmt"

var allEntities = []Answer{
	{ID: 1, Title: "one"},
	{ID: 2, Title: "two"},
	{ID: 3, Title: "three"},
	{ID: 4, Title: "four"},
	{ID: 5, Title: "five"},
	{ID: 6, Title: "six"},
	{ID: 7, Title: "seven"},
	{ID: 8, Title: "eight"},
	{ID: 9, Title: "nine"},
	{ID: 10, Title: "ten"},
}

type Answer struct {
	ID    uint64
	Title string
}

func (el Answer) String() string {
	return fmt.Sprintf("ID:%d, Title: %s", el.ID, el.Title)
}
