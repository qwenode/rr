package main

import (
    "context"
    "errors"
    "log"
    "time"

    "github.com/qwenode/rr"
    "github.com/qwenode/rr/random"
)

func main() {
    timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
    async := rr.Async(timeout, func(ctx context.Context) error {
        log.Println("任务开始")
        time.Sleep(time.Second * 3)
        log.Println("任务结束")
        if random.IntRange(1, 100) > 50 {
            return errors.New("随机错误")
        }
        return nil
    })
    for {
        select {
        case <-timeout.Done():
            cancelFunc()
            log.Println("超时退出")
        case <-async.Done():
            cancelFunc()
            log.Println("任务完成",async.Get())
        case <-time.After(time.Second * 1):
            log.Println("等待中...",async.IsDone())
            continue
        }
        break
    }
    log.Println("启动一个后台任务")
    a := rr.Async(context.Background(), func(ctx context.Context) error {
        log.Println("这是一个后台任务")
        time.Sleep(time.Second *5)
        panic("测试panic处理")
        log.Println("后台任务完成")
        return errors.New("后台任务错误")
    })
    log.Println("等待后台任务完成")
    log.Println("后台任务结果:", a.Get())
}
