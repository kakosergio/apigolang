package main

import (
	"apigolang/configs"
	"apigolang/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main () {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)
}