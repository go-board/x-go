# Concurrent 并发工具

## Concurrent 并发请求并等待的工具
```go
type Concurrent struct {}
func NewConcurrent(ctx context.Context) (*Concurrent, context.Context)
// 安排一个异步任务并返回
func (c *Concurrent) SpawnContext(ctx context.Context, fn func(ctx context.Context) error)
// 安排一个异步任务并返回
func (c *Concurrent) Spawn(fn func() error)
// 安排一些任务并等待结果返回
func (c *Concurrent) SpawnAndWait(ctx context.Context, fns ...func(context.Context) error) error
// 等待任务的执行结束，和SpawnContext与Spawn结合使用
func (c *Concurrent) Wait(ctx context.Context) error
```

## Pool 协程池
```go
type Task struct {
	ctx    context.Context
	cancel context.CancelFunc

	fn   func(ctx context.Context) error
	done chan error
}
// 取消一个任务，目前仅限于任务还未进行调度之前
func (t *Task) Cancel()
// 同步等待一个任务的完成
func (t *Task) Wait() error
// 带有超时的等待一个任务的完成
func (t *Task) WaitTimeout(timeout time.Duration) error
type Pool struct {
    // TODO: 使用无锁数据结构来提高性能
	taskCh    chan *Task
	workerNum int
	done      chan struct{}
}
// 创建一个协程池，worker数量为workerNum，任务队列大小为taskNum
func NewPool(workerNum int, taskNum int) *Pool
// 当前任务队列大小
func (p *Pool) JobQueueSize() int
// 调度一个任务并返回任务的token，用来进行控制
func (p *Pool) Spawn(ctx context.Context, fn func(ctx context.Context) error) *Task
// 关闭协程池
func (p *Pool) Shutdown()
```
## Semaphore 信号量，并发量
```go
type Semaphore struct {...}
// 创建一个并发度为n的信号量
func NewSemaphore(n int64) *Semaphore
// 获取token，并执行，如果没有可用的token，会等待
func (s *Semaphore) SpawnContext(ctx context.Context, fn func(ctx context.Context) error) error
// 获取token，并执行，如果没有可用的token，会等待
func (s *Semaphore) Spawn(fn func() error) error
// 获取token，并执行，如果没有可用的token，会立即返回
func (s *Semaphore) TrySpawnContext(ctx context.Context, fn func(ctx context.Context) error) error
// 获取token，并执行，如果没有可用的token，会立即返回
func (s *Semaphore) TrySpawn(fn func() error) error
```
