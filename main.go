package main

import (
	"fmt"

	"github.com/xtt28/gamma/data"
	"github.com/xtt28/gamma/impl"
)

func main() {
	fmt.Println(impl.LessonToTextQuiz(data.Lesson{
		TargetLanguage: "es-ES",
		Elements: []data.Element{
			{
				Type:        data.ElementTypeDescription,
				Title:       "A title",
				Description: "A description",
			},
			{
				Type:      data.ElementTypeSelOneFromMany,
				Title:     "One from many",
				MCOptions: []string{"One", "Two", "Three", "Four"},
			},
			{
				Type:      data.ElementTypeSelManyFromMany,
				Title:     "Many from many",
				MCOptions: []string{"One", "Two", "Three", "Four"},
			},
			{
				Type:  data.ElementTypeMatchPairs,
				Title: "Pairs",
				Sides: [][]string{
					{
						"Left 1",
						"Left 2",
						"Left 3",
						"Left 4",
					},
					{
						"Right 1",
						"Right 2",
						"Right 3",
						"Right 4",
					},
				},
			},
			{
				Type: data.ElementTypeShortAns,
				Title: "Short answer",
			},
		},
	}))
}
