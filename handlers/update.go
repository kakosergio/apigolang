package handlers

import (
	"apigolang/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)

	// tratamento de erro, se err for diferente de nil, loga o erro e responde a request com um http.Error, passando como parâmetro o
	// ResponseWriter w (que é a gravação da resposta do cliente), a string de resposta para o cliente e o status code, que nesse caso
	// estamos escrevendo e enviando como resposta o status code 500 (internal server error, ou bad request)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)

	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Erro: foram atualizados %v registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "dados atualizados com sucesso!",
	}
	
	// Feito o tratamento de erro e salva a resposta na variável resp, solicita ao ResponseWriter w que adicione ao Header do HTTP que
	// o conteúdo do body é um JSON
	w.Header().Add("Content-Type", "application/json")
	// Chama a função NewEncoder que cria um encode no ResponseWriter w, e chama o Encode para codificar em JSON a variável resp e incluir
	// na stream (pelo que entendi ResponseWriter é uma stream)
	json.NewEncoder(w).Encode(resp)
}
