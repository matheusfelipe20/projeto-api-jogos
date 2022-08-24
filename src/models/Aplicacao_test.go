package models

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
)

//TESTES CAMPEONATO
//Cadastrar Campeonato
func TestPostCampeonato(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/campeonatos", "application/json",
		bytes.NewBuffer([]byte(`{"id":1,"titulo":"Copa do Mundo"}`)))
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

	esperado := []byte(`{"id":1,"titulo":"Copa do Mundo"}`)
	//Aqui será comparado os dois Json e vai retornar um bool na variavel 'eq'
	eq, err := JSONBytesEqual(body, esperado)
	if err != nil {
		log.Println(err)
	}

	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, esperado)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Sem sucesso!! %v", string(body))
	}

}

//Teste para Listar Campeonatos disponiveis
func TestGetCampeonatoDisponivel(t *testing.T) {
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

	//O teste vai verificar se a lista está vazia ou não, se estiver vazia será adicionado um campeonato...
	esperado := []byte(`null`)
	eq, err := JSONBytesEqual(body, esperado)
	if err != nil {
		log.Println(err)
	}

	if eq == true {

		valoresCampeonato := []byte(`{"id":8,"titulo":"Brasileirão - serie B"}`)

		resp, err := http.Post("http://localhost:5000/campeonatos", "application/json",
			bytes.NewBuffer(valoresCampeonato))
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
}

//TESTES JOGOS (INCOMPLETO)

//Teste para Listar Jogos disponiveis
func TestGetJogosDisponiveis(t *testing.T) {
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

	//O teste vai verificar se a lista está vazia ou não, se estiver vazia será adicionado um jogo...
	esperado := []byte(`null`)
	eq, err := JSONBytesEqual(body, esperado)
	if err != nil {
		log.Println(err)
	}

	if eq == true {

		valorJogo := []byte(`{
    	"titulo": "Brasil x Sérvia",
    	"id_campeonato": 1,
    	"data": "2022-08-20 T16:00:00Z",
    	"opcoes": [
        	{ "1": 1.2 },
        	{ "x": 3.0 },
        	{ "2": 2.1 }
    	],
    	"limites": [
    	    { "1": 200 },
    	    { "x": 500 },
    	    { "2": 300 }
    	]}`)

		resp, err := http.Post("http://localhost:5000/jogos", "application/json",
			bytes.NewBuffer(valorJogo))
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
		pro := Jogo{}
		err = json.Unmarshal([]byte(string(body)), &pro)
		if err != nil {
			log.Println(err)
		}

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Sem sucesso!! %v", string(body))
		}
	}
}

//Cadastrar Jogos

func TestPostJogos(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/jogos", "application/json",
		bytes.NewBuffer([]byte(`{
			"id":5,
    		"titulo": "Brasil x Argentina",
    		"id_campeonato": 5,
    		"data": "2022-11-28 T13:30:00Z",
    		"opcoes": [
     		   	{ "1": 2.8 },
     		   	{ "x": 3.0 },
     		   	{ "2": 4.0 }
    		],
    		"limites": [
    		    { "1": 400 },
    		    { "x": 700 },
    		    { "2": 500 }
    		]}`)))
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

	esperado := []byte(`{
		"id":5,
    	"titulo": "Brasil x Argentina",
    	"id_campeonato": 5,
    	"data": "2022-11-28 T13:30:00Z",
    	"opcoes": [
        	{ "1": 2.8 },
        	{ "x": 3.0 },
        	{ "2": 4.0 }
    	],
    	"limites": [
    	    { "1": 400 },
    	    { "x": 700 },
    	    { "2": 500 }
    	]}`)
	//Aqui será comparado os dois Json e vai retornar um bool na variavel 'eq'
	eq, err := JSONBytesEqual(body, esperado)
	if err != nil {
		log.Println(err)
	}

	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, esperado)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Sem sucesso!! %v", string(body))
	}

}

//Testes de usuario

//Cadastrar Usuario (Error menor de idade)
func TestErroCadastroIdade(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/usuarios", "application/json",
		bytes.NewBuffer([]byte(`{"cpf":11223344556, "nome":"Matheus", "nascimento": "20-01-2005"}`)))
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
	pro := Usuario{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	esperado := []byte(`{"cpf":11223344556, "nome":"Matheus", "nascimento": "20-01-2005"}`)
	//Aqui será comparado os dois Json e vai retornar um bool na variavel 'eq'
	eq, err := JSONBytesEqual(body, esperado)
	if err != nil {
		log.Println(err)
	}

	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, esperado)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Sem sucesso!! %v", string(body))
	}

}

//Teste para verficar Erro: Cadastro de Usuario com CPF inválido (necessario 11 digitos)
func TestErroCadastroCPF(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/usuarios", "application/json",
		bytes.NewBuffer([]byte(`{"cpf":1, "nome":"Matheus", "nascimento": "20-01-2002"}`)))
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
	pro := Usuario{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	esperado := []byte(`{"cpf":1, "nome":"Matheus", "nascimento": "20-01-2002"}`)
	//Aqui será comparado os dois Json e vai retornar um bool na variavel 'eq'
	eq, err := JSONBytesEqual(body, esperado)
	if err != nil {
		log.Println(err)
	}

	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, esperado)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Sem sucesso!! %v", string(body))
	}

}

//Teste para verficar Erro: Campo de nome vazio
func TestErroCadastroNome(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/usuarios", "application/json",
		bytes.NewBuffer([]byte(`{"cpf":12312312345, "nome":"", "nascimento": "20-01-2002"}`)))
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
	pro := Usuario{}
	err = json.Unmarshal([]byte(string(body)), &pro)
	if err != nil {
		log.Println(err)
	}

	esperado := []byte(`{"cpf":12312312345, "nome":"", "nascimento": "20-01-2002"}`)
	//Aqui será comparado os dois Json e vai retornar um bool na variavel 'eq'
	eq, err := JSONBytesEqual(body, esperado)
	if err != nil {
		log.Println(err)
	}

	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, esperado)
	}

	if resp.StatusCode != http.StatusCreated {
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
