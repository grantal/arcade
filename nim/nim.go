package main

import (
    "fmt"
)

/*
game: nim
This is the single computer version of nim
players alternate picking up 1-3 sticks and whoever picks up the 
last stick wins
*/
func main() {

    n := 8
    fmt.Print("Enter number of sticks: ")
    fmt.Scanf("%d", &n)
    turn := 2 // which player's turn it is

    for n > 0 {
        turn = ((turn - 1)^1) + 1 
        i := 0
        for i < n {
            fmt.Print("| ")
            i++
        }
        fmt.Print("\n")
        fmt.Printf("Player %d's turn. How many sticks to pick up?\n", turn)
        
        num := 0
        for num < 1 || num > 3{
            fmt.Print("Enter a number from 1-3: ")
            fmt.Scanf("%d", &num)
        }
        n -= num
    }

    fmt.Printf("Player %d wins!\n", turn)
    


}
