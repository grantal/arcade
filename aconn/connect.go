// aconn is short for "arcade connect"
package aconn
// connect.go
// contains functions managing how servers connect and register
// on the arcade and how clients find servers on the arcade 
// *** INSTALLATION ***
// make sure this file is at $GOPATH/src/arcade/aconn/connect.go
// it needs that in order to be run properly

import (
    "cs221"
    "fmt"
    "os"
    "strconv"
)


// hardcoded arcade address
var ArcadeHostname string = "ravioli"
var ArcadePort int = 5888

// Connects the client to the arcade and has them choose a server 
// address
// game name should have no spaces in it
func ClientConnect(game string) (string,int){
    out, in, e := cs221.MakeConnection(ArcadeHostname,ArcadePort,game)
    if e != nil {
            fmt.Println(e.Error())
            os.Exit(1)
    }

    out <- game + " client\n\n"
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
    return hostname,port
}

// registers the server on the arcade and then checks in with
// the arcade every time it asks
// WARNING: This function loops forever
func ServerConnect(game string, hostname string, port int) {
    
    out, in, e := cs221.MakeConnection(ArcadeHostname,ArcadePort,game)
    if e != nil {
            fmt.Println(e.Error())
            os.Exit(1)
    }
    
    fmt.Printf("Connected to arcade at %s:%d\n",ArcadeHostname,ArcadePort)
    // wait a sec to make sure arcade is ready
    out <- game + " server\n\n"
    <- in
    out <- fmt.Sprintf("%s %d\n\n", hostname,port)

    // check in with arcade
    for {
        <- in
        out <- "Still Here\n\n"
    }

}
