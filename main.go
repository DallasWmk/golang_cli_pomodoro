package main

import (
    "flag"
    "fmt"
    "strings"
    "time"
)

func main () {

    process_args()

    for i := 0; i <= 10; i++ {
        fmt.Printf("%s\r", strings.Repeat("#", i))
        time.Sleep(1 * time.Second)
    }
}

func process_args() {
    num_pomodoros := flag.Int(
        "pomos",
        5,
        "-pomos=<num pomodoros>; default: 5",
    )

    pomodoro_length := flag.Int(
        "length",
        25,
        "-length=<# minutes>; default: 25; must be divisible by 5",
    )

    flag.Parse()
    fmt.Println("you want:", *num_pomodoros, "pomodoros of length", *pomodoro_length)
}
