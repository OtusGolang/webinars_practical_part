package otime_test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/cabify/timex"
	"github.com/cabify/timex/timextest"
	"github.com/stretchr/testify/require"
)

var now = time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)

func TestTimexNow(t *testing.T) {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		require.Equal(t, now, timex.Now())

		newNow := now.Add(time.Hour)
		mockedtimex.SetNow(newNow)
		require.Equal(t, newNow, timex.Now())
	})
}

func TestTimexSleep(t *testing.T) {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		var i int64
		doneCh := make(chan struct{})

		go func() {
			atomic.AddInt64(&i, 1)
			timex.Sleep(time.Nanosecond)
			atomic.AddInt64(&i, 1)
			close(doneCh)
		}()

		sleepCall := <-mockedtimex.SleepCalls
		require.EqualValues(t, 1, i)
		require.Equal(t, time.Nanosecond, sleepCall.Duration)

		sleepCall.WakeUp()
		<-doneCh
		require.EqualValues(t, 2, i)
	})
}
