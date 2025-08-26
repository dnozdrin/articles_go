package model

import (
	"fmt"

	"github.com/google/uuid"

	"examples/internal/course/vo"
)

type CourseOption func(*Course) error

func NewCourse(id uuid.UUID, opts ...CourseOption) (Course, error) {
	var course Course
	course.SetID(id)

	for _, opt := range opts {
		if err := opt(&course); err != nil {
			return Course{}, fmt.Errorf("new course: %w", err)
		}
	}

	return course, nil
}

func WithName(name string) CourseOption {
	return func(c *Course) error {
		// perform validations

		c.name = name
		return nil
	}
}

func WithComplexity(complexity vo.Complexity) CourseOption {
	return func(c *Course) error {
		// perform validations

		if !complexity.IsDefined() {
			return vo.ErrUndefinedComplexityLevels
		}

		c.complexity = complexity
		return nil
	}
}

func WithChapters(chapters []Chapter) CourseOption {
	return func(c *Course) error {
		// perform validations

		c.chapters = chapters
		return nil
	}
}

func NewChapter(id uuid.UUID, opts ...ChapterOption) (Chapter, error) {
	var chapter Chapter
	chapter.SetID(id)
	for _, opt := range opts {
		if err := opt(&chapter); err != nil {
			return Chapter{}, fmt.Errorf("new chapter: %w", err)
		}
	}

	return chapter, nil
}

type ChapterOption func(*Chapter) error

func WithChapterName(name string) ChapterOption {
	return func(c *Chapter) error {
		// perform validations

		if name == "" {
			return ErrEmptyName
		}

		c.name = name
		return nil
	}
}
