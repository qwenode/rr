package main

import (
    "log"
    "rr/random"
)

func main() {
    log.Println(random.String(10))
    log.Println(random.StringRange(1, 10))
    log.Println(random.Random(10, random.ASCIILettersLowercase))
    log.Println(random.Random(10, random.ASCIILettersLowercase))
    for i := 0; i < 100; i++ {
        log.Println(random.IntRange(0, 10))
    }
}
