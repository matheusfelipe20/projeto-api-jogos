package controllers
/*

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusfelipe20/projeto-api-jogos/src/db"
	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/repositories"
	"github.com/matheusfelipe20/projeto-api-jogos/src/respostas"
)

// RealizarAposta é a função que irá cadastrar uma aposta
func RealizarAposta(w http.ResponseWriter, r *http.Request) {
	
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Converter o corpo da requisição para um objeto
	var ap models.Aposta
	if err := json.Unmarshal(requestBody, &ap); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}


	// abre a conexão com o banco de dados
	db, err := db.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Insere a aposta no banco de dados
	repositorio := repositories.NovoRepositorioDeAposta(db)
	ap.Id, err = repositorio.CriarAposta(ap)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, ap)
	w.Write([]byte(fmt.Sprintf("Aposta realizada com sucesso. ID: %d", ap.Id)))
}

// ListarApostasByComprador irá listar apostas feitas por um apostador
func ListarApostasByComprador(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	// Abrindo conexão com o banco de dados
	db, err := db.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeAposta(db)
	apostas, err := repositorio.BuscarApostaByNomeComprador(parametros["nome"])
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, apostas)
} */