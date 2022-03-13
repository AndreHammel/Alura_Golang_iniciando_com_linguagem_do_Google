package main

import (
	"fmt"
	"os"
)

func main() {
	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	// if comando == 1 {
	// 	fmt.Println("É o 1")
	// } else if comando ==2 {
	// 	fmt.Println("É o 2")
	// }

	switch comando {
		case 1:
			fmt.Println("Monitoriando ...")
		case 2:
			fmt.Println("Exibindo logs")
		case 0:
			fmt.Printf("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1) // status 255
	}

}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá Mundo", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar o Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}
