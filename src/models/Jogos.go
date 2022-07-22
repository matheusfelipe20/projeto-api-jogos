package models

type Jogo struct {
	ID             uint64   `json:"id,omitempty"`
	Titulo         string   `json:"titulo,omitempty"`
	ID_Campeonatos uint64   `json:"campeonatos,omitempty"`
	Data           string   `json:"data,omitempty"`
	Opcoes         []string `json:"opcoes,omitempty"`
	Limites        []string `json:"limites,omitempty"`
}

type Jogos struct {
	Jogos []Jogo `json:"jogos"`
}
