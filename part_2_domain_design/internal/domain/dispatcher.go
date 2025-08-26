package domain

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

var ErrInvalidEventType = errors.New("invalid event type")

type EventListener interface {
	Listen(ctx context.Context, event Event) error
}

type EventDispatcher struct {
	listeners map[string][]EventListener
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		listeners: make(map[string][]EventListener),
	}
}

func (d *EventDispatcher) MustSubscribe(instance Event, listeners ...EventListener) {
	for i := range listeners {
		if listeners[i] == nil {
			panic("development mistake: nil listener provided")
		}
	}

	t := mustDetectType(instance)
	d.listeners[t] = append(d.listeners[t], listeners...)
}

func mustDetectType(e Event) string {
	if t := detectEventType(e); t != "" {
		return t
	}

	panic("development mistake: unsupported event type")
}

func detectEventType(e Event) string {
	t := reflect.TypeOf(e)
	if t.Kind() != reflect.Struct {
		return ""
	}

	return t.Name()
}

func (d *EventDispatcher) Dispatch(ctx context.Context, events ...Event) error {
	for i := range events {
		if err := d.handle(ctx, events[i]); err != nil {
			return err
		}
	}

	return nil
}

func (d *EventDispatcher) handle(ctx context.Context, event Event) error {
	t := detectEventType(event)
	for _, h := range d.listeners[t] {
		if err := h.Listen(ctx, event); err != nil {
			return fmt.Errorf("handle event %s: %w", t, err)
		}
	}

	return nil
}
