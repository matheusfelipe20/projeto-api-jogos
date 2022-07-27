package controllers

import "github.com/matheusfelipe20/projeto-api-jogos/src/models"

var Campeonatos []models.Campeonato = []models.Campeonato{
	models.Campeonato{
		ID:     10,
		Titulo: "Brasileirão - Serie A",
	},
	models.Campeonato{
		ID:     15,
		Titulo: "Copa do Brasil",
	},
	models.Campeonato{
		ID:     20,
		Titulo: "Copa Sul-Americana",
	},
}
