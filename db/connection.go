package db

import (
	"apigolang/configs"
	"database/sql"
	"fmt"
)

// cria função para abrir a conexão com o db
func OpenConnection() (*sql.DB, error){
	// cria variável para pegar as configurações de acesso ao db
	conf := configs.GetDB()

	// cria a string de acesso ao DB
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
	// abre a conexão com o DB e grava caso haja erro na segunda variável declarada
	conn, err := sql.Open("postgres", sc)

	// valida se há erro e encerra a aplicação
	if err != nil {
		// esse panic encerra a aplicação e retorna o erro.
		//! não usar panic em produção!
		panic(err)
	}

	// se passar da validação, faz um ping no db
	err = conn.Ping()

	// retorna as duas variáveis
	return conn, err
}
