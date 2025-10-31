package rr

import (
    "context"
    "fmt"
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
