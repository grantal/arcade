package aconn
// connect.go
// contains functions managing how servers connect and register
// on the arcade and how clients find servers on the arcade 
// *** INSTALLATION ***
// in order for this file to be found, the project will need to be
// in your $GOPATH/src
// So, nim_c.go will be at $GOPATH/src/arcade/nim_c.go and this file
// will be at $GOPATH/src/arcade/aconn/connect.go


// hardcoded arcade address
var ArcadeHostname string = "ravioli"
var ArcadePort int = 5888
