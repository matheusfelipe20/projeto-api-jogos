package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/db"
	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/repositories"
	"github.com/matheusfelipe20/projeto-api-jogos/src/respostas"
)

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

// ListarJogos é a função que irá listar os jogos
func ListarJogos(w http.ResponseWriter, r *http.Request) {

	// Abrindo conexão com o banco de dados
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
	respostas.JSON(w, http.StatusOK, jogos)

}
