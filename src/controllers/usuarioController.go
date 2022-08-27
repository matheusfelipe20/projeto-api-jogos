package controllers

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

// CadastrarUsuario irá cadastrar um usuário no banco de dados
func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	// Ler o corpo da requisição
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Converter o corpo da requisição para um objeto
	var usr models.Usuario
	if err := json.Unmarshal(requestBody, &usr); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	// Verifica os dados de cadastro do usuário
	if err = usr.ValidarUsuario(); err != nil {
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

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usr.ID, err = repositorio.AdicionarUsuario(usr)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, usr)
	w.Write([]byte(fmt.Sprintf("Usuário criado com sucesso. ID: %d", usr.ID)))
}

// ListarUsuario irá buscar um usuário através do CPF no banco de dados
func ListarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioCPF, err := strconv.Atoi(parametros["cpf"])
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
	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usr, err := repositorio.BuscarUsuarioByCpf(usuarioCPF)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, usr)

}
*/
