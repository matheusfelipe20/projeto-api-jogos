package models

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

//Servidores_Terceiro
func ServCampeonatos() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dataCamp := []Campeonato{
			{
				ID:     30,
				Titulo: "Brasileirão - Serie A",
			},
			{
				ID:     35,
				Titulo: "Copa América - Feminina",
			},
			{
				ID:     36,
				Titulo: "Uruguai - Primeira Divisão",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(dataCamp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading campeonatos"))
		}
	})
}

func ServJogos() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dataJog := []Jogo{
			// Brasileirão - Serie A
			{
				ID:            354858757161272,
				Titulo:        "São Paulo x Flamengo",
				ID_Campeonato: 30,
				Data:          "2022-08-31",
			},
			{
				ID:            354858757161273,
				Titulo:        "Fluminense x Palmeiras",
				ID_Campeonato: 30,
				Data:          "2022-07-18",
			},
			{
				ID:            354858757161274,
				Titulo:        "Botafogo x Santos",
				ID_Campeonato: 30,
				Data:          "2022-07-15",
			},
			{
				ID:            354858757161275,
				Titulo:        "Vasco x Atlético",
				ID_Campeonato: 30,
				Data:          "2022-07-16",
			},
			{
				ID:            354858757161276,
				Titulo:        "Ceará x Avaí",
				ID_Campeonato: 30,
				Data:          "2022-07-22",
			},
			// Copa América - Feminina
			{
				ID:            354858324654689,
				Titulo:        "Colômbia x Chile",
				ID_Campeonato: 35,
				Data:          "2022-07-22",
			},
			{
				ID:            354858324654690,
				Titulo:        "Equador x Paraguai",
				ID_Campeonato: 35,
				Data:          "2022-07-15",
			},
			// Uruguai - Primeira Divisão
			{
				ID:            65489162165498,
				Titulo:        "Liverpool FC x AlbionFC",
				ID_Campeonato: 36,
				Data:          "2022-07-15",
			},
			{
				ID:            65489162165499,
				Titulo:        "Deportivo Maldonado x Torque da Cidade de Montevideu",
				ID_Campeonato: 36,
				Data:          "2022-07-18",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(dataJog); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading campeonatos"))
		}
	})
}

//Teste: Listar Campeonatos disponíveis (SUCESSO)
func Test_ServCampeonatos(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/campeonatos", ServCampeonatos())
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, _ := http.Get(ts.URL + "/campeonatos")
	body, _ := ioutil.ReadAll(res.Body)

	campeonatos := []byte(`[{"id":30,"titulo": "Brasileirão - Serie A"},{"id":35,"titulo": "Copa América - Feminina"},{"id":36,"titulo": "Uruguai - Primeira Divisão"}]`)

	eq, err := JSONBytesEqual(body, campeonatos)
	if err != nil {
		log.Println(err)
	}

	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, campeonatos)
	}

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

//Teste: Listar Jogos disponíveis (SUCESSO)
func Test_ServJogos(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/jogos", ServCampeonatos())
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/jogos")

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
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

//Criar bilhete de aposta (Erro limite do valor da aposta excedido)
func TestCriarVendaErro_LimiteExcedido(t *testing.T) {
	resp, err := http.Post("http://localhost:5000/venda", "application/json",
		bytes.NewBuffer([]byte(`{
			"id_jogo": 354858757161272,
			"titulo_jogo": "São Paulo x Flamengo",
 			"campeonato": "Brasileirão - Serie A",
  			"data_jogo": "2022-08-31",
  			"opcao_aposta": "empate",
			"valor_aposta": 500,
			"limite_aposta": 100,
			"cliente_nome": "Valdir",
			"cliente_cpf": "231.300.114-80",
			"cliente_nascimento": "15/02/2000"
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

//Funções para comparar os Json's

// JSONEqual comparando dois Json
func JSONEqual(a, b io.Reader) (bool, error) {
	var j, j2 interface{}
	d := json.NewDecoder(a)
	if err := d.Decode(&j); err != nil {
		return false, err
	}
	d = json.NewDecoder(b)
	if err := d.Decode(&j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

// JSONBytesEqual compara o JSON em fatias de dois bytes.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}
