package controllers

import (
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/respostas"
)

var Jogos []models.Jogo = []models.Jogo{
	// Brasileirão - Serie A
	{
		ID:            354858757161272,
		Titulo:        "São Paulo x Flamengo",
		ID_Campeonato: 30,
		Data:          "2022-08-31",
		Opcoes:        mapOpcoes(2.0, 3.5, 1.6),
		Limite:        mapLimite(900.0, 400.0, 350.0),
	},
	{
		ID:            354858757161273,
		Titulo:        "Fluminense x Palmeiras",
		ID_Campeonato: 30,
		Data:          "2022-07-18",
		Opcoes:        mapOpcoes(3.6, 5.5, 2.6),
		Limite:        mapLimite(500.0, 700.0, 550.0),
	},
	{
		ID:            354858757161274,
		Titulo:        "Botafogo x Santos",
		ID_Campeonato: 30,
		Data:          "2022-07-15",
		Opcoes:        mapOpcoes(2.0, 4.0, 3.6),
		Limite:        mapLimite(800.0, 200.0, 550.0),
	},
	{
		ID:            354858757161275,
		Titulo:        "Vasco x Atlético",
		ID_Campeonato: 30,
		Data:          "2022-07-16",
		Opcoes:        mapOpcoes(3.2, 2.5, 1.6),
		Limite:        mapLimite(950.0, 800.0, 750.0),
	},
	{
		ID:            354858757161276,
		Titulo:        "Ceará x Avaí",
		ID_Campeonato: 30,
		Data:          "2022-07-22",
		Opcoes:        mapOpcoes(1.8, 2.5, 1.6),
		Limite:        mapLimite(750.0, 300.0, 550.0),
	},
	// Copa América - Feminina
	{
		ID:            354858324654689,
		Titulo:        "Colômbia x Chile",
		ID_Campeonato: 35,
		Data:          "2022-07-22",
		Opcoes:        mapOpcoes(2.3, 5.5, 3.6),
		Limite:        mapLimite(880.0, 330.0, 400.0),
	},
	{
		ID:            354858324654690,
		Titulo:        "Equador x Paraguai",
		ID_Campeonato: 35,
		Data:          "2022-07-15",
		Opcoes:        mapOpcoes(2.6, 3.1, 4.2),
		Limite:        mapLimite(300.0, 200.0, 150.0),
	},
	// Uruguai - Primeira Divisão
	{
		ID:            65489162165498,
		Titulo:        "Liverpool FC x AlbionFC",
		ID_Campeonato: 36,
		Data:          "2022-07-15",
		Opcoes:        mapOpcoes(2.7, 3.0, 4.6),
		Limite:        mapLimite(970.0, 400.0, 950.0),
	},
	{
		ID:            65489162165499,
		Titulo:        "Deportivo Maldonado x Torque da Cidade de Montevideu",
		ID_Campeonato: 36,
		Data:          "2022-07-18",
		Opcoes:        mapOpcoes(2.0, 3.5, 1.6),
		Limite:        mapLimite(700.0, 500.0, 450.0),
	},
}

// ListarJogos é a função que irá listar os jogos
func ListarJogos(w http.ResponseWriter, r *http.Request) {

	respostas.JSON(w, http.StatusOK, Jogos)

}

func mapOpcoes(a float64, b float64, c float64) []map[string]float64 {
	opcoes := []map[string]float64{}

	opcoes = append(opcoes, map[string]float64{"casa": a})
	opcoes = append(opcoes, map[string]float64{"empate": b})
	opcoes = append(opcoes, map[string]float64{"fora": c})

	return opcoes
}

func mapLimite(a float64, b float64, c float64) []map[string]float64 {
	limite := []map[string]float64{}

	limite = append(limite, map[string]float64{"casa": a})
	limite = append(limite, map[string]float64{"empate": b})
	limite = append(limite, map[string]float64{"fora": c})

	return limite
}
