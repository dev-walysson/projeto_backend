package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	//Vai cadastrar usuários
	{
		URI:                "/usuarios",     //caminho da rota
		Metodo:             http.MethodPost, //tipo de requisição
		funcao:             controllers.InserirUsuario,
		RequerAutenticacao: false,
	},
	//Vai buscar todos os usuários
	{
		URI:                "/usuarios",    //caminho da rota
		Metodo:             http.MethodGet, //tipo de requisição
		funcao:             controllers.ListarUsuarios,
		RequerAutenticacao: false,
	},
	//Vai buscar o usuário pelo ID
	{
		URI:                "/usuarios/{usuarioId}", //caminho da rota
		Metodo:             http.MethodGet,          //tipo de requisição
		funcao:             controllers.ListarUsuarioID,
		RequerAutenticacao: false,
	},
	//Vai atualizar o usuário pelo ID
	{
		URI:                "/usuarios/{usuarioId}", //caminho da rota
		Metodo:             http.MethodPut,          //tipo de requisição
		funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	//vai deletar o usuário pelo ID
	{
		URI:                "/usuarios/{usuarioId}", //caminho da rota
		Metodo:             http.MethodDelete,       //tipo de requisição
		funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
