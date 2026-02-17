package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
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
