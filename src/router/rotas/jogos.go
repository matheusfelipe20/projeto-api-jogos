package rotas

import "net/http"

var rotasJogos = []Rota{
	{
		URI:             "/campeonatos",
		Metodo:          http.MethodGet, 
		Funcao:          func(w http.ResponseWriter, r *http.Request) {},
	},
	{
		URI:             "/jogos",
		Metodo:          http.MethodGet, 
		Funcao:          func(w http.ResponseWriter, r *http.Request) {},
	},
}