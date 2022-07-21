package router

import (
	"github.com/matheusfelipe20/projeto-api-jogos/src/router/rotas"
	"github.com/gorilla/mux"
)

// GerarRota irá retornar uma rota configurada
func GerarRota() *mux.Router{
	r := mux.NewRouter()

	return rotas.Configurar(r)
}