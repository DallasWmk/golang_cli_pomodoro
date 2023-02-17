package main

import (
    "flag"
    "fmt"
    "strings"
    "time"
)

func main () {

    num_pomodoros, pomodoro_length := process_args()
    //print_banner()

    fmt.Println(pomodoro_length)

    pomo_string := build_pomodoro_string(num_pomodoros, pomodoro_length)
    fmt.Println(pomo_string)

    for pomo_counter := 0; pomo_counter <= num_pomodoros; pomo_counter++ {
        fmt.Printf("%s\r", strings.Repeat("#", pomo_counter))
        time.Sleep(5 * time.Minute)
    }
}

func process_args() (int, int) {
    /*
        Process arguments passed as flags in command like evokation
    */
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

    return *num_pomodoros, *pomodoro_length
}

func build_pomodoro_string(num_pomos int, pomo_length int) string {
    /*
        based on the desired # of pomodoros and the resulting breaks req'd 
        generate the initial string to display to the console
    */

    if !(num_pomos > 0 && pomo_length % 5 == 0) {
        // cant build string for no pomos
        // idk maybe make a funny joke
    }
    
    breaks_reqd := num_pomos - 1
    pomo_segments := pomo_length / 5
    pomodoro_string := ""

    for rests := 0; rests <= breaks_reqd; rests++ {
        pomodoro_string = pomodoro_string + "|" + strings.Repeat("-", pomo_segments)
        if rests != 0 && rests % 4 == 0 {
            pomodoro_string = pomodoro_string + "|" + strings.Repeat("-", 3)
            break
        } 

        pomodoro_string = pomodoro_string + "|" + strings.Repeat("-", 1)
        
    }
    pomodoro_string = pomodoro_string + "|"
    return pomodoro_string

}

func print_pomodoro_tracker() {
    /*
        print the updated pomodoro tracker to the console.
    */
}
