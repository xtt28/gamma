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
	TargetLanguage Language  `json:"target"`
	Elements       []Element `json:"elements"`
}

// Element is an element present in a lesson for a language, i.e. a description
// or quiz question.
type Element struct {
	Type        ElementType `json:"type"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	MCOptions   []string    `json:"multipleChoiceOptions"`
	Validator   Validator   `json:"validator"`
}

// Validator represents a set of rules that answers to an element of a lesson
// will be checked against.
type Validator struct {
	// MustBeIn is an array of possible answers to the question. The learner's
	// answer must be in this array in order to be considered valid.
	MustBeIn []string `json:"mustBeIn"`
	// MustMatchOne is an array of regular expressions that match possible answers
	// to the question. The learner's answer must match one of these regular
	// expressions in order to be considered valid.
	MustMatchOne []string `json:"mustMatchOne"`
	// PairsMustBe is an array of string pairs used for validating pair matching
	// questions. The pairs created by the learner must match the pairs specified
	// here in order for the answer to be considered valid.
	PairsMustBe [][]string `json:"pairsMustBe"`
}
