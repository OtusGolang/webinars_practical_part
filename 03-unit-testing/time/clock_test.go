package otime_test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/require"
)

func TestClockNow(t *testing.T) {
	mock := clock.NewMock()

	now := mock.Now()
	time.Sleep(10 * time.Nanosecond)
	require.Equal(t, now, mock.Now())

	mock.Add(time.Hour)
	require.Equal(t, now.Add(time.Hour), mock.Now())
}

func TestClockSleep(t *testing.T) {
	var i int64
	doneCh := make(chan struct{})
	mock := clock.NewMock()

	go func() {
		atomic.AddInt64(&i, 1)
		mock.Sleep(time.Nanosecond)
		atomic.AddInt64(&i, 1)
		close(doneCh)
	}()

	waitFor(t, func() bool {
		return 1 == atomic.LoadInt64(&i)
	})

	mock.Add(time.Nanosecond)

	waitFor(t, func() bool {
		return 2 == atomic.LoadInt64(&i)
	})
}

func waitFor(t *testing.T, fn func() bool) {
	t.Helper()
	timer := time.NewTimer(time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			t.Fatal("wait for second but still not happened")
		default:
			if fn() {
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}
