package models

import "apigolang/db"

// Função que recebe um ID e exclui as informações desse ID do database
// Retorna um int e um erro, indicando quais linhas foram afetadas
func Delete(id int64) (int64, error) {
	// abre a conexão com o db
	conn, err := db.OpenConnection()

	// trata o erro. Se tiver erro, encerra a chamada à função
	if err != nil {
		return 0, err
	}
	// difere o fechamento da conexão a final
	defer conn.Close()

	// executa o DELETE no banco de dados e retorna as linhas afetadas e se houve algum erro
	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
	// se houve erro, encerra a função retornando err e no lugar das linhas afetadas retorna zero
	if err != nil {
		return 0, err
	}

	// retorna as linhas afetadas e err, que deverá aqui ser nil
	return res.RowsAffected()
}
