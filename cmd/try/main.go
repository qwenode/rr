package main

import (
    "log"

    "github.com/qwenode/rr/random"
)

func main() {
	log.Println(random.String(10))
	log.Println(random.StringRange(1, 10))
	log.Println(random.Random(10, random.ASCIILettersLowercase))
	log.Println(random.Random(10, random.ASCIILettersLowercase))
	for i := 0; i < 100; i++ {
		log.Println(random.IntRange(0, 10))
	}
    
    for i := 0; i < 10; i++ {
        log.Println(random.IntRange(0,1),"AA")
    }
     for i := 0; i < 10; i++ {
        log.Println(random.IntRange(0,2),"BB")
    }
     for i := 0; i < 10; i++ {
       log.Println( random.IntRange(1,2),"CC")
    }
    
}
