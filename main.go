package main

import "fmt"

func main() {
	nome := "Lucas"
	const idade = 20
	const versao float32 = 1.1

	fmt.Println("Hello,", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair")

	var comando int

	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", comando)
}
