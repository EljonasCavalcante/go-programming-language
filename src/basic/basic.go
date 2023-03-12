package main

import "fmt"
import "os"
// import "net/http"

func main() {
    // exibeIntruducao()
    // exibeMenu()
    nome, idade := devolveNomeEIdade()
    fmt.Println(nome, "E tenho", idade)

    comando := lerComando()
// local para colocar o "if" que está comentado em baixo
    switch comando {
        case 1:
            iniciarMonitoriamento()            
        case 2 : 
            fmt.Println("Exibindo Logs...")
        case 0:
            fmt.Println("Saindo do programa...")
            os.Exit(0)
        default:
            fmt.Println("Não conheço este comando")
            os.Exit(-1) // mostra que aconteceu um erro
    }
}

func devolveNomeEIdade() (string, int) {
    nome := "Jonas"
    idade := 24
    return nome, idade
}

func exibeIntruducao() {
    nome := "Jonas"
    versao := 1.1
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}

func lerComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O comando escolhido foi", comandoLido)
    return comandoLido
}

func iniciarMonitoriamento() {
    fmt.Println("Iniciar Monitoriamento...")
    // site := "https://www.alura.com.br"
    // resp,err := http.Get(site)
}

// colocar em cima para testar
    // if comando == 1 {
    //     fmt.Println("Monitorando...")
    // } else if comando == 2 {
    //     fmt.Println("Exibindo Logs...")
    // } else if comando == 0 {
    //     fmt.Println("Saindo do programa...")
    // } else {
    //     fmt.Println("Não conheço este comando")
    // }