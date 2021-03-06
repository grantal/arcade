package main
// Handles the server for the nim game
// usage: go run nim_s.go <hostname> <port>

import (
    "cs221"
    "fmt"
    "strconv"
    "os"
    "arcade/aconn"
)

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
    
    // break off new routine to communicate with arcade
    go aconn.ServerConnect("nim", hostname, port)

    fmt.Printf("Listening on %s:%d\n",hostname,port)
    e := cs221.HandleConnections(hostname, port, handleNim, "nim", nil)
    if e != nil {
        fmt.Println(e.Error())
        os.Exit(0)
    }


}

