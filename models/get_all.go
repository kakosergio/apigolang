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

	// salva em uma variável o resultado da query realizada no banco de dados e guarda eventual erro na query, considerando que a função
	// tem dois retornos
	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	// um loop que vai passar por cada um dos itens recuperados do database e fazer um Scan para gravar as informações na variável todo
	// a função Scan funciona como uma função void que realiza uma operação sem retornar nada. Contudo, ela tem um retorno, que é do
	// tipo error, por isso a função é rodada dentro da variável err, porque se retornar algo diferente de nil, é considerado erro
	// Scan parece ser uma função do package sql, que lê as informações recuperadas do database e converte em tipos comuns do go, como int64, string
	// entre outros. Além disso, o Next, que lê item por item de uma lista, parece ser uma função do package sql também.
	for rows.Next(){
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	}

	return
}