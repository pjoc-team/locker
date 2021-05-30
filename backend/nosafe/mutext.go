package nosafe

import (
	"context"
	"github.com/pjoc-team/locker"
	"sync"
)

// nosafeLocker 不安全的锁，该方式仅限本地测试
type nosafeLocker struct {
	lockers sync.Map
}

// NewNosafeLocker create local lock, don't use this for production!!!
func NewNosafeLocker() (locker.DistributedLocker, error) {
	return &nosafeLocker{}, nil
}

// GetLocker 获取锁
func (n *nosafeLocker) GetLocker(ctx context.Context, key string) (locker.Locker, error) {
	mutex, _ := n.lockers.LoadOrStore(key, &mutexLocker{})
	return (mutex).(*mutexLocker), nil
}

type mutexLocker struct {
	sync.Mutex
}

func (m *mutexLocker) Close(ctx context.Context) error {
	return nil
}

// Lock 加锁
func (m *mutexLocker) Lock(ctx context.Context) error {
	m.Mutex.Lock()
	return nil
}

// Unlock 释放锁
func (m *mutexLocker) Unlock(ctx context.Context) error {
	m.Mutex.Unlock()
	return nil
}
