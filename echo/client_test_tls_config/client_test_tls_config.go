package main

import (
	"bufio"
	"log"

	tls "github.com/secure-for-ai/goktls"
)

var tlsVersion = map[uint16]string{
	tls.VersionTLS10: "TLS 1.0",
	tls.VersionTLS11: "TLS 1.1",
	tls.VersionTLS12: "TLS 1.2",
	tls.VersionTLS13: "TLS 1.3",
}

// Kernel TLS 1.2 CipherSuites
var cipherSuitesTLS12 = map[uint16]string{
	// AES_128_GCM
	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:   "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256: "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
	tls.TLS_RSA_WITH_AES_128_GCM_SHA256:         "TLS_RSA_WITH_AES_128_GCM_SHA256",
	// AES_256_GCM
	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:   "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384: "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
	tls.TLS_RSA_WITH_AES_256_GCM_SHA384:         "TLS_RSA_WITH_AES_256_GCM_SHA384",
	//CHACHA20_POLY1305
	tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256:   "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
	tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256: "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
}

// Kernel TLS 1.3 CipherSuites
var cipherSuitesTLS13 = map[uint16]string{
	tls.TLS_AES_128_GCM_SHA256: "TLS_AES_128_GCM_SHA256",
	tls.TLS_AES_256_GCM_SHA384: "TLS_AES_256_GCM_SHA384",
	tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256: "TLS_CHACHA20_POLY1305_SHA256",
}

func test_tls_config(conf *tls.Config) {
	var cipherSuites map[uint16]string
	if conf.MaxVersion == tls.VersionTLS12 {
		cipherSuites = cipherSuitesTLS12
		// log.Printf("tls config: %s CipherSuites: %s", tlsVersion[conf.MaxVersion], cipherSuites[uint16(conf.CipherSuites[0])])
	} else {
		cipherSuites = cipherSuitesTLS13
		// log.Printf("tls config: %s CipherSuites: %s", tlsVersion[conf.MaxVersion], cipherSuites[uint16(conf.CipherSuites[0])])
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:9003", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	w := bufio.NewWriter(conn)
	r := bufio.NewReader(conn)
	for {
		n, err := w.Write([]byte("Hello World! " + tlsVersion[conf.MaxVersion] + " " + cipherSuites[uint16(conf.CipherSuites[0])] + "\n"))
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

func main() {
	log.SetFlags(log.Lshortfile)
	// test TLS 1.2
	log.Println("=======================")
	log.Println("Test TLS 1.2")
	log.Println("=======================")
	for ciphersuite := range cipherSuitesTLS12 {
		conf := tls.Config{
			InsecureSkipVerify: true,
			MaxVersion:         tls.VersionTLS12,
			CipherSuites:       []uint16{ciphersuite},
		}
		test_tls_config(&conf)
	}

	// test TLS 1.3
	log.Println("=======================")
	log.Println("Test TLS 1.3")
	log.Println("=======================")
	for ciphersuite := range cipherSuitesTLS13 {
		conf := tls.Config{
			InsecureSkipVerify: true,
			MaxVersion:         tls.VersionTLS13,
			CipherSuites:       []uint16{ciphersuite},
		}
		test_tls_config(&conf)
	}
}
