package domain

import (
	"errors"
	"time"
)

var ErrInvalidAggregateVersion = errors.New("invalid aggregate version")

type Event interface {
	OccurredAt() time.Time
}

type Aggregate struct {
	Entity
	events  []Event
	version int64
}

func (a *Aggregate) AddEvent(event Event) {
	a.events = append(a.events, event)
}

func (a *Aggregate) Events() []Event {
	return a.events
}

func (a *Aggregate) Reset() {
	a.events = nil
}

func (a *Aggregate) SetVersion(version int64) {
	a.version = version
}

func (a *Aggregate) Version() int64 {
	return a.version
}
