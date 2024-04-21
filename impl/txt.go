package impl

import (
	"fmt"
	"strings"

	"github.com/xtt28/gamma/data"
)

func intToLetter(i int) rune {
	return rune('a' - 1 + i)
}

func LessonToTextQuiz(lesson data.Lesson) string {
	bob := strings.Builder{}
	bob.WriteString(fmt.Sprintf("Quiz for language %s\n\n", string(lesson.TargetLanguage)))
	for i, v := range lesson.Elements {
		bob.WriteString(fmt.Sprint(i+1) + ") ")
		bob.WriteString(v.Title + "\n\n")
		if v.Description != "" {
			bob.WriteString(v.Description + "\n\n")
		}
		switch v.Type {
		case data.ElementTypeSelOneFromMany, data.ElementTypeSelManyFromMany:
			for j, opt := range v.MCOptions {
				bob.WriteRune(intToLetter(j + 1))
				bob.WriteString(". ")

				bob.WriteString(opt)
				bob.WriteRune('\n')
			}
		case data.ElementTypeMatchPairs:
			for i := 0; i < len(v.Sides[0]); i++ {
				bob.WriteString(strings.Join([]string{v.Sides[0][i], v.Sides[1][i]}, "\t"))
				bob.WriteRune('\n')
			}
		case data.ElementTypeShortAns:
			bob.WriteString("\n")
			bob.WriteString("____________________________");
		}
		bob.WriteString("\n")
	}
	
	return bob.String()
}