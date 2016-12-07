package main

import (
	"cs221"
	"fmt"
	"strconv"
	"os"
	"math/rand"
)

func handleEcho(out chan<- string, in <-chan string, info interface{}) {
	for {
		message := <-in
		if rand.Intn(10) == 0 {
			out <- fmt.Sprintf("Come again, eh?!??\n")
			out <- "\n"
		} else {
			msg := cs221.HeadLine(message)
			out <- fmt.Sprintf("%s, eh!\n",msg)
			out <- "\n"
		}
	}
}

func main() {
	hostname := os.Args[1]
	port,_ := strconv.Atoi(os.Args[2])
	e := cs221.HandleConnections(hostname, port, handleEcho, "Echo", nil)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
}




