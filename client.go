package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func searchFileOnSearchServer(hash string) {
	// Conecta ao servidor 
	conn, err := net.Dial("tcp", "127.0.0.1:9000") 
	if err != nil {
		fmt.Printf("Erro ao conectar ao servidor de busca: %v\n", err)
		return
	}
	defer conn.Close()

	// Envia o hash ao servidor
	fmt.Fprintf(conn, hash+"\n")

	// Recebe os resultados do servidor\
	fmt.Println("Servidores com o arquivo:")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler a resposta do servidor de busca: %v\n", err)
	}
}

func main() {
	if len(os.Args) < 3 || os.Args[1] != "search" {
		fmt.Println("Uso: ./client search <file_hash>")
		return
	}

	// Extrai o hash do arquivo dos argumentos
	hash := os.Args[2]

	// Faz a busca no servidor de busca
	searchFileOnSearchServer(hash)
}
