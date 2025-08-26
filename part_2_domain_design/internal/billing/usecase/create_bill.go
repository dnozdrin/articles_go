package usecase

import (
	"context"
)

type CreateBill struct{}

func (c CreateBill) CreateBill(ctx context.Context, courseName string) error {
	panic("implement me")
}
