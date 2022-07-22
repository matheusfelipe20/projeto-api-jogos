package models

type Jogo struct {
	ID             uint64 `json:"id,omitempty"`
	Titulo         string `json:"titulo,omitempty"`
	ID_Campeonatos uint64 `json:"campeonatos,omitempty"`
	Data           string `json:"data,omitempty"`
}

type Opcoes struct {
	Vitoria string `json:"1,omitempty"`
	Empate  string `json:"x,omitempty"`
	Derrota string `json:"2,omitempty"`
}

type Limites struct {
	Vitoria string `json:"1,omitempty"`
	Empate  string `json:"x,omitempty"`
	Derrota string `json:"2,omitempty"`
}

type Jogos struct {
	Jogos []Jogo `json:"jogos"`
}
