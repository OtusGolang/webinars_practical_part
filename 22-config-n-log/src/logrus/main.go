package main

import (
	"github.com/sirupsen/logrus"
	"sync/atomic"
)

func main() {
	var log = logrus.New()
	log.AddHook(&Hook{})
	log.Infoln(errorCounter)
	log.Error("Error")
	log.Infoln(errorCounter)
}

var errorCounter, warningCounter uint64

//Hook for logrus
type Hook struct {
}

//Execute hook
func (hook *Hook) Fire(entry *logrus.Entry) error {
	switch entry.Level {
	case logrus.PanicLevel:
		atomic.AddUint64(&errorCounter, 1)
	case logrus.FatalLevel:
		atomic.AddUint64(&errorCounter, 1)
	case logrus.ErrorLevel:
		atomic.AddUint64(&errorCounter, 1)
	case logrus.WarnLevel:
		atomic.AddUint64(&warningCounter, 1)
	}

	return nil
}

//Return slice of logrus levels with witch hook work
func (hook *Hook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	}
}
