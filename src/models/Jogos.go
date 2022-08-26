package models

import (
	"errors"

	"github.com/matheusfelipe20/projeto-api-jogos/src/models/validators"
)

type Jogo struct {
	ID            uint64               `json:"id,omitempty"`
	Titulo        string               `json:"titulo,omitempty"`
	ID_Campeonato uint64               `json:"id_campeonato,omitempty"`
	Data          string               `json:"data_jogo,omitempty"`
	Opcoes        []map[string]float64 `json:"opcoes,omitempty"`
	Limite        []map[string]int     `json:"limite,omitempty"`

// Verifica as validações dos campos de cadastro de um jogo
func (jg *Jogo) ValidarJogo() error {

	if verificarID := validators.ValidadeID(uint64(jg.ID)); !verificarID {
		return errors.New("Falha ao cadastrar jogo com o ID zero")
	}

	if verficarTitulo := validators.ValidarTitulo(jg.Titulo); !verficarTitulo {
		return errors.New("Falha ao cadastrar, insira o título do jogo")
	}

	if ValidarIDCampeonato := validators.ValidarIDCamp(jg.ID_Campeonato); !ValidarIDCampeonato {
		return errors.New("Falha ao cadastrar jogo, o ID do campeonado não pode ser zero")
	}

	if ValidarDataJogo := validators.ValidarData(jg.Data); !ValidarDataJogo {
		return errors.New("Falha ao cadastrar, insira a data do jogo")
	}

	return nil
}
