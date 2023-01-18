package handlers

import (
	"apigolang/models"
	"encoding/json"
	"log"
	"net/http"
)
// Função que recupera uma lista de todas as tarefas cadastradas no banco de dados e retorna ela em formato JSON
// Por curiosidade, o primeiro parâmetro http.ResponseWriter funciona como um return da função, mesmo ela sendo void (sem retorno), pois
// é através dele que o http retorna a resposta para o cliente que fez a requisição (Request). Por isso que a última função chamada no
// código é o json.NewEncoder(w), que cria um encode das tarefas (todos) e 'escreve' a lista de tarefas no response, em formato JSON
// sem a necessidade de transformar em uma variável 'map' (chave e valor)
func List(w http.ResponseWriter, r *http.Request) {
	// cria duas variáveis para guardar os retornos da função GetAll
	todos, err := models.GetAll()

	// trata o erro e retorna uma mensagem caso não consiga recuperar os dados do banco de dados
	if err != nil {
		log.Printf("Erro ao recuperar os dados: %v", err)
	}
	// adiciona um header http à resposta w, indicando o tipo de conteúdo, que no caso presente é um application/json
	w.Header().Add("Content-Type", "application/json")
	// faz um encode da lista de tarefas (todos) em formato json e 'escreve' no ResponseWriter w para retorno ao requisitante
	json.NewEncoder(w).Encode(todos)
}
