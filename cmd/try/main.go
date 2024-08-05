package main

import (
    "log"
    "rr/random"
)

func main() {
    log.Println(random.String(10))
    log.Println(random.StringRange(1, 10))
    log.Println(random.Random(10, random.ASCIILettersLowercase))
    log.Println(random.RandomInSecure(10, random.ASCIILettersLowercase))
    log.Println(random.RandomInSecure(10, random.ASCIILettersLowercase))
}
