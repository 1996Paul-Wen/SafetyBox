package ratelimit

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestLimit(t *testing.T) {
	l := LoadLimiter("test")
	now := time.Now()
	wg := &sync.WaitGroup{}
	var failedCount int64
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			// test allow
			pass := l.AllowN(1)
			if !pass {
				atomic.AddInt64(&failedCount, 1)
				fmt.Printf("index %d failed", index)
				return
			}

			// // test wait
			// err := l.Wait(context.Background())
			// if err != nil {
			//  atomic.AddInt64(&failedCount, 1)
			// 	fmt.Printf("index %d failed", index)
			//  return
			// }
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(now).Seconds())
	fmt.Println("failed: ", failedCount)
}
