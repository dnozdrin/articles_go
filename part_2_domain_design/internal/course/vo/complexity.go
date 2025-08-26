package vo

import (
	"errors"
)

// Complexity — value object (immutable).
type Complexity struct {
	content Level
	tasks   Level
}

type Level byte

const (
	UndefinedLevel Level = iota
	EntryLevel
	IntermediateLevel
	AdvancedLevel
)

var ErrUndefinedComplexityLevels = errors.New("complexity: levels not defined")

func NewComplexity(content, tasks Level) (Complexity, error) {
	if content == UndefinedLevel || tasks == UndefinedLevel {
		return Complexity{}, ErrUndefinedComplexityLevels
	}

	return Complexity{content: content, tasks: tasks}, nil
}

func (c Complexity) Content() Level {
	return c.content
}

func (c Complexity) Tasks() Level {
	return c.tasks
}

func (c Complexity) IsDefined() bool {
	return c.content != UndefinedLevel && c.tasks != UndefinedLevel
}
