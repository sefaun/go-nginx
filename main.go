package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"spider-server/configuration"
)

var (
	counter = 0
	listen_address = []string{}
	server = []string{}
)

func main() {
	config := &configuration.Configuration{}
	listen_address, server = config.ReadConfiguration()

	listener, err := net.Listen("tcp", listen_address[0])

	if err != nil {
		log.Fatal("Failed to Listen: %s", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("Failed to Accept Connection: %s", err)
		}

		backend := chooseBackend()
		fmt.Printf("Backend= %s\n", backend)

		go func() {
			err := proxy(backend, conn)
			if err != nil {
				log.Printf("WARNING: Proxying Failed: %v", err)
			}
		}()
	}

}

func chooseBackend() string {
	s := server[counter%len(server)]
	counter++
	return s
}

func proxy(backend string, conn net.Conn) error {
	bc, err := net.Dial("tcp", backend)

	if err != nil {
		return fmt.Errorf("Failed to Connect to Backend %s:%v", backend, err)
	}

	//c -> bc
	go io.Copy(bc, conn)
	//bc -> c
	go io.Copy(conn, bc)

	return nil
}
