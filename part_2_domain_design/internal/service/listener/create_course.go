package listener

import (
	"context"
	"fmt"

	"examples/internal/course/model"
	"examples/internal/domain"
)

type BillCreator interface {
	CreateBill(ctx context.Context, courseName string) error
}

type CreateBillOnCreateCourse struct {
	billCreator BillCreator
}

func NewCreateBillOnCreateCourse(billCreator BillCreator) *CreateBillOnCreateCourse {
	return &CreateBillOnCreateCourse{
		billCreator: billCreator,
	}
}

func (l *CreateBillOnCreateCourse) Listen(ctx context.Context, input domain.Event) error {
	event, isValid := input.(model.CourseCreatedEvent)
	if !isValid {
		return fmt.Errorf("CreateBillOnCreateCourse: invalid event type: %w", domain.ErrInvalidEventType)
	}

	if err := l.billCreator.CreateBill(ctx, event.Name()); err != nil {
		return fmt.Errorf("CreateBillOnCreateCourse: create bill: %w", err)
	}

	return nil
}
