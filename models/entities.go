package models

// cria struct para guardar informações que irão para o banco de dados
type Todo struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done"`
}