package main

import (
    "cs221"
    "fmt"
    "strconv"
    "os"
    "sync"
    "time"
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
    gamemap map[string][]*gameservice  
    lock    *sync.Mutex
}

func makeGameService(h string, p int, g string) *gameservice{
    return &gameservice{host:h,port:p,game:g}
}

func makeRegistry() *registry{
    return &registry{gamemap:make(map[string][]*gameservice),lock:&sync.Mutex{}}
}

// locks the registry mutex and adds a gameservice to the gamemap
// then unlocks
func (r *registry) addGame(g *gameservice){
    r.lock.Lock()
    r.gamemap[g.game] = append(r.gamemap[g.game], g)
    r.lock.Unlock()
}

// locks the registry mutex and removes a gameservice from the 
// gamemap by pointer to the gameservice
func (r *registry) removeGame(g *gameservice){
    r.lock.Lock()
    i := 0
    for i < len(reg.gamemap[g.game]) {
        if r.gamemap[g.game][i] == g {
            r.gamemap[g.game] = append(r.gamemap[g.game][:i], r.gamemap[g.game][i+1:]...)
        }
        i++
    }
    reg.lock.Unlock()
}

// registry to be shared between threads
var reg *registry = makeRegistry()

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
        fmt.Printf("%s server on %s:%d\n",game,host,port)
        // object to be added to registry
        gam := makeGameService(host,port,game)
        // add gam to registry
        reg.addGame(gam)
        // check up on server to make sure its still running
        response := "Still Here\n\n"
        for response == "Still Here\n\n" {
            time.Sleep(10*time.Second)   
            out <- "Are You still there?\n\n"
            response = <- in
        }
        // remove gam from registry
        reg.removeGame(gam)
        fmt.Printf("%s:%d removed\n",host,port)
        
    } else {
        // give client list of services with their game
        report := fmt.Sprintf("%s services connected\n",game)
        i := 0
        reg.lock.Lock()
        for i < len(reg.gamemap[game]) {
            igame := reg.gamemap[game][i].game
            host := reg.gamemap[game][i].host
            port := reg.gamemap[game][i].port
            report += fmt.Sprintf("%d. %s on %s:%d\n",i+1,igame,host,port)
            i++
        }
        reg.lock.Unlock()
        out <- report + "\n"
        // get client choice
        report = <- in
        var choice int
        fmt.Sscanf(report, "%d",&choice)
        // send client serve info they requested
        reg.lock.Lock()
        host := reg.gamemap[game][choice-1].host 
        port := reg.gamemap[game][choice-1].port 
        reg.lock.Unlock()
        out <- fmt.Sprintf("%s %d\n\n",host,port)
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
