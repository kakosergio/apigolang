package handlers

import (
	"apigolang/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r *http.Request) {
	// cria duas variáveis para guardar o parse da Request a fim de recuperar o id, e um possível erro
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	// tratamento de erro, caso o id seja nulo
	if err != nil {
		log.Printf("Erro ao realizar o parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// cria duas variáveis para gravar as informações recuperadas
	todo, err := models.Get(int64(id))

	// faz o tratamento de erro para retornar caso haja falha de comunicação com o banco de dados
	if err != nil {
		log.Printf("Erro ao recuperar o registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// adiciona um header http à resposta w, indicando o tipo de conteúdo, que no caso presente é um application/json
	w.Header().Add("Content-Type", "application/json")
	// Chama a função NewEncoder que cria um encode no ResponseWriter w, e chama o Encode para codificar em JSON a variável todo e incluir
	// na stream (pelo que entendi ResponseWriter é uma stream)
	json.NewEncoder(w).Encode(todo)
}
