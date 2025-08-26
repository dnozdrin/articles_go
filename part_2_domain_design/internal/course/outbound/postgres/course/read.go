package course

import (
	"context"

	"github.com/google/uuid"

	"examples/internal/course/model"
)

func (repo *Repository) GetByID(ctx context.Context, id uuid.UUID) (model.Course, error) {
	panic("implement me")
}
