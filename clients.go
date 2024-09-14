package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var serverIPs = []string{
	"127.0.0.1:8080", 
	"127.0.0.1:8081", 
	"127.0.0.1:8082", 
	"127.0.0.1:8083",
}

func searchFile(hash string) []string {
	var foundServers []string

	for _, server := range serverIPs {
		fmt.Printf("Conectando ao servidor: %s\n", server)

		conn, err := net.Dial("tcp", server)
		if err != nil {
			fmt.Printf("Erro ao conectar ao servidor %s: %v\n", server, err)
			continue
		}
		defer conn.Close()

		fmt.Fprintf(conn, hash+"\n")

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Erro ao ler a resposta do servidor %s: %v\n", server, err)
			continue
		}
		response = strings.TrimSpace(response)

		if response == "FOUND" {
			foundServers = append(foundServers, server)
		}
	}

	return foundServers
}

func main() {
	if len(os.Args) < 3 || os.Args[1] != "search" {
		fmt.Println("Uso: ./program_bay search <file_hash>")
		return
	}

	hash := os.Args[2]

	foundServers := searchFile(hash)

	if len(foundServers) > 0 {
		for _, server := range foundServers {
			fmt.Println(server)
		}
	} else {
		fmt.Println("Nenhum servidor possui o arquivo com esse hash.")
	}
}
