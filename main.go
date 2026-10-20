package main

import (
	"http-server/server"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("error creating listener:", err)
	}

	log.Print("server running on port: :8080")

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal("error accepting connection:", err)
		}

		go server.HandleConnection(conn)
	}
}
