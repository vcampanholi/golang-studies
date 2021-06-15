package main

import "fmt"

func main() {
	//ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String 2"
	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 //numero = numero - 20
	fmt.Println(numero)

	numero *= 3 //numero = numero * 3
	fmt.Println(numero)

	numero /= 10
	numero %= 3
	fmt.Println(numero)

	//OPERADORES TERNÁRIOS
	//GO não possui sintaxe para operador ternário
	// texto := numero > 5 ? "Maior que 5" : "Menor que 5" (Isso não funciona em GO)

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)

}
