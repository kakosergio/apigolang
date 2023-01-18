package handlers

import (
	"apigolang/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))

	if err != nil {
		log.Printf("Erro ao excluir o registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Erro: foram excluídos %v registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "dados excluídos com sucesso!",
	}

	// Feito o tratamento de erro e salva a resposta na variável resp, solicita ao ResponseWriter w que adicione ao Header do HTTP que
	// o conteúdo do body é um JSON
	w.Header().Add("Content-Type", "application/json")
	// Chama a função NewEncoder que cria um encode no ResponseWriter w, e chama o Encode para codificar em JSON a variável resp e incluir
	// na stream (pelo que entendi ResponseWriter é uma stream)
	json.NewEncoder(w).Encode(resp)
}
