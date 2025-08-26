package course

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"examples/internal/course/model"
	"examples/internal/course/outbound/postgres"
	"examples/internal/domain"
)

type Repository struct {
	manager    *postgres.TransactionManager
	dispatcher *domain.EventDispatcher
}

func NewRepository(manager *postgres.TransactionManager, dispatcher *domain.EventDispatcher) Repository {
	return Repository{
		manager:    manager,
		dispatcher: dispatcher,
	}
}

func (repo *Repository) Save(ctx context.Context, agg model.Course) error {
	events := agg.Events()
	if len(events) == 0 {
		return nil
	}

	return repo.manager.InTransaction(ctx, func(ctx context.Context, tx bun.Tx) error {
		var updatedAt time.Time
		for i := range events {
			switch event := events[i].(type) {
			case model.CourseCreatedEvent:
				if err := insertCourse(ctx, tx, event); err != nil {
					return fmt.Errorf("persist CourseCreatedEvent: %w", err)
				}
			case model.CourseRenamedEvent:
				if err := updateCourseName(ctx, tx, event); err != nil {
					return fmt.Errorf("persist CourseRenamedEvent: %w", err)
				}
			case model.ChapterRenamedEvent:
				if err := updateChapterName(ctx, tx, event); err != nil {
					return fmt.Errorf("persist ChapterRenamedEvent: %w", err)
				}
				// etc
			}

			updatedAt = events[i].OccurredAt()
		}

		if err := incrementVersion(ctx, tx, agg.ID(), agg.Version(), updatedAt); err != nil {
			return fmt.Errorf("increment course version: %w", err)
		}

		if err := repo.dispatcher.Dispatch(ctx, events...); err != nil {
			return fmt.Errorf("dispatch course events: %w", err)
		}

		agg.Reset()
		return nil
	}, nil)
}

func incrementVersion(ctx context.Context, tx bun.Tx, id uuid.UUID, version int64, updatedAt time.Time) error {
	res, err := tx.NewUpdate().
		Table("courses").
		Set("version = version + 1").
		Set("updated_at = ?", updatedAt).
		Where("id = ?", id).
		Where("version = ?", version).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	num, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if num == 0 {
		return fmt.Errorf("result: %w", domain.ErrInvalidAggregateVersion)
	}

	return nil
}

func insertCourse(ctx context.Context, tx bun.Tx, event model.CourseCreatedEvent) error {
	panic("implement me")
}

func updateCourseName(ctx context.Context, tx bun.Tx, event model.CourseRenamedEvent) error {
	panic("implement me")
}

func updateChapterName(ctx context.Context, tx bun.Tx, event model.ChapterRenamedEvent) error {
	panic("implement me")
}
