package executor

import (
	"context"
	"sync"
)

func NewExecutor(f func() string) func(ctx context.Context) string {
	var result *string
	var done chan interface{}
	mu := sync.Mutex{}
	isRunning := false

	return func(ctx context.Context) string {
		mu.Lock()
		if !isRunning {
			isRunning = true
			done = make(chan interface{})
			go func() {
				r := f()
				mu.Lock()
				result = &r
				close(done)
				isRunning = false
				mu.Unlock()
			}()
		}
		mu.Unlock()

		res := ""
		select {
		case <-done:
			res = *result
		case <-ctx.Done():
		}

		return res
	}
}
