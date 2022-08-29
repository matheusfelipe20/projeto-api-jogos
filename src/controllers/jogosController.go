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
	// Brasileirão - Serie A
	{
		ID:            354858757161272,
		Titulo:        "São Paulo x Flamengo",
		ID_Campeonato: 30,
		Data:          "2022-08-31",
	},
	{
		ID:            354858757161273,
		Titulo:        "Fluminense x Palmeiras",
		ID_Campeonato: 30,
		Data:          "2022-07-18",
	},
	{
		ID:            354858757161274,
		Titulo:        "Botafogo x Santos",
		ID_Campeonato: 30,
		Data:          "2022-07-15",
	},
	{
		ID:            354858757161275,
		Titulo:        "Vasco x Atlético",
		ID_Campeonato: 30,
		Data:          "2022-07-16",
	},
	{
		ID:            354858757161276,
		Titulo:        "Ceará x Avaí",
		ID_Campeonato: 30,
		Data:          "2022-07-22",
	},
	// Copa América - Feminina
	{
		ID:            354858324654689,
		Titulo:        "Colômbia x Chile",
		ID_Campeonato: 35,
		Data:          "2022-07-22",
	},
	{
		ID:            354858324654690,
		Titulo:        "Equador x Paraguai",
		ID_Campeonato: 35,
		Data:          "2022-07-15",
	},
	// Uruguai - Primeira Divisão
	{
		ID:            65489162165498,
		Titulo:        "Liverpool FC x AlbionFC",
		ID_Campeonato: 36,
		Data:          "2022-07-15",
	},
	{
		ID:            65489162165499,
		Titulo:        "Deportivo Maldonado x Torque da Cidade de Montevideu",
		ID_Campeonato: 36,
		Data:          "2022-07-18",
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
