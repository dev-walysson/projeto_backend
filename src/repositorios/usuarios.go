package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Essa struct vai guardar a conexão com o banco de dados
type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuario(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// metodo inserir do struct Usuarios
func (repositorio Usuarios) Inserir(usuario modelos.Usuario) (uint64, error) {
	//Prepare, prepara o banco para query, é importante que seja fechada no final da execução da função
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha)values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	//Quando a função terminar vai fechar
	defer statement.Close()

	//vai substituir as ?,?,?,? do insert e executar
	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return 0, erro
	}

	//vai retornar o ID do usuário inserido
	ultimoIDInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	//retornando o ID inserido e nulo para o erro
	return uint64(ultimoIDInserido), nil
}
