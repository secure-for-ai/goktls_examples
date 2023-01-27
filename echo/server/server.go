package main

import (
	"bufio"
	"log"
	"math"
	"net"

	tls "github.com/secure-for-ai/goktls"
)

func main() {
	log.SetFlags(log.Lshortfile)
	cerRSA, err := tls.LoadX509KeyPair("../../certs/server.rsa4096.crt", "../../certs/server.rsa4096.key")
	if err != nil {
		log.Println(err)
		return
	}

	cerDSA, err := tls.LoadX509KeyPair("../../certs/server.ecdsa.secp521r1.crt", "../../certs/server.ecdsa.secp521r1.key")
	if err != nil {
		log.Println(err)
		return
	}

	// server only uses TLS 1.2 and TLS 1.3
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		Certificates: []tls.Certificate{cerRSA, cerDSA},
	}
	ln, err := tls.Listen("tcp", "localhost:9003", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		length := int(math.Min(100, float64(len(msg))))
		log.Printf("client msg: %s", msg[:length])

		n, err := conn.Write([]byte(msg))
		if err != nil {
			log.Println(n, err)
		}
		return
	}
}
