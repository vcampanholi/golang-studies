package main

import (
	"errors"
	"fmt"
)

func main() {
	//---------------------- INTEIROS -----------------
	numeroArquitetura := 50
	fmt.Println(numeroArquitetura)

	var numero int16 = 100
	fmt.Println(numero)

	//unsigned int
	var numero2 uint = 1500
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 1234
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//---------------------- NÚMEROS REAIS -----------------

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1230.87
	fmt.Println(numeroReal3)

	//---------------------- STRINGS -----------------

	var srt string = "Texto"
	fmt.Println(srt)

	str2 := "Texto 2"
	fmt.Println(str2)

	//---------------------- BOOLEAN -----------------

	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool = false
	fmt.Println(booleano2)

	//---------------------- ERRO -----------------

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

	//---------------------- VALOR ZERO -----------------
	//Valor inicial da variáveis
	//String = vazio
	//Int = 0
	//Boolean = false

	var texto string
	fmt.Println(texto)

	var texto2 int8
	fmt.Println(texto2)

	var boolean bool
	fmt.Println(boolean)

}
