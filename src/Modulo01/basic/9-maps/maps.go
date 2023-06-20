package main

import "fmt"

func main() {
	fmt.Println("--- Maps ---")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"]) // para acessar algo especifico

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "joao",
			"ultimo":   "pedro",
		},

		"curso": {
			"nome":   "engenharia",
			"campus": "campus 1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)
}
