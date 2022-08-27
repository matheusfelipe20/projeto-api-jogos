package rotas

import (
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/controllers"
	// "github.com/matheusfelipe20/projeto-api-jogos/src/controllers"
)

var rotasJogos = []Rota{
	{
		URI:    "/campeonatos",
		Metodo: http.MethodGet,
		Funcao: controllers.ListarCampeonatos,
	},
	{
		URI:    "/jogos",
		Metodo: http.MethodGet,
		Funcao: controllers.ListarJogos,
	},
	{
		URI:    "/usuarios/{cpf}",
		Metodo: http.MethodGet,
		Funcao: func(h http.ResponseWriter, r *http.Request) {},
	},
	{
		URI:    "/venda/",
		Metodo: http.MethodPost,
		Funcao: func(h http.ResponseWriter, r *http.Request) {},
	},
}
