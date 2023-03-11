package main

import "fmt"

func main() {
    exibeIntruducao()
    exibeMenu()
    comando := lerComando()
// local para colocar o "if" que está comentado em baixo
    switch comando {
        case 1:
            fmt.Println("Monitorando...")
        case 2 : 
            fmt.Println("Exibindo Logs...")
        case 0:
            fmt.Println("Saindo do programa...")
        default:
            fmt.Println("Não conheço este comando")

    }
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