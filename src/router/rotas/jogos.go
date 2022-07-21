package rotas

import (
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/router/models"
)

var rotasJogos = []Rota{
	{
		URI:    "/campeonatos",
		Metodo: http.MethodGet,
		Funcao: models.GetCampeonatos,
	},
	{
		URI:    "/jogos",
		Metodo: http.MethodGet,
		Funcao: models.GetJogos,
	},
	{
		URI:    "/cpf/{numero}",
		Metodo: http.MethodGet,
		Funcao: models.GetCPF,
	},
}
