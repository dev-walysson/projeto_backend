package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota, é uma estrutura de definição para rotas da API
type Rota struct {
	URI                string
	Metodo             string
	funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

//Configurar está recebendo um ponteiro de Router
func Configurar(r *mux.Router) *mux.Router {
	//rotas está recebendo um slice de struct
	rotas := rotasUsuarios
	//percorre todo o slice de rotas retornando os valores 
	for _, rota := range rotas {
		//quando acessar essa URI, executa essa função .Methods(rota.Metodo) só responderá ao metodo declarado
		r.HandleFunc(rota.URI, rota.funcao).Methods(rota.Metodo)
		
	}
	return r
}
