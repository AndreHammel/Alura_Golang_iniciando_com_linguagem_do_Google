package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const MONITORAMENTO = 5
const DELAY = 3

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


func iniciarMonitoramento() {
	fmt.Println("Monitoriando ...")
	sites := leSitesDoArquivo()

	for i:=0; i < MONITORAMENTO; i++ {
		time.Sleep(DELAY * time.Second)
		for i, site := range sites{
			fmt.Println("Testando site", i, ":", site )
			testaSite(site)
		}
		fmt.Println()
	}
	fmt.Println()
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro na requisição:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site: ", site, "está fora do ar. Status Code: ", resp.StatusCode)
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

func leSitesDoArquivo() ([]string) {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}
