package main

import (
	"net/http"
)

// Função que inicializa nossa aplicação
func main() {
	// Cria um novo router utilizando a biblioteca padrão
	router := http.NewServeMux()
	// Roteia os requests no path /helloworld para a HandlerFunc helloWorld
	router.HandleFunc("/helloworld", helloWorld)
	// Roteia os requests no path /ping para o Handler pingHandler
	ph := &pingHandler{}
	router.Handle("/ping", ph)
	// Inicia nosso servidor HTTP na porta 8080 utilizando o router criado acima
	http.ListenAndServe(":8080", router)
}

// Em Go temos um tipo de interface chamado HandlerFunc que aceita um ResponseWriter e um Request
// Essa interface é utilizada pela maioria dos routers criados pela comunidade
func helloWorld(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Hello world!"))
}

// Cria o tipo pingHandler que irá implementar a interface http.Handler
type pingHandler struct{}

// adiciona o método ServeHTTP no tipo pingHandler para que ele implemente a interface
func (*pingHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("pong"))
}
