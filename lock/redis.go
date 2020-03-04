package lock

import (
	"context"
	"time"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

type redisLocker struct {
	locker *redsync.Redsync
}

func (r *redisLocker) Do(ctx context.Context, name string, duration time.Duration, fn func(ctx context.Context) error) error {
	m := r.locker.NewMutex(name, redsync.SetExpiry(duration), redsync.SetTries(2), redsync.SetRetryDelayFunc(func(tries int) time.Duration {
		return time.Millisecond * time.Duration(tries) * 10
	}))
	if err := m.Lock(); err != nil {
		return err
	}
	defer m.Unlock()
	return fn(ctx)
}

func NewRedisLocker(address string, password string, db int) Locker {
	pool := &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", address,
				redis.DialPassword(password),
				redis.DialReadTimeout(time.Millisecond*100),
				redis.DialWriteTimeout(time.Millisecond*100),
				redis.DialDatabase(db),
			)
		},
		MaxIdle:     50,
		MaxActive:   100,
		IdleTimeout: time.Second * 50,
	}
	return &redisLocker{locker: redsync.New([]redsync.Pool{
		pool,
	})}
}
