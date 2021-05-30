package locker

import (
	"context"
)

// DistributedLocker 分布式锁
type DistributedLocker interface {
	// GetLocker 获取锁
	GetLocker(ctx context.Context, key string) (Locker, error)
}

// Locker 锁
type Locker interface {
	// Lock 加锁
	Lock(ctx context.Context) error

	// Unlock 解锁
	Unlock(ctx context.Context) error

	// Close 关闭锁，释放资源
	Close(ctx context.Context) error
}
