package templates

import (
	"github.com/matcornic/hermes"
)

type Button struct {
	Color string
	Text  string
	Link  string
}

type Confirmation struct {
	Name         string
	Intros       []string
	Instructions string
	Button
	Outros []string
}

func (c Confirmation) Init() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name:   c.Name,
			Intros: c.Intros,
			Actions: []hermes.Action{
				{
					Instructions: c.Instructions,
					Button: hermes.Button{
						Color: c.Button.Color,
						Text:  c.Button.Text,
						Link:  c.Button.Link,
					},
				},
			},
			Outros: c.Outros,
		},
	}
}
