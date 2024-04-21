package data

// Language is a representation of a language code according to the ISO 639-1
// and 3166-1 Alpha-2 standards, e.g. en-US.
type Language string

// ElementType represents the type of an element in a language lesson.
type ElementType string

const (
	// ElementTypeDescription is an element that provides information to the
	// learner.
	ElementTypeDescription = "description"
	// ElementTypeSelOneFromMany represents a multiple-choice question with one
	// answer.
	ElementTypeSelOneFromMany = "selectOneFromMany"
	// ElementTypeSelManyFromMany represents a multiple-choice question with more
	// than one answer.
	ElementTypeSelManyFromMany = "selectManyFromMany"
	// ElementTypeMatchPairs represents a question in which the learner must form
	// pairs of words in the source and target languages.
	ElementTypeMatchPairs = "matchPairs"
	// ElementTypeShortAns represents a question that accepts a textual answer.
	ElementTypeShortAns = "shortAnswer"
)

// Lesson is a representation of a lesson for a language.
type Lesson struct {
	TargetLanguage Language `json:"target"`
	Elements       []Element `json:"elements"`
}

// Element is an element present in a lesson for a language, i.e. a description
// or quiz question.
type Element struct {
	Type ElementType `json:"type"`
}