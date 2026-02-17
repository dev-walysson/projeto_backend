package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConexaoBanco, é a string de conexão com o MySQL
	StringConexaoBanco = ""

	//Porta onde a API vai rodar
	Porta = 0
)

// Vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}


	//está pegando a porta lá do .env, se retornar erro vai setar porta 4000
	Porta, erro = strconv.Atoi(os.Getenv("PORTA"))
	if erro != nil {
		Porta = 4000
	}

	//está sendo pego do arquivo .env / variáveis de ambiente
	StringConexaoBanco = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_BANCO"),
	)
}
