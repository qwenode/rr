package main

import (
    "context"
    "log"
    "time"

    "github.com/qwenode/rr"
)

func main() {
    timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
    async := rr.Async(timeout, func(ctx context.Context) error {
        time.Sleep(time.Second * 6)
        return nil
    })
    select {
    case <-timeout.Done():
        cancelFunc()
    case <-async.Done():
        log.Println("执行完成")
    case <-time.After(time.Second * 1):
        log.Println("等待中...")
    }
}
