package course

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"examples/internal/course/model"
	"examples/internal/course/vo"
)

var ErrBrokenData = errors.New("broken data")

type Course struct {
	bun.BaseModel     `bun:"table:сourses"`
	ID                uuid.UUID  `bun:"id,pk"`
	Name              string     `bun:"name"`
	ContentComplexity vo.Level   `bun:"content_complexity"`
	TasksComplexity   vo.Level   `bun:"tasks_complexity"`
	Chapters          []*Chapter `bun:"rel:has-many,join:id=course_id"`
	Version           int64      `bun:"version"`
}

func (c Course) toAggregate() (model.Course, error) {
	complexity, err := vo.NewComplexity(c.ContentComplexity, c.TasksComplexity)
	if err != nil {
		return model.Course{}, fmt.Errorf("Course.toAggregate: %w", err)
	}

	chapters, err := toChapters(c.Chapters)
	if err != nil {
		return model.Course{}, fmt.Errorf("Course.toAggregate: %w", err)
	}

	course, err := model.NewCourse(
		c.ID,
		model.WithName(c.Name),
		model.WithChapters(chapters),
		model.WithComplexity(complexity),
	)
	if err != nil {
		return model.Course{}, fmt.Errorf("Course.toAggregate: %w", err)
	}

	course.SetVersion(c.Version)

	return course, nil
}

type Chapter struct {
	ID   uuid.UUID `bun:"id,pk"`
	Name string    `bun:"name"`
}

func toChapters(chapters []*Chapter) ([]model.Chapter, error) {
	result := make([]model.Chapter, len(chapters))
	var err error
	for i := range chapters {
		if chapters[i] == nil {
			return nil, fmt.Errorf("toChapters: chapter is nil: %w", ErrBrokenData)
		}

		result[i], err = model.NewChapter(chapters[i].ID, model.WithChapterName(chapters[i].Name))
		if err != nil {
			return nil, fmt.Errorf("toChapters: invalid data: %w", err)
		}
	}

	return result, nil
}
