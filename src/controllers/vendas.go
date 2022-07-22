package controllers

import "net/http"

// GetCPF busca os dados de um usuário pelo CPF salvo no banco de dados
func GetCPF(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando os dados do usuário!"))
}

// GetCampeonatos busca os campeonatos salvo no banco
func GetCampeonatos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Campeonatos!"))
}

// GetJogos busca os jogos salvo no banco
func GetJogos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Jogos!"))
}
