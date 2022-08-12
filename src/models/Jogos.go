package models

import (
	"errors"
	"log"
	"time"
)

type Jogo struct {
	ID             uint64 `json:"id,omitempty"`
	Titulo         string `json:"titulo,omitempty"`
	ID_Campeonatos uint64 `json:"id_campeonato,omitempty"`
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

func (j *Jogo) validadeVenda(action string) error {

	dataJogo := j.Data

	parsed, err := time.Parse("02/01/2006", dataJogo)
	if err != nil {
		log.Fatal(err)
	}
	beforeDay := parsed.AddDate(0, 0, 1) //Somar 1 dia a data do evento
	today := time.Now()

	comparetedEvento := beforeDay.Before(today)
	if !comparetedEvento {
		return errors.New("falha ao realizar a ação. A posta deve ser feita até um dia antes do evento")
	}

	return nil
}
