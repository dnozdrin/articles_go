package events

import (
	"examples/internal/billing/usecase"
	"examples/internal/course/model"
	"examples/internal/domain"
	"examples/internal/service/listener"
)

type dummyDependencies struct {
	billCreator usecase.CreateBill
}

func setupEventListeners(dispatcher *domain.EventDispatcher, dd dummyDependencies) {
	dispatcher.MustSubscribe(
		model.CourseCreatedEvent{},
		listener.NewCreateBillOnCreateCourse(dd.billCreator),
	)
}
