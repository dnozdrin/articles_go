package model

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"examples/internal/course/vo"
	"examples/internal/domain"
)

var (
	ErrEmptyName = errors.New("empty name provided")
	ErrSameName  = errors.New("the same name provided")
	ErrNameTaken = errors.New("the name is already in use in this course")
)

type Course struct {
	domain.Aggregate
	name       string
	complexity vo.Complexity
	chapters   []Chapter
}

func (c *Course) Name() string {
	return c.name
}

func (c *Course) Complexity() vo.Complexity {
	return c.complexity
}

func (c *Course) Chapters() []Chapter {
	return c.chapters
}

func (c *Course) Rename(newName string, now time.Time) error {
	if newName == "" {
		return ErrEmptyName
	}

	if c.name == newName {
		return ErrSameName
	}

	c.AddEvent(CourseRenamedEvent{
		courseID:     c.ID(),
		previousName: c.name,
		newName:      newName,
		occurredAt:   now,
	})

	c.name = newName
	return nil
}

func (c *Course) RenameChapter(chapterID uuid.UUID, newName string, now time.Time) error {
	if newName == "" {
		return ErrEmptyName
	}

	if c.isChapterNameUsed(newName) {
		return ErrNameTaken
	}

	for i := range c.chapters {
		if c.chapters[i].ID() == chapterID {
			return c.renameChapterByIdx(i, newName, now)
		}
	}

	return nil
}

func (c *Course) isChapterNameUsed(name string) bool {
	for i := range c.chapters {
		if c.chapters[i].Name() == name {
			return true
		}
	}

	return false
}

func (c *Course) renameChapterByIdx(idx int, newName string, now time.Time) error {
	previousName := c.chapters[idx].Name()
	if err := c.chapters[idx].rename(newName); err != nil {
		return err
	}

	c.AddEvent(ChapterRenamedEvent{
		courseID:     c.ID(),
		chapterID:    c.chapters[idx].ID(),
		previousName: previousName,
		newName:      newName,
		occurredAt:   now,
	})

	return nil
}

type Chapter struct {
	domain.Entity
	name string
}

func (c *Chapter) Name() string {
	return c.name
}

func (c *Chapter) rename(newName string) error {
	if c.name == newName {
		return ErrSameName
	}

	c.name = newName
	return nil
}
