package main

import (
    "context"
    "errors"
    "flag"
    "log"
    "time"

    "github.com/qwenode/rr"
    "github.com/qwenode/rr/random"
)

// 本示例演示 rr.Async 与 rr.AsyncResult 的常见用法：
// 1) Async：启动无返回值但可能失败的异步任务，支持 Done()/IsDone()/Get() 等待与查询。
// 2) AsyncResult[T]：启动有返回值的异步任务，通过 Get() 获取 (T, error)。
// 3) 建议在任务函数内自行遵循 ctx（如 select<-ctx.Done()）以支持取消/超时。
// 4) 任务中 panic 会被捕获为 error 返回，避免直接崩溃。
//
// 使用 go:generate 单独执行每个示例：
//   - 在本目录下执行：go generate
//   - 或运行某个示例：go run . -demo=1（1~7）

// 示例1：Async + 超时控制 + 轮询 IsDone
//go:generate go run . -demo=1
// 示例2：Async + panic 处理（panic 将被转换为 error 返回）
//go:generate go run . -demo=2
// 示例3：Async + 主动取消（任务内遵循 ctx）
//go:generate go run . -demo=3
// 示例4：AsyncResult 基本用法（获取返回值）
//go:generate go run . -demo=4
// 示例5：AsyncResult + Done + 超时/取消（任务内支持 ctx）
//go:generate go run . -demo=5
// 示例6：AsyncResult 扇出/聚合（并发多个任务再汇总）
//go:generate go run . -demo=6
// 示例7：HeartbeatWait 心跳等待（等待任务完成时定期执行心跳回调）
//go:generate go run . -demo=7
func main() {
    demo := flag.Int("demo", 0, "选择要运行的示例(1~7); 0 或不传运行全部")
    flag.Parse()

    if *demo == 0 {
        demo1()
        demo2()
        demo3()
        demo4()
        demo5()
        demo6()
        demo7()
        return
    }

    switch *demo {
    case 1:
        demo1()
    case 2:
        demo2()
    case 3:
        demo3()
    case 4:
        demo4()
    case 5:
        demo5()
    case 6:
        demo6()
    case 7:
        demo7()
    default:
        log.Println("未知 demo 编号:", *demo)
    }
}

// 示例1：Async + 超时控制 + 轮询 IsDone
func demo1() {
    log.Println("示例1：Async + 超时控制 + 轮询 IsDone")
    timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
    defer cancelFunc()

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
            log.Println("超时退出")
        case <-async.Done():
            log.Println("任务完成", async.Get())
        case <-time.After(time.Second * 1):
            log.Println("等待中...", async.IsDone())
            continue
        }
        break
    }
}

// 示例2：Async + panic 处理（panic 将被转换为 error 返回）
func demo2() {
    log.Println("示例2：Async + panic 处理（panic 将被转换为 error 返回）")
    a := rr.Async(context.Background(), func(ctx context.Context) error {
        log.Println("这是一个后台任务")
        time.Sleep(time.Second * 2)
        panic("测试panic处理")
        // 下行不会执行，仅为演示
        // log.Println("后台任务完成")
        // return errors.New("后台任务错误")
    })
    log.Println("等待后台任务完成")
    log.Println("后台任务结果:", a.Get())
}

// 示例3：Async + 主动取消（任务内遵循 ctx）
// 说明：Async 本身不会自动感知 ctx，需要在任务函数中自己 select ctx.Done。
func demo3() {
    log.Println("示例3：Async + 主动取消（任务内遵循 ctx）")
    ctx3, cancel3 := context.WithCancel(context.Background())
    a3 := rr.Async(ctx3, func(ctx context.Context) error {
        ticker := time.NewTicker(300 * time.Millisecond)
        defer ticker.Stop()
        for i := 0; i < 10; i++ {
            select {
            case <-ctx.Done():
                return ctx.Err()
            case <-ticker.C:
                log.Println("工作中...", i)
            }
        }
        return nil
    })
    // 模拟工作一段时间后取消
    time.Sleep(900 * time.Millisecond)
    cancel3()
    log.Println("取消结果:", a3.Get())
}

// 示例4：AsyncResult 基本用法（获取返回值）
func demo4() {
    log.Println("示例4：AsyncResult 基本用法（获取返回值）")
    r1 := rr.AsyncResult(context.Background(), func(ctx context.Context) (int, error) {
        // 计算/IO ...
        time.Sleep(1 * time.Second)
        return random.IntRange(1, 100), nil
    })
    v1, err1 := r1.Get()
    log.Println("AsyncResult 基本用法: 值=", v1, " 错误=", err1)
}

// 示例5：AsyncResult + Done + 超时/取消（任务内支持 ctx）
func demo5() {
    log.Println("示例5：AsyncResult + Done + 超时/取消（任务内支持 ctx）")
    ctx5, cancel5 := context.WithTimeout(context.Background(), 800*time.Millisecond)
    defer cancel5()
    r2 := rr.AsyncResult(ctx5, func(ctx context.Context) (string, error) {
        // 用 select 支持超时/取消
        select {
        case <-time.After(2 * time.Second):
            return "OK", nil
        case <-ctx.Done():
            return "", ctx.Err()
        }
    })
    select {
    case <-r2.Done():
        v2, err2 := r2.Get()
        log.Println("AsyncResult Done 完成: 值=", v2, " 错误=", err2)
    case <-ctx5.Done():
        log.Println("AsyncResult 超时/取消:", ctx5.Err())
        // 任务函数里会很快返回，此处也可选择随后再 Get() 一次拿到最终 error
        v2, err2 := r2.Get()
        log.Println("AsyncResult 超时后最终: 值=", v2, " 错误=", err2)
    }
}

