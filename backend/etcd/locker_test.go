package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/integration"
	"go.etcd.io/etcd/pkg/testutil"
	"gopkg.in/go-playground/assert.v1"
	"sync"
	"testing"
	"time"
)

func Test_etcdLock_GetLocker(t *testing.T) {
	lock, err := NewEtcdLocker(
		&Config{
			Endpoints: endpoints,
			BasePath:  "/test",
		},
	)
	if err != nil {
		t.Fatal(err.Error())
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	count := 0
	size := 100
	wg := sync.WaitGroup{}
	for j := 0; j < size; j++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			locker, err := lock.GetLocker(ctx, fmt.Sprintf("test%d", index))
			if err != nil {
				t.Fatal(err.Error())
			}
			err = locker.Lock(ctx)
			if err != nil {
				t.Fatal(err.Error())
			}
			defer func() {
				_ = locker.Unlock(ctx)
				_ = locker.Close(ctx)
			}()
			count++
		}(0)
	}
	wg.Wait()
	fmt.Println(count)
	assert.Equal(t, count, size)
}

var endpoints []string

// TestMain sets up an etcd cluster for running the examples.
func TestMain(m *testing.M) {
	cfg := integration.ClusterConfig{Size: 1}
	clus := integration.NewClusterV3(nil, &cfg)
	endpoints = []string{clus.Client(0).Endpoints()[0]}
	fmt.Println("endpoints", endpoints)
	v := m.Run()
	fmt.Println("terminating etcd")
	clus.Terminate(nil)
	if err := testutil.CheckAfterTest(2 * time.Second); err != nil {
		// fmt.Fprintf(os.Stderr, "%v", err)
		// os.Exit(1)
	}
	if v == 0 && testutil.CheckLeakedGoroutine() {
		// os.Exit(1)
	}
	// os.Exit(v)
}
