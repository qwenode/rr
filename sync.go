package rr

import (
    "context"
    "sync"
    "sync/atomic"
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

type asyncResult[T any] struct {
    mu     sync.Mutex
    done   bool
    result T
    err    error
    wg     sync.WaitGroup
}

// 启动一个异步任务
func AsyncWithResult[T any](ctx context.Context, fn func(ctx context.Context) (T, error)) *asyncResult[T] {
    task := &asyncResult[T]{}
    task.wg.Add(1)
	
    go func() {
        defer task.wg.Done()

        result, err := fn(ctx)

        task.mu.Lock()
        task.result = result
        task.err = err
        task.done = true
        task.mu.Unlock()
    }()

    return task
}

// 阻塞等待任务完成并获取结果
func (t *asyncResult[T]) Get() (T, error) {
    t.wg.Wait()
    t.mu.Lock()
    defer t.mu.Unlock()
    return t.result, t.err
}

// 检查任务是否已完成（非阻塞）
func (t *asyncResult[T]) IsDone() bool {
    t.mu.Lock()
    defer t.mu.Unlock()
    return t.done
}
