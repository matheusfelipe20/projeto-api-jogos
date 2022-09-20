package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/controllers/services"
	"github.com/matheusfelipe20/projeto-api-jogos/src/db"
	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
	"github.com/matheusfelipe20/projeto-api-jogos/src/repositories"
	"github.com/matheusfelipe20/projeto-api-jogos/src/respostas"
)

// RealizarVenda irá registrar uma venda no banco de dados
func RealizarVenda(w http.ResponseWriter, r *http.Request) {

	// Lendo o corpo da requisição
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Criando um objeto do tipo venda
	var venda models.Vendas
	err = json.Unmarshal(requestBody, &venda)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	// adicionando as operações de apostas
	listaJogos := Jogos
	for _, jogo := range listaJogos {
		if jogo.ID == venda.Id_jogo {

			_, value := services.CompararOpcaoAposta(venda.Opcao_aposta, jogo.Opcoes) // opção de odds do jogo
			venda.Opcao_valor = value	// atribuição do valor da opção de aposta a venda
			val := services.CompararLimiteComOpcao(venda.Opcao_aposta, jogo.Limite) // limite de cada aposta do jogo
			venda.Limite_aposta = val	// atribuição do limite de aposta a venda

			ganho := services.CalcularGanho(venda.Valor_aposta, venda.Opcao_valor)
			venda.Ganho_provavel = ganho	// atribuição do ganho provável a venda
			break
		}
	}


	// fazendo a validação da venda
	val := venda.ValidarVendas()
	if val != nil {
		respostas.Erro(w, http.StatusBadRequest, val)
		return
	}

	// Abrindo conexão com o banco de dados
	db, err := db.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Registrando o objeto venda no repositório
	repositorio := repositories.NovoRepositorioDeVendas(db)
	venda.Id, err = repositorio.CriarVenda(venda)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resultadoOK := "Venda realizada com sucesso!"
	respostas.JSON(w, http.StatusCreated, resultadoOK)
}

// ListarVendas irá buscar todas as vendas no banco de dados
func ListarVendas(w http.ResponseWriter, r *http.Request) {

	// Abrindo conexão com o banco de dados
	db, err := db.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Buscando todas as vendas no repositório
	repositorio := repositories.NovoRepositorioDeVendas(db)
	vendas, err := repositorio.BuscarVendas()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, vendas)
}
