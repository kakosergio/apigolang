package db

import (
	"apigolang/configs"
	"database/sql"
	"fmt"
	// importa o driver de conexão com o PostgreSQL
	_ "github.com/lib/pq"
)

// Função que abre a conexão com o database e retorna uma conexão e/ou um erro
func OpenConnection() (*sql.DB, error){
	// cria variável para pegar as configurações de acesso ao db
	conf := configs.GetDB()

	// cria a string de acesso ao DB
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
	// abre a conexão com o DB e grava caso haja erro na segunda variável declarada
	// me parece que quando uma função tem dois retornos, você precisa declarar duas variáveis
	// no caso abaixo, a função Open do pacote sql tem dois retornos, sendo um deles a conexão e o outro um erro, que pode ser nulo ou n
	// por esse motivo é necessário declarar duas variáveis, para guardar os dois retornos
	conn, err := sql.Open("postgres", sc)

	// valida se há erro e encerra a aplicação
	if err != nil {
		// esse panic encerra a aplicação e retorna o erro no terminal.
		//! não usar panic em produção!
		panic(err)
	}

	// se passar da validação, ou seja, se err contiver nil, faz um ping no db. Se o retorno for nil, tá tudo certo.
	err = conn.Ping()

	// retorna a conexão e a variável err para quem chamou a função
	return conn, err
}
