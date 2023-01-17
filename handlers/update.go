package handlers

import (
	"log"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi .URLParam(r, "id"))
	teste := 0

	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}