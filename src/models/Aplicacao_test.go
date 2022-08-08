package models

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestUserNotFound(t *testing.T) {
	resp, err := http.Get("http://localhost:5000/cpf")
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

	if resp.StatusCode != 404 {
		t.Fatalf("bad status code: %d", resp.StatusCode)
	}
}

//Teste para Listar Usuario existente a partir do CPF
func TestGetUserTESTE(t *testing.T) {
	resp, err := http.Get("http://localhost:5000/cpf/55512377788")
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

	esperado := []byte(`{"cpf":55512377788,"nome":"felipe","nascimento":"02/02/2022"}`)

	//Aqui será comparado os dois Json e vai retornar um bool na variavel 'eq'
	eq, err := JSONBytesEqual(body, esperado)
	if err != nil {
		log.Println(err)
	}

	if !eq {
		t.Errorf("Sem sucesso!! valor recebido: '%s', valor esperado: '%s'", body, esperado)
	}

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
