package concurrent

import (
	"context"
	"errors"
	"runtime"
	"time"
)

type TaskHandle struct {
	ctx    context.Context
	cancel context.CancelFunc

	fn   func(ctx context.Context) error
	done chan error
}

func (t *TaskHandle) Cancel() {
	if t.cancel != nil {
		t.cancel()
	}
}

func (t *TaskHandle) Wait() error {
	return <-t.done
}

func (t *TaskHandle) WaitTimeout(timeout time.Duration) error {
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case <-timer.C:
		return errors.New("context deadline exceeded")
	case err := <-t.done:
		return err
	}
}

type Pool struct {
	// TODO: 使用无锁数据结构来提高性能
	taskCh    chan *TaskHandle
	workerNum int
	done      chan struct{}
}

func (p *Pool) poll() {
	for i := 0; i < p.workerNum; i++ {
		go func() {
			for {
				select {
				case <-p.done:
					return
				case task := <-p.taskCh:
					select {
					case <-task.ctx.Done():
						task.done <- task.ctx.Err()
					default:
						task.done <- task.fn(task.ctx)
					}
				}
			}
		}()
	}
}

func (p *Pool) JobQueueSize() int {
	return len(p.taskCh)
}

func (p *Pool) Spawn(ctx context.Context, fn func(ctx context.Context) error) *TaskHandle {
	ctx, cancel := context.WithCancel(ctx)
	t := &TaskHandle{
		ctx:    ctx,
		cancel: cancel,
		fn:     fn,
		done:   make(chan error, 1),
	}
	p.taskCh <- t
	return t
}

func (p *Pool) Shutdown() {
	close(p.done)
}

func NewPool(workerNum int, taskNum int) *Pool {
	pool := &Pool{
		taskCh:    make(chan *TaskHandle, taskNum),
		workerNum: workerNum,
		done:      make(chan struct{}),
	}
	pool.poll()
	return pool
}

var globalPool = NewPool(runtime.NumCPU(), 10000)

func Spawn(ctx context.Context, fn func(ctx context.Context) error) *TaskHandle {
	return globalPool.Spawn(ctx, fn)
}
