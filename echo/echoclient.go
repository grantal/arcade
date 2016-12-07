package main

import (
	"cs221"
	"fmt"
	"strconv"
	"os"
)

func main() {
	hostname := os.Args[1]
	port,_ := strconv.Atoi(os.Args[2])

	out, in, e := cs221.MakeConnection(hostname,port,"Shout")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}

	fmt.Println("Enter lines of text below:")
	for {
		var message string	
		fmt.Scanln(&message)	
		out <- message + "\n"
		out <- "\n"
		reply := <-in
		fmt.Println(cs221.HeadLine(reply))
	}
}
