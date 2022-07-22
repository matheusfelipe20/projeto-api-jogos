package models

type Campeonato struct {
	ID     uint64 `json:"id,omitempty"`
	Titulo string `json:"titulo,omitempty"`
}

type Campeonatos struct {
	Campeonatos []Campeonato `json:"campeonatos"`
}
