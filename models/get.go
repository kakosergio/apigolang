package models

import "apigolang/db"

// Função que retorna as informações de um determinado ID
func Get(id int64) (todo Todo, err error) {
	// declarada a variável conn e assigned a ela e à variável err a abertura de conexão
	conn, err := db.OpenConnection()
	// testa se a conexão retornou erro e sai da chamada em caso positivo. O nil significa sem erros para o struct error
	if err != nil {
		return
	}
	// difere para o final do código o fechamento da conexão com o database
	defer conn.Close()

	// salva em uma variável o resultado da query realizada no banco de dados
	row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)

	// ainda não entendi o que isso faz, mas provavelmente tenta pegar as informações recuperadas do database.
	// se tiver erro, retorna ele, senão retorna as informações requeridas e aponta um endereço de memória para elas
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}