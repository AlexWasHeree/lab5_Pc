package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var fileHashes = []string{
	"ghi789",
	"jkl112",
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	hash, _ := bufio.NewReader(conn).ReadString('\n')
	hash = strings.TrimSpace(hash)

	found := false
	for _, fileHash := range fileHashes {
		if hash == fileHash {
			found = true
			break
		}
	}

	if found {
		conn.Write([]byte("FOUND\n"))
	} else {
		conn.Write([]byte("NOT_FOUND\n"))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8083")
	if err != nil {
		fmt.Println("Erro ao iniciar:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Ouvindo na porta 8083...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		go handleConnection(conn)
	}
}
