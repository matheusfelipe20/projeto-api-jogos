package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/db"
	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/repositories"
)

// CadastrarCampeonato cadastra um campeonato no banco de dados
func CadastrarCampeonato(w http.ResponseWriter, r *http.Request) {

	// ler o corpo da requisição
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// converter o corpo da requisição para um objeto
	var camp models.Campeonato
	if err := json.Unmarshal(requestBody, &camp); err != nil {
		log.Fatal(err)
	}

	// abrindo conexão com o banco de dados
	db, err := db.Conectar()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// insere um campeonato no banco de dados
	repositorio := repositories.NovoRepositorioDeCampeonatos(db)
	lastId, err := repositorio.AdicionarCampeonato(camp)
	if err != nil {
		log.Printf("Erro ao fazer o insert no banco de dados: %v, id: %d", err, lastId)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Campeonato criado com sucesso. ID: %d", lastId)))
}

// ListarCampeonatos lista todos os campeonatos
func ListarCampeonatos(w http.ResponseWriter, r *http.Request) {

	// abrindo conexão com o banco de dados
	db, err := db.Conectar()
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	// acessa o repositorio de campeonatos para fazer a busca
	repositorio := repositories.NovoRepositorioDeCampeonatos(db)
	campeonatos, err := repositorio.BuscarCampeonatos()
	if err != nil {
		log.Printf("Erro ao fazer o select no banco de dados: %v", err)
		log.Fatal(err)
	}

	log.Println(campeonatos)
}
