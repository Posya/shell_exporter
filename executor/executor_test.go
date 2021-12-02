package executor

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	counter int32 = 0
)

func TestExecutor(t *testing.T) {

	want := "RESULT"

	f := func() string {
		time.Sleep(time.Millisecond * 50)
		atomic.AddInt32(&counter, 1)
		return want
	}

	e := NewExecutor(f)
	ctx := context.Background()
	wg := sync.WaitGroup{}

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			res := e(ctx)
			if res != want {
				t.Errorf("Run() = %v, want %v", res, want)
			}
			wg.Done()
		}()
		time.Sleep(time.Millisecond * 5)
	}
	wg.Wait()

	if counter > 120 {
		t.Errorf("function runs too many times (%d)", counter)
	}

}
