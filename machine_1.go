package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Lista de hashes de arquivos que esta maquina possui
var fileHashes = []string{
	"abc123", 
	"def456",
	"ghi789",
	"jkl112",
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Lê o hash enviado pelo cliente
	hash, _ := bufio.NewReader(conn).ReadString('\n')
	hash = strings.TrimSpace(hash)

	// Verifica se esta maquina possui o arquivo com o hash
	found := false
	for _, fileHash := range fileHashes {
		if hash == fileHash {
			found = true
			break
		}
	}

	// Responde ao cliente com FOUND ou NOT_FOUND
	if found {
		conn.Write([]byte("FOUND\n"))
	} else {
		conn.Write([]byte("NOT_FOUND\n"))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Erro ao iniciar:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Ouvindo na porta 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		go handleConnection(conn)
	}
}
