package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/pjoc-team/locker"
	"github.com/pjoc-team/tracing/logger"

	"go.etcd.io/etcd/clientv3/concurrency"
	"time"
)

type etcdLock struct {
	config *Config
	client *clientv3.Client
}

type closableLock struct {
	session *concurrency.Session
	locker  *concurrency.Mutex
}

// Config etcd配置
type Config struct {
	Endpoints []string
	BasePath  string
}

// NewEtcdLocker 创建etcd锁
func NewEtcdLocker(config *Config) (locker.DistributedLocker, error) {
	client, err := clientv3.New(
		clientv3.Config{
			Endpoints:   config.Endpoints,
			DialTimeout: 3 * time.Second,
		},
	)
	if err != nil {
		return nil, err
	}
	l := &etcdLock{
		config: config,
		client: client,
	}
	return l, nil
}

func (s *etcdLock) GetLocker(ctx context.Context, key string) (locker.Locker, error) {
	session, err := concurrency.NewSession(s.client, concurrency.WithContext(ctx))
	if err != nil {
		log := logger.ContextLog(ctx)
		log.Errorf("failed to create session, err: %v", err.Error())
		return nil, err
	}
	lockerPath := s.buildLockerPath(key)
	lck := concurrency.NewMutex(session, lockerPath)
	l := &closableLock{
		session: session,
		locker:  lck,
	}
	return l, nil
}

func (s *etcdLock) buildLockerPath(key string) string {
	path := fmt.Sprintf("%v/%v", s.config.BasePath, key)
	return path
}

func (c *closableLock) Lock(ctx context.Context) error {
	return c.locker.Lock(ctx)
}

func (c *closableLock) Unlock(ctx context.Context) error {
	return c.locker.Unlock(ctx)
}

func (c *closableLock) Close(ctx context.Context) error {
	if c.session != nil {
		err := c.session.Close()
		if err != nil {
			log := logger.ContextLog(ctx)
			log.Errorf("failed to close session, err: %v", err.Error())
		}
		return err
	}
	return nil
}
