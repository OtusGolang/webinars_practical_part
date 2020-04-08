package entities

import (
	"time"

	"github.com/satori/go.uuid"
)

type Event struct {
	Id        uuid.UUID
	Owner     string
	Title     string
	Text      string
	StartTime *time.Time
	EndTime   *time.Time
}
