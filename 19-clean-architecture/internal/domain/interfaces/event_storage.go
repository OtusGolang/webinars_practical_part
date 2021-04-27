package interfaces

import (
	"context"
	"time"

	"github.com/OtusGolang/webinars_practical_part/19-clean-architecture/internal/domain/entities"
)

type EventStorage interface {
	SaveEvent(ctx context.Context, event *entities.Event) error
	GetEventByID(ctx context.Context, id string) (*entities.Event, error)
	GetEventsByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) ([]entities.Event, error)
}
