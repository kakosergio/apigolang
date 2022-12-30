package models

import "apigolang/db"

// cria a função insert para inserir informações no banco de dados
func Insert(todo Todo) (id int64, err error) {
	// abre conexão com o db
	conn, err := db.OpenConnection()

	// verifica se houve erro
	if err != nil {
		return
	}
	// fecha a conexão assim que todo o código for executado (para isso serve a palavra-chave DEFER, igual DIFERIR pagamento de custas a final)
	defer conn.Close()

	// cria uma query para incluir as informações no banco de dados e pede para retornar o ID gerado automaticamente
	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}