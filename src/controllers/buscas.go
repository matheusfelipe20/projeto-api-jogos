package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCPF busca os dados de um usuário pelo CPF salvo no banco de dados
func GetCPF(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	cpf, err := strconv.ParseUint(vars["numero"], 10, 64)
	if err != nil {
		http.Error(w, "Cpf inválido!", http.StatusBadRequest)
		return
	}

	for _, usuario := range Usuarios {
		if usuario.Cpf == uint64(cpf) {
			json.NewEncoder(w).Encode(usuario)
			return
		}
	}
	/*
		db, erro := db.Conectar()
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
		defer db.Close()
	*/
	w.WriteHeader(http.StatusNotFound)
}

// GetCampeonatos busca os campeonatos salvo no banco
func GetCampeonatos(w http.ResponseWriter, r *http.Request) {

	//Validação para verificar se há ou não Campeonatos disponíveis
	validateVazio := Campeonatos
	var verificationLen = len(validateVazio)
	if verificationLen >= 1 {
		w.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(Campeonatos)
	} else {
		w.Write([]byte("Sem Campeonatos disponíveis no momento!"))
	}
	/*
		db, erro := db.Conectar()
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
		defer db.Close()
	*/
}

// GetJogos busca os jogos salvo no banco
func GetJogos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Jogos!"))
}
