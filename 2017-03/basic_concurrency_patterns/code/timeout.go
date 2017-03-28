package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(time.Millisecond * 250)
    boom := time.After(time.Second * 1)
    for {
        select { // HL
        case <-ticker.C: // HL
            fmt.Println("tick")
        case <-boom: // HL
            fmt.Println("boom!")
            return
        } // HL
    }
}
