package models

// Struct utilizada para guardar informações que irão para o banco de dados, em formato JSON
type Todo struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done"`
}