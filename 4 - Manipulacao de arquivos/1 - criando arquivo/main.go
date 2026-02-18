package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// criação de arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("Ola mundo")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	defer f.Close()

	// leitura
	// arquivo, err := os.ReadFile("arquivo.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(arquivo))

	// leitura linha a linha

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

}
