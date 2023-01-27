package main

import (
	"bufio"
	"log"
	"os"

	tls "github.com/secure-for-ai/goktls"
)

func main() {
	log.SetFlags(log.Lshortfile)
	kw, err := os.OpenFile("same-key.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Println(err)
		return
	}
	conf := &tls.Config{
		InsecureSkipVerify: true,
		MaxVersion:         tls.VersionTLS13,
		MinVersion:         tls.VersionTLS12,
		KeyLogWriter:       kw,
		CipherSuites:       nil,
	}

	conn, err := tls.Dial("tcp", "localhost:9003", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	w := bufio.NewWriter(conn)
	r := bufio.NewReader(conn)
	for {
		n, err := w.Write([]byte("Hello World!\n"))
		if err != nil {
			log.Println(n, err)
			return
		}

		err = w.Flush()
		if err != nil {
			log.Println(n, err)
			return
		}

		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("server msg: %s", msg)
		return
	}
}
