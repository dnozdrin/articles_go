package model

import (
	"time"

	"github.com/google/uuid"
)

type CourseCreatedEvent struct {
	courseID   uuid.UUID
	name       string
	occurredAt time.Time
}

func (e CourseCreatedEvent) CourseID() uuid.UUID {
	return e.courseID
}

func (e CourseCreatedEvent) Name() string {
	return e.name
}

func (e CourseCreatedEvent) OccurredAt() time.Time {
	return e.occurredAt
}

type CourseRenamedEvent struct {
	courseID     uuid.UUID
	previousName string
	newName      string
	occurredAt   time.Time
}

func (e CourseRenamedEvent) CourseID() uuid.UUID {
	return e.courseID
}

func (e CourseRenamedEvent) PreviousName() string {
	return e.previousName
}

func (e CourseRenamedEvent) NewName() string {
	return e.newName
}

func (e CourseRenamedEvent) OccurredAt() time.Time {
	return e.occurredAt
}

type ChapterRenamedEvent struct {
	courseID     uuid.UUID
	chapterID    uuid.UUID
	previousName string
	newName      string
	occurredAt   time.Time
}

func (e ChapterRenamedEvent) CourseID() uuid.UUID {
	return e.courseID
}

func (e ChapterRenamedEvent) ChapterID() uuid.UUID {
	return e.chapterID
}

func (e ChapterRenamedEvent) PreviousName() string {
	return e.previousName
}

func (e ChapterRenamedEvent) NewName() string {
	return e.newName
}

func (e ChapterRenamedEvent) OccurredAt() time.Time {
	return e.occurredAt
}
