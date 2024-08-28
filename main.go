package main

import "tcp-simple-server/server"

var myServer server.Server = server.Server{
	Address: ":8080",
	Network: "tcp",
}

func main() {
	func() {
		server.CreateNewServer(&myServer)
	}()
}
