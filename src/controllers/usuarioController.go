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
	"github.com/matheusfelipe20/projeto-api-jogos/src/respostas"
)

func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usr models.Usuario
	if err := json.Unmarshal(requestBody, &usr); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usr.ID, err = repositorio.AdicionarUsuario(usr)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, usr)
	w.Write([]byte(fmt.Sprintf("Usuário criado com sucesso. ID: %d", usr.ID)))
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
