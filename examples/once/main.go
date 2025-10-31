package main

import (
    "errors"
    "log"

    "github.com/qwenode/rr"
    "github.com/qwenode/rr/random"
)

// 本示例展示了如何使用 rr.Once 来确保某个函数只被执行一次。
// rr.Once 是一个类似于 sync.Once 的工具，但它允许在函数执行失败时进行重试。
//
// 示例中，我们尝试执行一个函数，该函数会随机返回错误。
// 如果函数返回错误，rr.Once 会允许我们再次尝试执行该函数。
// 这种机制非常适合需要确保某些操作只成功执行一次的场景，例如初始化操作。
//
// 代码逻辑：
// 1. 定义了一个 rr.Once 实例。
// 2. 使用一个循环模拟多次尝试调用 once.Do。
// 3. 在 Do 方法中，定义了一个匿名函数，该函数会随机返回错误。
// 4. 如果函数执行成功，退出循环；否则继续尝试。
// 5. 最终确保函数只成功执行一次。
//
// 注意：random.IntRange 用于生成随机数，模拟函数可能失败的场景。

func main() {
    var once rr.Once

    for i := 0; i < 10; i++ {
        err := once.Do(func() error {
            log.Println("Executing function...", i)
            if random.IntRange(0, 100) > 50 {
                return errors.New("Simulated error, will retry")
            }
            return nil
        })
        if err != nil {
            log.Println("Error:", err)
            continue
        }
        log.Println("Function executed successfully on attempt", i)
    }
    log.Println("Function executed successfully")
}
