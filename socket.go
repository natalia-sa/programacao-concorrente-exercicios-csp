package main

import {
	"io"
	"log"
	"net"
	"time"
}

func main() {
	listener, err := net.listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// cria uma rontina para cada conexao
		go handleConn(conn)
	}

}

fun handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("02:05:00\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}