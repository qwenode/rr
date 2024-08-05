package main

import (
    "log"
    "rr"
)

func main() {
    t := rr.NewExceptionT("abc").WithT("abbb")
    log.Println(t)
    log.Println(t.StackMessages())
    log.Println(t.Error())
    log.Println(t.IsT("abbb"))
}