// 示例6：AsyncResult 扇出/聚合（并发多个任务再汇总）
func demo6() {
    log.Println("示例6：AsyncResult 扇出/聚合（并发多个任务再汇总）")
    tasks := []rr.AsyncResultTask[int]{
        rr.AsyncResult(context.Background(), func(ctx context.Context) (int, error) {
            time.Sleep(300 * time.Millisecond)
            return 1, nil
        }),
        rr.AsyncResult(context.Background(), func(ctx context.Context) (int, error) {
            time.Sleep(500 * time.Millisecond)
            // 模拟偶发错误
            if random.IntRange(1, 10) > 5 {
                return 0, errors.New("计算失败")
            }
            return 2, nil
        }),
        rr.AsyncResult(context.Background(), func(ctx context.Context) (int, error) {
            time.Sleep(200 * time.Millisecond)
            return 3, nil
        }),
    }
    sum := 0
    var firstErr error
    for idx, t := range tasks {
        v, err := t.Get()
        if err != nil && firstErr == nil {
            firstErr = err
        }
        log.Println("任务", idx, "完成: 值=", v, " 错误=", err)
        sum += v
    }
    log.Println("聚合结果: sum=", sum, " 首个错误=", firstErr)
}

// 示例7：HeartbeatWait 心跳等待（等待任务完成时定期执行心跳回调）
func demo7() {
    log.Println("示例7：HeartbeatWait 心跳等待（等待任务完成时定期执行心跳回调）")

    // 场景1：正常完成 - 任务在心跳期间完成
    log.Println("场景1：任务正常完成，观察心跳")
    ctx1 := context.Background()
    async1 := rr.Async(ctx1, func(ctx context.Context) error {
        log.Println("  任务1开始执行...")
        time.Sleep(2 * time.Second)
        log.Println("  任务1执行完成")
        return nil
    })

    heartbeatCount := 0
    err := async1.HeartbeatWait(ctx1, 500*time.Millisecond, func() {
        heartbeatCount++
        log.Printf("  ❤️ 心跳 #%d - 任务仍在运行中...", heartbeatCount)
    })
    log.Println("  任务1结果: err=", err, " 心跳次数=", heartbeatCount)

    // 场景2：超时取消 - 任务执行时间过长，通过 context 超时
    log.Println("\n场景2：任务超时，心跳停止")
    ctx2, cancel2 := context.WithTimeout(context.Background(), 1500*time.Millisecond)
    defer cancel2()
    async2 := rr.Async(ctx2, func(ctx context.Context) error {
        log.Println("  任务2开始执行（会运行5秒）...")
        select {
        case <-time.After(5 * time.Second):
            log.Println("  任务2正常完成")
            return nil
        case <-ctx.Done():
            log.Println("  任务2被取消")
            return ctx.Err()
        }
    })

    heartbeatCount2 := 0
    err2 := async2.HeartbeatWait(ctx2, 400*time.Millisecond, func() {
        heartbeatCount2++
        log.Printf("  ❤️ 心跳 #%d - 任务仍在运行中...", heartbeatCount2)
    })
    log.Println("  任务2结果: err=", err2, " 心跳次数=", heartbeatCount2)

    // 场景3：主动取消 - 模拟外部取消
    log.Println("\n场景3：主动取消任务")
    ctx3, cancel3 := context.WithCancel(context.Background())
    async3 := rr.Async(ctx3, func(ctx context.Context) error {
        log.Println("  任务3开始执行...")
        select {
        case <-time.After(10 * time.Second):
            return nil
        case <-ctx.Done():
            log.Println("  任务3收到取消信号")
            return ctx.Err()
        }
    })

    // 在后台goroutine中，模拟1秒后主动取消
    go func() {
        time.Sleep(1 * time.Second)
        log.Println("  主动取消任务3")
        cancel3()
    }()

    heartbeatCount3 := 0
    err3 := async3.HeartbeatWait(ctx3, 300*time.Millisecond, func() {
        heartbeatCount3++
        log.Printf("  ❤️ 心跳 #%d", heartbeatCount3)
    })
    log.Println("  任务3结果: err=", err3, " 心跳次数=", heartbeatCount3)

    // 场景4：任务失败返回错误
    log.Println("\n场景4：任务执行失败")
    ctx4 := context.Background()
    async4 := rr.Async(ctx4, func(ctx context.Context) error {
        log.Println("  任务4开始执行...")
        time.Sleep(1 * time.Second)
        log.Println("  任务4遇到错误")
        return errors.New("任务执行失败")
    })

    heartbeatCount4 := 0
    err4 := async4.HeartbeatWait(ctx4, 300*time.Millisecond, func() {
        heartbeatCount4++
        log.Printf("  ❤️ 心跳 #%d", heartbeatCount4)
    })
    log.Println("  任务4结果: err=", err4, " 心跳次数=", heartbeatCount4)
}

