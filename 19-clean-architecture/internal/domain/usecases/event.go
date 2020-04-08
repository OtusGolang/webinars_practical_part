package usecases

import (
	"context"
	"time"

	"github.com/satori/go.uuid"

	"github.com/otusteam/go/cleancalendar/internal/domain/entities"
	"github.com/otusteam/go/cleancalendar/internal/domain/interfaces"
)

type EventUsecases struct {
	EventStorage interfaces.EventStorage
}

func (es *EventUsecases) CreateEvent(ctx context.Context, owner, title, text string, startTime *time.Time, endTime *time.Time) (*entities.Event, error) {
	// TODO: persistence, validation
	event := &entities.Event{
		Id:        uuid.NewV4(),
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
