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

	time.Sleep(100 * time.Millisecond)
	require.Eventually(t, func() bool {
		return 1 == atomic.LoadInt64(&i)
	}, time.Second, 10*time.Millisecond)

	mock.Add(time.Nanosecond)

	require.Eventually(t, func() bool {
		select {
		case <-doneCh:
			return true
		default:
			return false
		}
	}, time.Second, 10*time.Millisecond)

	require.EqualValues(t, 2, i)
}
