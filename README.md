# go-programming-language
## Introdução <br>
> A Go é uma linguagem de programação de código aberto, compilada e estaticamente tipada, o que significa que o código-fonte é executado diretamente pelo sistema operacional ou pelo processador após ser traduzido por um compilador, e os tipos das variáveis são definidos no código e conhecidos durante a compilação. Essa linguagem permite a programação simultânea, gerenciamento de memória e execução adiada de algumas funções, e é semelhante à linguagem C, porém mais robusta e bem aceita entre os developers. É considerada de fácil entendimento, alto desempenho e excelente nível prático, e é uma opção para programadores(as) que já estão acostumados com o estilo C ou Java. Embora se comente que ela veio para substituir o Java, ainda há espaço e preferências para ambas as linguagens.<br>

 [Documentação do Golang](https://pkg.go.dev/fmt)

* Os verbos:
    * Em geral:
    ```
    %v o valor em um formato padrão
	ao imprimir structs, o sinalizador de mais (%+v) adiciona nomes de campo
    %#v uma representação da sintaxe Go do valor
    %T uma representação da sintaxe Go do tipo do valor
    %% um sinal de porcentagem literal; não consome nenhum valor
    ```
    * Boleano :
    ```
    %t a palavras true ou false
    ```
    * Inteiro
    ```
    %b base 2 
    %c o caractere representado pelo ponto de código Unicode correspondente 
    %d base 10 
    %o base 8 
    %O base 8 com 0o prefixo 
    %q um literal de caractere entre aspas simples escapado com segurança com a sintaxe Go. 
    %x base 16, com letras minúsculas para af 
    %X base 16, com letras maiúsculas para AF 
    %U Formato Unicode: U+1234; o mesmo que "U+%04X"
    ```
    [Visão Geral - Documentação](https://pkg.go.dev/fmt)