package interfaces

import (
	"context"
	"time"

	"github.com/otusteam/go/cleancalendar/internal/domain/entities"
)

type EventStorage interface {
	SaveEvent(ctx context.Context, event *entities.Event) error
	GetEventById(ctx context.Context, id string) (*entities.Event, error)
	GetEventsByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) []*entities.Event
}
