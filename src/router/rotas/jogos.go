package rotas

import (
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/controllers"
)

var rotasJogos = []Rota{
	{
		URI:    "/campeonatos",
		Metodo: http.MethodGet,
		Funcao: controllers.GetCampeonatos,
	},
	{
		URI:    "/jogos",
		Metodo: http.MethodGet,
		Funcao: controllers.GetJogos,
	},
	{
		URI:    "/cpf/{numero}",
		Metodo: http.MethodGet,
		Funcao: controllers.GetCPF,
	},
}
