package interfaces

import (
	"context"
	"github.com/otusteam/go/cleancalendar/internal/domain/entities"
	"time"
)

type EventStorage interface {
	SaveEvent(ctx context.Context, event *entities.Event) error
	GetEventById(ctx context.Context, id string) (*entities.Event, error)
	GetEventsByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) []*entities.Event
}
