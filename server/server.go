package server

import (
	"bufio"
	"log"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		req, err := Parse(reader)

		log.Printf("%v", req)

		if err != nil {
			log.Printf("failed to parse data: %v", err)
		}
	}
}
