package main

import (
	"fmt"
	"reflect"
)

func main() {
	//ARRAYS
	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 10
	fmt.Println(array1)

	array2 := [5]string{"Pos 01", "Pos 02", "Pos 03", "Pos 04", "Pos 05"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//SLICES
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	//Adiciona um item novo e retorna o slice com o novo item
	slice = append(slice, 18)
	fmt.Println(slice)

	//Cria um novo slice a partir de uma fatia de um array
	slice2 := array2[1:3] //1 -> inclusivo, 3 -> exclusivo
	fmt.Println(slice2)

	array2[1] = "Pos Alterada"
	fmt.Println(slice2)

	//Exibe o tipo de uma variável
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	//ARRAYS INTERNOS

	//Função MAKE aloca um espaço na memória para uma determinada estrutura
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	//Se estourar o tamanho do slice automaticamente ele irá dobrar a capacidade
	fmt.Println(len(slice4)) //length
	fmt.Println(cap(slice4)) //capacidade

}
