package rr

import (
    "context"
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

// 实现支持错误返回的once,并且执行失败的时候,第二次还会执行,直到成功
type Once struct {
    done uint32
    m    sync.Mutex
}

func (o *Once) Do(f func() error) error {
    if atomic.LoadUint32(&o.done) == 1 {
        return nil
    }
    return o.doSlow(f)
}
func (o *Once) doSlow(callback func() error) error {
    o.m.Lock()
    defer o.m.Unlock()
    var err error
    if o.done == 0 {
        err = callback()
        if err == nil {
            atomic.StoreUint32(&o.done, 1)
        }
    }
    return err
}

type async struct {
    err    error
    doneCh chan struct{}
    once   sync.Once
    done   atomic.Bool
}
type AsyncTask interface {
    Get() error
    IsDone() bool
    Done() <-chan struct{}
    HeartbeatWait(c context.Context, interval time.Duration, onHeartbeat func()) error
}

// 启动一个异步任务
func Async(ctx context.Context, fn func(ctx context.Context) error) AsyncTask {
    task := &async{doneCh: make(chan struct{})}

    go func() {
        defer func() {
            if r := recover(); r != nil {
                task.err = fmt.Errorf("async task panicked: %v", r)
            }
            task.once.Do(func() {
                task.done.Store(true)
                close(task.doneCh)
            })
        }()

        task.err = fn(ctx)
    }()

    return task
}
func (t *async) HeartbeatWait(c context.Context, interval time.Duration, onHeartbeat func()) error {
    ticker := time.NewTicker(interval)
    defer ticker.Stop()
    
    for {
        select {
        case <-c.Done():
            return c.Err()
        case <-t.Done():
            return t.err
        case <-ticker.C:
            if onHeartbeat != nil {
                onHeartbeat()
            }
        }
    }
}

// 阻塞等待任务完成
func (t *async) Get() error {
    <-t.doneCh
    return t.err
}

// 检查任务是否已完成（非阻塞）
func (t *async) IsDone() bool {
    return t.done.Load()
}

func (t *async) Done() <-chan struct{} {
    return t.doneCh
}

// -------------------- AsyncResult[T] --------------------

// 支持有返回值的异步任务
// 使用方式：
//   res := AsyncResult(ctx, func(ctx context.Context) (T, error) { ... })
//   v, err := res.Get()

type asyncResult[T any] struct {
    val    T
    err    error
    doneCh chan struct{}
    once   sync.Once
    done   atomic.Bool
}

// 泛型版本的任务接口
type AsyncResultTask[T any] interface {
    // Get 阻塞等待，返回结果与错误
    Get() (T, error)
    // IsDone 非阻塞检查是否完成
    IsDone() bool
    // Done 返回完成通知通道
    Done() <-chan struct{}
}

// AsyncResult 启动一个带返回值的异步任务
func AsyncResult[T any](ctx context.Context, fn func(ctx context.Context) (T, error)) AsyncResultTask[T] {
    task := &asyncResult[T]{doneCh: make(chan struct{})}

    go func() {
        defer func() {
            if r := recover(); r != nil {
                task.err = fmt.Errorf("async result task panicked: %v", r)
            }
            task.once.Do(func() {
                task.done.Store(true)
                close(task.doneCh)
            })
        }()

        v, err := fn(ctx)
        task.val, task.err = v, err
    }()

    return task
}

// Get 阻塞等待并返回结果
func (t *asyncResult[T]) Get() (T, error) {
    <-t.doneCh
    return t.val, t.err
}

// IsDone 非阻塞检查
func (t *asyncResult[T]) IsDone() bool { return t.done.Load() }

// Done 完成通知
func (t *asyncResult[T]) Done() <-chan struct{} { return t.doneCh }
