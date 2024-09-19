package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Lista de IPs
var machines = []string{
	"127.0.0.1:8080",
	"127.0.0.1:8081",
	"127.0.0.1:8082",
	"127.0.0.1:8083",
}

// Função que faz a busca
func searchFile(hash string) []string {
	var foundServers []string

	// Itera sobre a lista de IPs das maquinas
	for _, server := range machines {
		fmt.Printf("Conectando ao servidor: %s\n", server)

		// Conecta ao servidor de arquivo
		conn, err := net.Dial("tcp", server)
		if err != nil {
			fmt.Printf("Erro ao conectar ao servidor %s: %v\n", server, err)
			continue
		}
		defer conn.Close()

		// Envia o hash ao servidor de arquivo
		fmt.Fprintf(conn, hash+"\n")

		// Recebe a resposta do servidor
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Erro ao ler a resposta do servidor %s: %v\n", server, err)
			continue
		}
		response = strings.TrimSpace(response)

		// Se a maquina tiver o arquivo, adiciona o IP na lista
		if response == "FOUND" {
			foundServers = append(foundServers, server)
		}
	}

	return foundServers
}

// Função que lida com as conexões dos clientes de busca
func handleClientConnection(conn net.Conn) {
	defer conn.Close()

	// Lê o hash enviado pelo cliente
	hash, _ := bufio.NewReader(conn).ReadString('\n')
	hash = strings.TrimSpace(hash)

	// Faz a busca nos servidores de arquivos
	foundServers := searchFile(hash)

	// Envia os resultados de volta ao cliente
	if len(foundServers) > 0 {
		for _, server := range foundServers {
			conn.Write([]byte(server + "\n"))
		}
	} else {
		conn.Write([]byte("Nenhum servidor possui o arquivo com esse hash.\n"))
	}
}

func main() {
	// Inicia o servidor 
	ln, err := net.Listen("tcp", ":9000") 
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor de busca:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Servidor de busca ouvindo na porta 9000...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		go handleClientConnection(conn)
	}
}
