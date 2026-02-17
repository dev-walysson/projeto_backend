package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func InserirUsuario(w http.ResponseWriter, r *http.Request) {

	//Pegando as informações que estão vindo no corpo da requisição
	corpoRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	//variavel usuario do tipo Usuario, que é uma struct
	var usuario modelos.Usuario

	//O slice de byte é recebido e o JSON interpretado e os campos dos usuario que foi criado vazio é preenchido.
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	//conectando o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	//Objeto responsável por manipular o usuário no banco
	repositorio := repositorios.NovoRepositorioDeUsuario(db)

	//inserindo dados no banco
	usuarioID, erro := repositorio.Inserir(usuario)
	if erro != nil {
		log.Fatal(erro)
	}
	//retornando uma resposta com o ID do usuário inserido
	w.Write([]byte(fmt.Sprintf("Usuário de id: %d inserido", usuarioID)))
}
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando todos os Usuários"))
}
func ListarUsuarioID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando Usuário pelo ID"))
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário pelo ID"))
}
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário pelo ID"))
}
