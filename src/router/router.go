package router

import (
	"github.com/gorilla/mux"
	"github.com/matheusfelipe20/projeto-api-jogos/src/router/rotas"
)

// GerarRota irá retornar uma rota configurada
func GerarRota() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return rotas.Configurar(r)
}
