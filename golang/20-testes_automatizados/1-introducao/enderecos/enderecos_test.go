package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//TESTE DE UNIDADE
//Exemplo com STRUCTS
func TestTipoDeEndereco2(t *testing.T) {
	//Executa testes em paralelo
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia Dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA SEM NÚMERO", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de endereço é diferente do esperado! Esperado: %s,  Recebido: %s",
				cenario.retornoEsperado,
				tipoDeEnderecoRecebido)
		}
	}
}

//Exemplo simples
func TestTipoDeEndereco1(t *testing.T) {
	//Executa testes em paralelo
	t.Parallel()

	enderecoParaTeste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo de endereço é diferente do esperado! Esperado: %s,  Recebido: %s",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido)
	}
}
