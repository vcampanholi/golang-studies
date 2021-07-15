package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

//Em GO não há herança, mas a estrutura pessoa pode ser usada em outras estruturas sem declarar o tipo
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"João", "Pedro", 10, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "PUCRS"}
	fmt.Println(e1)

	//Recuperar o nome do estudande baseado na struct pessoa
	fmt.Println(e1.nome + " " + e1.sobrenome)
}
