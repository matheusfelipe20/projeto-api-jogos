package controllers

import (
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/respostas"
)

/*

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matheusfelipe20/projeto-api-jogos/src/db"
	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/repositories"
	"github.com/matheusfelipe20/projeto-api-jogos/src/respostas"
)
*/

var Jogos []models.Jogo = []models.Jogo{
	{
		ID:            10,
		Titulo:        "Flamengo x Athletico-PR",
		ID_Campeonato: 1,
		Data:          "01-01-2022",
	},
	{
		ID:            11,
		Titulo:        "Palmeiras x Corinthians",
		ID_Campeonato: 1,
		Data:          "02-02-2022",
	},
	{
		ID:            12,
		Titulo:        "Internacional x Botafogo",
		ID_Campeonato: 1,
		Data:          "03-03-2022",
	},

	{
		ID:            20,
		Titulo:        "Fluminense x River Plate",
		ID_Campeonato: 2,
		Data:          "10-04-2022",
	},
	{
		ID:            21,
		Titulo:        "Santos x Banfield",
		ID_Campeonato: 2,
		Data:          "11-05-2022",
	},
	{
		ID:            22,
		Titulo:        "Lanús x Melgar",
		ID_Campeonato: 2,
		Data:          "12-06-2022",
	},

	{
		ID:            30,
		Titulo:        "Real Madrid x Liverpool",
		ID_Campeonato: 3,
		Data:          "29-07-2022",
	},
	{
		ID:            31,
		Titulo:        "PSG x Manchester City",
		ID_Campeonato: 3,
		Data:          "29-08-2022",
	},
	{
		ID:            32,
		Titulo:        "Atlético de Madrid x Porto",
		ID_Campeonato: 3,
		Data:          "29-09-2022",
	},
}

// ListarJogos é a função que irá listar os jogos
func ListarJogos(w http.ResponseWriter, r *http.Request) {

	respostas.JSON(w, http.StatusOK, Jogos)

	// Abrindo conexão com o banco de dados
	/*
		db, err := db.Conectar()
		if err != nil {
			respostas.Erro(w, http.StatusInternalServerError, err)
			return
		}
		defer db.Close()

		// Acessa o repositorio de jogos para fazer a busca
		repositorio := repositories.NovoRepositorioDeJogos(db)
		jogos, err := repositorio.BuscarJogos()
		if err != nil {
			respostas.Erro(w, http.StatusInternalServerError, err)
			return
		}
	*/

}

/*
// CadastrarJogo é a função que irá cadastrar um jogo
func CadastrarJogo(w http.ResponseWriter, r *http.Request) {
	// Ler o corpo da requisição
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Converter o corpo da requisição para um objeto
	var jg models.Jogo
	if err := json.Unmarshal(requestBody, &jg); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	// Verifica os dados de cadastro do jogo
	if err = jg.ValidarJogo(); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	// Abre a conexão com o banco de dados
	db, err := db.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Insere o jogo no banco de dados
	repositorio := repositories.NovoRepositorioDeJogos(db)
	jg.ID, err = repositorio.Criar(jg)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, jg)
	w.Write([]byte(fmt.Sprintf("Jogo criado com sucesso. ID: %d", jg.ID)))
}

// ListarJogosByID irá buscar um jogo através do ID no banco de dados
func ListarJogosByID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	jogoID, err := strconv.Atoi(parametros[ID])
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	// Abre a conexão com o banco de dados
	db, err := db.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Acessa o repositorio de jogos para fazer a busca
	repositorio := repositories.NovoRepositorioDeJogos(db)
	jgID, err := repositorio.BuscarJogoByID(jogoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, jgID)
}

// ListarJogosByData é a função que irá listar os jogos pela data
func ListarJogosByData(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	// Abre a conexão com o banco de dados
	db, err := db.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Acessa o repositorio de jogos para fazer a busca
	repositorio := repositories.NovoRepositorioDeJogos(db)
	jg, err := repositorio.BuscarJogosByData(parametros[Data])
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, jg)

}
*/
