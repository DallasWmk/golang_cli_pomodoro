package main

import (
    "fmt"
    "strings"
    "time"
)

func main () {
    for i := 0; i <= 10; i++ {
        fmt.Printf("%s\r", strings.Repeat("#", i))
        time.Sleep(1 * time.Second)
    }
}
