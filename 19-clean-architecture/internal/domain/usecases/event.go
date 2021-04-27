package usecases

import (
	"context"
	"time"

	"github.com/OtusGolang/webinars_practical_part/19-clean-architecture/internal/domain/entities"
	"github.com/OtusGolang/webinars_practical_part/19-clean-architecture/internal/domain/interfaces"
	uuid "github.com/satori/go.uuid"
)

type EventUseCases struct {
	EventStorage interfaces.EventStorage
}

func (es *EventUseCases) CreateEvent(ctx context.Context, owner, title, text string, startTime *time.Time, endTime *time.Time) (*entities.Event, error) {
	// TODO: persistence, validation
	event := &entities.Event{
		ID:        uuid.NewV4(),
		Owner:     owner,
		Title:     title,
		Text:      text,
		StartTime: startTime,
		EndTime:   endTime,
	}
	err := es.EventStorage.SaveEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (es *EventUseCases) GetEventByID(ctx context.Context, id string) (*entities.Event, error) {
	panic("implement me")
}

func (es *EventUseCases) GetEventsByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) ([]entities.Event, error) {
	panic("implement me")
}
