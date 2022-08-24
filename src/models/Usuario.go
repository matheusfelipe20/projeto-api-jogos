package models

import (
	"errors"

	"github.com/matheusfelipe20/projeto-api-jogos/src/models/validators"
)

type Usuario struct {
	ID         uint64 `json:"id,omitempty"`
	Cpf        uint64 `json:"cpf,omitempty"`
	Nome       string `json:"nome,omitempty"`
	Nascimento string `json:"nascimento,omitempty"`
}

// Preparar vai chamar ps métodos para validar e formatar o usuário
// func (usuario *Usuario) Preparar() error {
// 	if erro := validators.ValidadeDataNascimento(""); erro != nil {
// 		return erro
// 	}
// 	return nil
// }

// Verifica as validações dos campos de cadastro de um usuário
func (usr *Usuario) ValidarUsuario() error {

	if verificarNascimento := validators.ValidadeDataNascimento(usr.Nascimento); !verificarNascimento {
		return errors.New("Falha ao cadastrar, usuário menor de idade")
	}

	if verficarCPF := validators.VerificarCPF(int(usr.Cpf)); !verficarCPF {
		return errors.New("Falha ao cadastrar, usuário menor de idade")
	}

	if verficarNome := validators.ValidadeNome(usr.Nome); !verficarNome {
		return errors.New("Falha ao cadastrar, campo nome vazio")
	}

	return nil
}
