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
		URI:    "/jogos/datas/{data_jogo}",
		Metodo: http.MethodGet,
		Funcao: controllers.ListarJogosByData, 
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
	{
		URI:    "/aposta/",
		Metodo: http.MethodPost,
		Funcao: controllers.RealizarAposta,
	},
	{
		URI:    "/aposta/{nome}",
		Metodo: http.MethodGet,
		Funcao: controllers.ListarApostasByComprador,
	},
}
