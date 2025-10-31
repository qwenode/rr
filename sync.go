package rr

import (
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
