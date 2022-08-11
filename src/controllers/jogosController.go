package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"

	"github.com/matheusfelipe20/projeto-api-jogos/src/db"
	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/repositories"
)

// CadastrarJogo é a função que irá cadastrar um jogo
func CadastrarJogo(w http.ResponseWriter, r *http.Request){
	// ler o corpo da requisição
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// converter o corpo da requisição para um objeto
	var jg models.Jogo
	if err := json.Unmarshal(requestBody, &jg); err != nil {
		log.Fatal(err)
	}

	// abre a conexão com o banco de dados
	db, err := db.Conectar()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// insere o jogo no banco de dados
	repositorio := repositories.NovoRepositorioDeJogos(db)
	lastId, err := repositorio.Criar(jg)
	if err != nil {
		log.Printf("Erro ao fazer o insert no banco de dados: %v, id: %d", err, lastId)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Jogo criado com sucesso. ID: %d", lastId)))
}

// ListarJogos é a função que irá listar os jogos
func ListarJogos(w http.ResponseWriter, r *http.Request) {

	// abrindo conexão com o banco de dados
	db, err := db.Conectar()
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	// acessa o repositorio de jogos para fazer a busca
	repositorio := repositories.NovoRepositorioDeJogos(db)
	jogos, err := repositorio.BuscarJogos()
	if err != nil {
		log.Fatalf("Erro ao buscar jogos: %v", err)
	}

	log.Println(jogos)
	
}