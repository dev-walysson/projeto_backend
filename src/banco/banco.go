package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// vai abrir conexão com o banco de dados
func Conectar() (*sql.DB, error) {

	//vai validar as os argumentos
	db, erro := sql.Open("mysql", config.StringConexaoBanco)

	if erro != nil {
		return nil, erro
	}

	//testando conexão
	if erro := db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	//se ocorrer tudo bem, retorna a conexão
	return db, nil
}
