package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)
	wg.Add(1)
	go dealLongWithCtx(wg, ctx)
	wg.Wait()
}

func dealLongWithCtx(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randVal := r.Intn(2000)
	randTime := time.Duration(randVal) * time.Millisecond
	timer := time.NewTimer(randTime)
	fmt.Printf("wait for %s \n", randTime)
	select {
	case <-timer.C:
		fmt.Println("Done")
	case <-ctx.Done():
		fmt.Println("Canceled")
	}
}
