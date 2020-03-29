package entities

import (
	"github.com/satori/go.uuid"
	"time"
)

type Event struct {
	Id        uuid.UUID
	Owner     string
	Title     string
	Text      string
	StartTime *time.Time
	EndTime   *time.Time
}
