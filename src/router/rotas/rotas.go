package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}

// Configura as rotas
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasJogos

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
