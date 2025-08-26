package domain

import "github.com/google/uuid"

type Entity struct {
	id uuid.UUID
}

func (e *Entity) SetID(id uuid.UUID) {
	if e.id != uuid.Nil {
		panic("development mistake: dirty ID set")
	}

	if id == uuid.Nil {
		panic("development mistake: invalid ID")
	}
}

func (e *Entity) ID() uuid.UUID {
	return e.id
}
