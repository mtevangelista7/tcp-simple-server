package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Server struct {
	Address  string
	Network  string
	Listener net.Listener
}

func CreateNewServer(server *Server) {
	var err error
	server.Listener, err = net.Listen(server.Network, server.Address)

	defer server.Listener.Close()

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(">>> running...")

	acceptConnections(&server.Listener)
}

func acceptConnections(listener *net.Listener) {
	for {
		ln := *listener

		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("error:", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		message = strings.TrimSpace(message)

		fmt.Println(">>> ", message)

		if strings.ToLower(message) == "close" {
			fmt.Println(">>> see you...")
			return
		}
	}
}
