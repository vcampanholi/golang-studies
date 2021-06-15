package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var usuario1 usuario
	//Imprime struct vazio
	fmt.Println(usuario1)

	usuario1.nome = "Davi"
	usuario1.idade = 21
	fmt.Println(usuario1)

	enderecoExemplo := endereco{"Rua sem número", 100}

	usuario2 := usuario{"Pedro", 25, enderecoExemplo}
	fmt.Println(usuario2)

	//Informa somente um valor para a Struct
	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
}
