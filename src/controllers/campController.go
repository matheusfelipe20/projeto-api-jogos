package controllers

import (
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/respostas"
)

var Campeonatos []models.Campeonato = []models.Campeonato{
	{
		ID:     30,
		Titulo: "Brasileirão - Serie A",
	},
	{
		ID:     35,
		Titulo: "Copa América - Feminina",
	},
	{
		ID:     36,
		Titulo: "Uruguai - Primeira Divisão",
	},
}

// ListarCampeonatos lista todos os campeonatos
func ListarCampeonatos(w http.ResponseWriter, r *http.Request) {

	respostas.JSON(w, http.StatusOK, Campeonatos)
}
