package main

import (
	"log"
	"net/http"
)

//HTTP é um protocolo de comunicação - Base da comunicação WEB
// Cliente(Faz a requisição) - Servidor(Processa a requisição e envia resposta)

// Rotas:
//URI -> Identificador do recurso
//Método -> GET, POST, PUT, DELETE

func raiz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página Raiz"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Teste HTTP"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários"))
}

func main() {
	http.HandleFunc("/", raiz)
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
