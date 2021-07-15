package main

import "fmt"

func main() {

	//Chave e valores devem ser do mesmo tipo
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"segundo":  "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "PUCRS",
		},
	}
	fmt.Println(usuario2)

	//Remove uma chave/valor do MAP
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	//Adiciona uma chave/valor ao MAP
	usuario2["signo"] = map[string]string{
		"nome": "escorpião",
	}
	fmt.Println(usuario2)

}
