package models

import "apigolang/db"

// Função que recebe um ID e uma tarefa (to do), e atualiza a tarefa
func Update(id int64, todo Todo) (int64, error) {
	// abre a conexão com o banco de dados. Lembrando que a função OpenConnection retorna dois valores, a própria conexão e um erro
	// por isso que deve-se declarar duas variáveis quando quer guardar a função em variáveis, para que mantenha as duas respostas
	conn, err := db.OpenConnection()

	// trata o erro. Se for diferente de nulo, encerra a função retornando um int e um erro
	if err != nil {
		return 0, err
	}

	// difere o fechamento da conexão para o final da execução do código
	defer conn.Close()

	// cria duas variáveis, uma de resposta e uma de erro para gravar o retorno da função abaixo, que executa uma query no banco de dados
	// e nesse caso é para atualizar as informações de um determinado ID
	res, err := conn.Exec(`UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`, id, todo.Title, todo.Description, todo.Done)
	// tratamento de erro. Se for diferente de nulo, encerra a função retornando o err e zero porque não afetou nenhuma linha do bd
	if err != nil {
		return 0, err
	}

	// retorno das linhas afetadas e do valor de err, que nessa altura deve ser nil
	return res.RowsAffected()
}
