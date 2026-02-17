package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// vai retornar um router com todas as rotas configuradas
func Gerar() *mux.Router {

	//r é um ponteiro mux.Router e dentro desse Router tem um slice de rotas vazia dentro do Router
	r := mux.NewRouter()
	//Quando configurar rodar, todas as rotas lá do handle é adicionada nesse slice de router vazio
	return rotas.Configurar(r)
}
