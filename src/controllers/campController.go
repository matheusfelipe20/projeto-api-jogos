package controllers

import (
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/respostas"
)

/*

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/db"
	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/repositories"
	"github.com/matheusfelipe20/projeto-api-jogos/src/respostas"
)
*/

var Campeonatos []models.Campeonato = []models.Campeonato{
	{
		ID:     1,
		Titulo: "Brasileirão - Serie A",
	},
	{
		ID:     2,
		Titulo: "Copa Sul-Americana",
	},
	{
		ID:     3,
		Titulo: "Champions League",
	},
}

// ListarCampeonatos lista todos os campeonatos
func ListarCampeonatos(w http.ResponseWriter, r *http.Request) {

	respostas.JSON(w, http.StatusOK, Campeonatos)

	// Abrindo conexão com o banco de dados
	/*
		Erro(w, http.StatusInternalServerError, err)
					return
				}
				defer db.Close()

				// Acessa o repositorio de campeonatos para fazer a busca
				repositorio := repositories.NovoRepositorioDeCampeonatos(db)
				campeonatos, err := repositorio.BuscarCampeonatos()
				if err != nil {
					respostas.Erro(w, http.StatusInternalServerError, err)
					return
				}
	*/
}

/*

// CadastrarCampeonato cadastra um campeonato no banco de dados
func CadastrarCampeonato(w http.ResponseWriter, r *http.Request) {
	// Ler o corpo da requisição
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Converter o corpo da requisição para um objeto
	var camp models.Campeonato
	if err := json.Unmarshal(requestBody, &camp); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	// Abrindo conexão com o banco de dados
	db, err := db.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Insere um campeonato no banco de dados
	repositorio := repositories.NovoRepositorioDeCampeonatos(db)
	camp.ID, err = repositorio.AdicionarCampeonato(camp)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, camp)
}
*/
