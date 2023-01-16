package handlers

import (
	"apigolang/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
// Função que representa o C do CRUD (create), que recebe como parâmetros uma Request (que vem do cliente), e um ResponseWriter, que 
// vai gravar a resposta para ser devolvida para o cliente
func Create(w http.ResponseWriter, r *http.Request) {
	// cria uma variável para guardar as informações do TODO
	var todo models.Todo

	// Cria um decoder com base no body da request, que é basicamente um json vindo do front-end, e executa a função decode, que retorna
	// um erro ou não (o erro retornado pode ser nil ou conter especificamente um erro, por isso é executado e o retorno guardado no err)
	err := json.NewDecoder(r.Body).Decode(&todo)

	// tratamento de erro, se err for diferente de nil, loga o erro e responde a request com um http.Error, passando como parâmetro o
	// ResponseWriter w (que é a gravação da resposta do cliente), a string de resposta para o cliente e o status code, que nesse caso
	// estamos escrevendo e enviando como resposta o status code 500 (internal server error, ou bad request)
	if err != nil{
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// se não tiver erro, chama a função de insert que criamos para inserir as informações recebidas no banco de dados
	// lembrando que passamos duas variáveis porque o retorno da função Insert tem dois retornos, sendo um id da nova inserção
	// e um erro, que pode ser nil ou não
	id, err := models.Insert(todo)

	// criada uma variável do tipo map de string e "dynamic" para guardar as informações da resposta a ser enviada para o front-end
	var resp map[string]any

	// tratamento de erro da inserção no banco de dados. Se houver erro, guarda na variável resp uma representação de um json com a chave
	// "Error", guardando um valor booleano true e uma chave "Message" que guarda uma string que está sendo gerada em tempo de execução
	// com a função Sprintf, onde estará escrita a frase abaixo, passando o erro que foi retornado pelo banco de dados
	if err != nil {
		resp = map[string]any{
			"Error": true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
	// Se tudo correr bem e o banco de dados conseguir gravar as informações passadas, guarda na variável resp uma representação de um json com a chave
	// "Error", guardando um valor booleano false e uma chave "Message" que guarda uma string que está sendo gerada em tempo de execução
	// com a função Sprintf, onde estará escrita a frase abaixo, passando a ID que foi retornado pelo banco de dados
		resp = map[string]any{
			"Error": false,
			"Message": fmt.Sprintf("To do inserido com sucesso! ID: %d", id),
		}
	}
	// Feito o tratamento de erro e salva a resposta na variável resp, solicita ao ResponseWriter w que adicione ao Header do HTTP que
	// o conteúdo do body é um JSON
	w.Header().Add("Content-Type", "application/json")
	// Chama a função NewEncoder que cria um encode no ResponseWriter w, e chama o Encode para codificar em JSON a variável resp e incluir
	// na stream (pelo que entendi ResponseWriter é uma stream)
	json.NewEncoder(w).Encode(resp)
}