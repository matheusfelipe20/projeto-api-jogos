package controllers

import (
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/respostas"
)

var Usuario []models.Usuario = []models.Usuario{
	{
		Cpf:        "65910255452",
		Nome:       "Bello Moreira Alcântara",
		Nascimento: "01-01-2000",
	},
	{
		Cpf:        "84280875472",
		Nome:       "Chico Amaro Bugalho",
		Nascimento: "02-02-1997",
	},
	{
		Cpf:        "43717395475",
		Nome:       "Erick Vasconcelos Pessoa",
		Nascimento: "03-03-1995",
	},
	{
		Cpf:        "23130011480",
		Nome:       "Jonas Victor Alves da Silva",
		Nascimento: "02-03-2000",
	},
}

// ListarUsuario irá buscar um usuário através do CPF no banco de dados
func ListarUsuario(w http.ResponseWriter, r *http.Request) {

	respostas.JSON(w, http.StatusOK, Usuario)

}
