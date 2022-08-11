package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matheusfelipe20/projeto-api-jogos/src/db"
	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/repositories"
)

func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usr models.Usuario
	if err := json.Unmarshal(requestBody, &usr); err != nil {
		log.Fatal(err)
	}

	db, err := db.Conectar()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	lastID, err := repositorio.AdicionarUsuario(usr)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário criado com sucesso. ID: %d", lastID)))
}

func ListarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	usuarioCPF, err := strconv.Atoi(parametro["cpf"])
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.Conectar()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usr, err := repositorio.BuscarUsuarioByCpf(usuarioCPF)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(usr)

}
