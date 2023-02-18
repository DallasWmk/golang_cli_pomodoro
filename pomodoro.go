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

    pomo_string := build_pomodoro_string(num_pomodoros, pomodoro_length)
    fmt.Printf("%s\r", pomo_string)

    pomo_segments := pomodoro_length / 5
    idx := 1

    for pomo_counter := 0; pomo_counter < num_pomodoros; pomo_counter++ {
        // Long break
        if pomo_counter != 0 && pomo_counter % 3 == 0 { 
            take_break_long(&pomo_string, &idx)
        } else {
            study(&pomo_string, &idx, pomo_segments)
            idx++
            take_break_short(&pomo_string, &idx)
        }
        idx++
    }
}

func study(pomo_string *string, idx *int, pomo_segments int) {

    for segment_counter := 0; segment_counter < pomo_segments; segment_counter++ {
        time.Sleep(5 * time.Minute)
        update_pomo_string(pomo_string, idx)
        *idx++
    }
}

func update_pomo_string(pomo_string *string, idx *int) {
    pomo_temp := *pomo_string
    *pomo_string = pomo_temp[:*idx] + "#" + pomo_temp[*idx+1:]
    fmt.Printf("%s\r", *pomo_string)
}

func take_break_long(pomo_string *string, idx *int) {
    /*
        Take long break and update pomo_string
    */
    for break_counter := 0; break_counter < 3; break_counter++ {
        update_pomo_string(pomo_string, idx)
        *idx++
    }
    *idx++
}

func take_break_short(pomo_string *string, idx *int) {
    /*
        Take short break and update pomo_string
    */
    time.Sleep(5 * time.Minute)
    update_pomo_string(pomo_string, idx)
    *idx++
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
    fmt.Println("num pomodoros:", num_pomos)    
    breaks_reqd := num_pomos - 1
    pomo_segments := pomo_length / 5
    pomodoro_string := ""

    for rests := 0; rests <= breaks_reqd; rests++ {
        pomodoro_string = pomodoro_string + "|" + strings.Repeat("-", pomo_segments)
        if rests != 0 && (rests+1) % 3 == 0 {
            pomodoro_string = pomodoro_string + "|" + strings.Repeat("-", 3)
            continue
        } 

        pomodoro_string = pomodoro_string + "|" + strings.Repeat("-", 1)
        
    }
    pomodoro_string = pomodoro_string + "|"
    return pomodoro_string

}
