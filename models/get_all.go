package models

import "apigolang/db"

// Função que retorna as informações de tudo que for encontrado na tabela de to dos
func GetAll() (todos []Todo, err error) {
	// declara duas variáveis para receber os dois retornos da função OpenConnection, sendo um retorno a conexão em si e o outro um
	// possível erro de conexão, sendo que se o valor de err for igual a nil, não há erros.
	conn, err := db.OpenConnection()
	// testa se a conexão retornou erro e sai da chamada em caso positivo. O nil significa sem erros para o struct error
	if err != nil {
		return
	}
	// difere para o final do código o fechamento da conexão com o database
	defer conn.Close()

	// salva em uma variável o resultado da query realizada no banco de dados
	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	for rows.Next(){
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	}
	// ainda não entendi o que isso faz, mas provavelmente tenta pegar as informações recuperadas do database.
	// se tiver erro, retorna ele, senão retorna as informações requeridas e aponta um endereço de memória para elas

	return
}