package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)                  // devolver em bytes
	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) // aqui ele converte para json que conseguimos ler

	// outra maneira de fazer a mesma coisa

	c2 := map[string]string{
		"nome": "Max",
		"raca": "Viralata",
	}
	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
