package main

import (
    "cs221"
    "fmt"
    "strconv"
    "os"
)


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
