package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func emulateLongOperation(ctx context.Context, id int) {
	randVal := rand.Intn(5000)
	randTime := time.Duration(randVal) * time.Millisecond
	log.Printf("Job %d Will be evalutated for %d", id, randVal)
	timer := time.NewTimer(randTime)

	select {
	case <-timer.C:
		log.Printf("Successfully finished job %d", id)
	case <-ctx.Done():
		log.Printf("id %d timed out", id)
	}
}

func main() {
	wg := sync.WaitGroup{}
	ctx, _ := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			emulateLongOperation(ctx, id)
			//cancel()
			wg.Done()
		}(i)
	}

	wg.Wait()
}
