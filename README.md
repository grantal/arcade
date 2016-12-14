## Installation

Theres a few ways to do this:

1. Just put this directory in your `$GOPATH/src/`
2. Make a folder `$GOPATH/src/arcade/` and put `aconn/` in there

## Running the arcade

Right now, I have it set up so that the arcade runs on ravioli at port 5888. So if you log on to ravioli and do `go run arcade.go ravioli 5888` it should work fine. If you want to run it somewhere else, change the `ArcadeHostName` and `ArcadePort` variables in `aconn/connect.go`

## Description of Game

The game I recreated in go is called nim. The game starts with 8 sticks on a table and the players alternate picking up 1-3 sticks. Whoever picks up the last stick wins. The way it is right now, 1 client has to play for both players. 

In the `nim/` directory `nim.go` is the not-networked version, `nim_c.go` is the client that connects to the arcade to find a gameserver, and `nim_s.go` is the server that registers itself with the arcade. You need to run the server like this: `go run nim_s.go <hostname> <port>`. The client does not require any arguments.
