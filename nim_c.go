package main

import (
    "cs221"
    "fmt"
    "strconv"
    "os"
)

import "arcade/aconn"

// hardcoded arcade address
var arcadehostname string = aconn.ArcadeHostname 
var arcadeport int = aconn.ArcadePort

/*
game: nim
players alternate picking up 1-3 sticks and whoever picks up the 
last stick wins
*/
func main() {

    out, in, e := cs221.MakeConnection(arcadehostname,arcadeport,"nimc")
    if e != nil {
            fmt.Println(e.Error())
            os.Exit(1)
    }

    out <- "nim client\n\n"
    gameservers := <- in

    fmt.Print(gameservers)

    // ask user to choose a server
    fmt.Print("Enter number of server (ie 1,2,3): ")
    var choice int
    fmt.Scanf("%d", &choice)

    out <- strconv.Itoa(choice) + "\n\n"
    fmt.Print(strconv.Itoa(choice) + "\n\n")

    report := <- in
    // exit if you didn't choose a good server
    if report == "INVALID CHOICE\n\n" {
        fmt.Println("You didn't choose a valid server.")
        fmt.Println("This may be because there are no servers for your game.")
        os.Exit(1)
    }

    hostname := ""
    port := 0

    fmt.Sscanf(report, "%s %d",&hostname, &port)

    out, in, e = cs221.MakeConnection(hostname,port,"nimc")
    if e != nil {
            fmt.Println(e.Error())
            os.Exit(1)
    }

    fmt.Printf("Connected to %s:%d\n",hostname,port)
    
    n := 8 
    turn := 2
    out <- "Start!\n\n" //tell server to start game
    // get initial data
    data := <- in 
    fmt.Sscanf(data, "%d %d\n",&n, &turn)
    for n > 0 {
        
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
        out <- strconv.Itoa(num) + "\n\n"

        // set sticks and turn by server
        data := <- in 
        fmt.Sscanf(data, "%d %d\n",&n, &turn)
    }

    fmt.Printf("Player %d wins!\n", turn)
    


}
