package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	for {
		exibeIntroducao()
		exibeMenu()
		comando := leComando()
	
		switch comando {
			case 1:
				iniciarMonitoramento()
			case 2:
				fmt.Println("Exibindo logs")
			case 0:
				fmt.Printf("Saindo do programa")
				os.Exit(0)
			default:
				fmt.Println("Não conheço esse comando")
				os.Exit(-1) 
		}
	}
}

func exibeIntroducao() {
	versao := 1.1
	fmt.Println("Olá")
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

func iniciarMonitoramento() {
	fmt.Println("Monitoriando ...")
	site := "http://www.alura.com.br"
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site: ", site, "está fora do ar. Status Code: ", resp.StatusCode)
	}
}
