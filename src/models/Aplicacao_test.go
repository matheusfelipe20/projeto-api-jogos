package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

//Testes da aplicação

//Listar Campeonatos Disponiveis (Sucesso)
func TestListarCampeonatosDisponiveis(t *testing.T) {
	resp, err := http.Get("http://localhost:5000/campeonatos")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))

	if resp.StatusCode != 200 {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Listar Jogos Disponiveis (Sucesso)
func TestListarJogosDisponiveis(t *testing.T) {
	resp, err := http.Get("http://localhost:5000/jogos")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))

	if resp.StatusCode != 200 {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Criar bilhete de aposta (Sucesso)
func TestCriarVenda(t *testing.T) {

	bilhete := []byte(`{
    	"id_jogo": 354858757161272,
  		"titulo_jogo": "São Paulo x Flamengo",
  		"campeonato": "Brasileirão - Serie A",
  		"data_jogo": "2022-08-31",
  		"opcao_aposta": "casa",
  		"valor_aposta": 50,
  		"limite_aposta": 200,
  		"cliente_nome": "Bello Moreira Alcântara",
  		"cliente_cpf": "659.102.554-52",
  		"cliente_nascimento": "01/01/2000"
    	}`)

	resp, err := http.Post("http://localhost:5000/venda", "application/json",
		bytes.NewBuffer(bilhete))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))
	pro := Campeonato{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Listar Vendas realizadas (Sucesso)
func TestListarVendas(t *testing.T) {
	resp, err := http.Get("http://localhost:5000/venda")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))

	if resp.StatusCode != 200 {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Criar bilhete de aposta (Erro Data Passada)
func TestCriarVendaErro_Data(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/venda", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161273,
  			"titulo_jogo": "Fluminense x Palmeiras",
  			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-07-18",
  			"opcao_aposta": "fora",
  			"valor_aposta": 130,
  			"limite_aposta": 300,
  			"cliente_nome": "José",
  			"cliente_cpf": "231.300.114-80",
  			"cliente_nascimento": "02/06/1992"
    		}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))
	pro := Vendas{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Criar bilhete de aposta (Erro CPF invalido)
func TestCriarVendaErro_CPF(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/venda", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
			"titulo_jogo": "São Paulo x Flamengo",
 			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-08-31",
  			"opcao_aposta": "fora",
			"valor_aposta": 120,
			"limite_aposta": 300,
			"cliente_nome": "Valdir",
			"cliente_cpf": "502.098.729-08",
			"cliente_nascimento": "10/04/2000"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))
	pro := Vendas{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Criar bilhete de aposta (Erro Menor de idade)
func TestCriarVendaErro_DataNascimento(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/venda", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
			"titulo_jogo": "São Paulo x Flamengo",
 			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-08-31",
  			"opcao_aposta": "empate",
			"valor_aposta": 120,
			"limite_aposta": 300,
			"cliente_nome": "Valdir",
			"cliente_cpf": "231.300.114-80",
			"cliente_nascimento": "30/08/2006"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))
	pro := Vendas{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Criar bilhete de aposta (Erro Nome do cliente vazio)
func TestCriarVendaErro_NomeCliente(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/venda", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
			"titulo_jogo": "São Paulo x Flamengo",
 			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-08-31",
  			"opcao_aposta": "empate",
			"valor_aposta": 120,
			"limite_aposta": 300,
			"cliente_nome": "",
			"cliente_cpf": "231.300.114-80",
			"cliente_nascimento": "30/08/2001"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))
	pro := Vendas{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Criar bilhete de aposta (Erro Nome do campeonato vazio)
func TestCriarVendaErro_NomeCampeonato(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/venda", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
			"titulo_jogo": "São Paulo x Flamengo",
 			"campeonato": "",
  			"data_jogo": "2022-08-31",
  			"opcao_aposta": "empate",
			"valor_aposta": 120,
			"limite_aposta": 300,
			"cliente_nome": "Roberto",
			"cliente_cpf": "231.300.114-80",
			"cliente_nascimento": "30/08/2001"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))
	pro := Vendas{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Criar bilhete de aposta (Erro Nome do titulo vazio)
func TestCriarVendaErro_NomeTitulo(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/venda", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
			"titulo_jogo": "",
 			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-08-31",
  			"opcao_aposta": "empate",
			"valor_aposta": 120,
			"limite_aposta": 300,
			"cliente_nome": "Gabrielly",
			"cliente_cpf": "231.300.114-80",
			"cliente_nascimento": "30/08/2001"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))
	pro := Vendas{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Criar bilhete de aposta (Erro ID do jogo: 0)
func TestCriarVendaErro_IDjogo(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/venda", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 0,
			"titulo_jogo": "São Paulo x Flamengo",
 			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-08-31",
  			"opcao_aposta": "empate",
			"valor_aposta": 120,
			"limite_aposta": 300,
			"cliente_nome": "Helena",
			"cliente_cpf": "231.300.114-80",
			"cliente_nascimento": "30/08/2001"
    	}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))
	pro := Vendas{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}

//Listar Usuario por CPF(Sucesso)
func TestListarUsuarios(t *testing.T) {
	resp, err := http.Get("http://localhost:5000/usuarios/84280875472")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	log.Println(string(body))

	if resp.StatusCode != 200 {
		t.Errorf("Sem sucesso!! %v", string(body))
	}
}
