package nosafe

import (
	"context"
	"gopkg.in/go-playground/assert.v1"
	"sync"
	"testing"
)

func Test_nosafeLocker_GetLocker(t *testing.T) {
	locker, _ := NewNosafeLocker()
	ctx := context.Background()
	lk, err := locker.GetLocker(ctx, "test")
	if err != nil {
		t.Fatal(err.Error())
	}

	count := 0
	size := 10000
	wg := sync.WaitGroup{}
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := lk.Lock(ctx)
			if err != nil {
				t.Fatal(err.Error())
				return
			}
			defer func() {
				_ = lk.Unlock(ctx)
			}()
			count++
		}()
	}
	wg.Wait()
	assert.Equal(t, size, count)
}
