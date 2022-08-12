package rotas

import (
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/controllers"
)

var rotasJogos = []Rota{
	{
		URI:    "/campeonatos",
		Metodo: http.MethodGet,
		Funcao: controllers.ListarCampeonatos,
	},
	{
		URI:    "/campeonatos",
		Metodo: http.MethodPost,
		Funcao: controllers.CadastrarCampeonato,
	},
	{
		URI:    "/jogos",
		Metodo: http.MethodGet,
		Funcao: controllers.ListarJogos,
	},
	{
		URI:    "/jogos",
		Metodo: http.MethodPost,
		Funcao: controllers.CadastrarJogo,
	},
	{
		URI:    "/jogos/{id}",
		Metodo: http.MethodGet,
		Funcao: controllers.ListarJogosByID,
	},
	{
		URI:    "/jogos/{data}",
		Metodo: http.MethodGet,
		Funcao: controllers.ListarUsuario, // fazer mudanças aqui
	},
	{
		URI:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CadastrarUsuario,
	},
	{
		URI:    "/usuarios/{cpf}",
		Metodo: http.MethodGet,
		Funcao: controllers.ListarUsuario,
	},
}
