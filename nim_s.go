package main

import (
    "cs221"
    "fmt"
    "strconv"
    "os"
)

// hardcoded arcade address
var arcadehostname string = "localhost"
var arcadeport int = 5888

func handleNim(out chan<- string, in <-chan string, info interface{}) {
    fmt.Println("Client connected")

    n := 8
    turn := 2 // which player's turn it is
    <- in //wait for player to send something
    for n > 0 {
        turn = ((turn - 1)^1) + 1 
        // send turn and number of sticks to player
        report := fmt.Sprintf("%d %d\n\n", n, turn)
        out <- report
        
        num := 0
        for num < 1 || num > 3{
            input := <- in
            fmt.Sscanf(input, "%d", &num)
        }
        n -= num
    }
    // send one last report
    report := fmt.Sprintf("%d %d\n\n", n, turn)
    out <- report
    fmt.Println("Game Over")
    
}


func main() {

    hostname := os.Args[1]
    port,_ := strconv.Atoi(os.Args[2])
    
    out, in, e := cs221.MakeConnection(arcadehostname,arcadeport,"nimarcade")
    if e != nil {
            fmt.Println(e.Error())
            os.Exit(1)
    }
    

    fmt.Printf("Connected to arcade at %s:%d\n",arcadehostname,arcadeport)
    // wait a sec to make sure arcade is ready
    //time.Sleep(10*time.Second) 
    //fmt.Println("Sleep Over")
    out <- "nim server\n\n"
    <- in
    out <- fmt.Sprintf("%s %d\n\n", hostname,port)


    fmt.Printf("Listening on %s:%d\n",hostname,port)
    e = cs221.HandleConnections(hostname, port, handleNim, "nim", nil)
    if e != nil {
        fmt.Println(e.Error())
        os.Exit(0)
    }
}

