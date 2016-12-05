package main

import (
    "cs221"
    "fmt"
    "strconv"
    "os"
    "sync"
)

// holds info related to each game server
type gameservice struct {
    host string
    port int
    game string
}

// will hold the map of game services and 
// the lock to that map
type registry struct {
    // gamemap is a map from the name of the game to all of the 
    // registered game services running that game
    gamemap map[string][]gameservice  
    lock    *sync.Mutex
}

func makeGameService(h string, p int, g string) *gameservice{
    return &gameservice{host:h,port:p,game:g}
}

func makeRegistry() *registry{
    return &registry{gamemap:make(map[string][]*gameservice),lock:&sync.Mutex{}}
}

func handleArcade(out chan<- string, in <-chan string, info interface{}) {
    // the person connecting should say if they're a server or a client and what game they are
    report := <- in
    game := ""
    ctype := ""
    fmt.Sscanf(report, "%s %s\n", &game, &ctype)
    if ctype == "server" {
        out <- "send host and port\n\n"
        report = <- in
        host := ""
        port := 0
        fmt.Sscanf(report, "%s %d", &host, &port)
        fmt.Printf("%s server on %s:%d",game,host,port)
    }
}


func main() {
    hostname := os.Args[1]
    port,_ := strconv.Atoi(os.Args[2])
    fmt.Printf("Arcade listening on %s:%d\n",hostname,port)
    fmt.Println("Will print out connected services as they connect")
    e := cs221.HandleConnections(hostname, port, handleArcade, "arcade", nil)
    if e != nil {
        fmt.Println(e.Error())
        os.Exit(0)
    }
}
