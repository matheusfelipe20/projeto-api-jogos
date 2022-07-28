package models

import (
	"errors"
	"log"
	"time"
)

type Usuario struct {
	Cpf        uint64 `json:"cpf,omitempty"`
	Nome       string `json:"nome,omitempty"`
	Nascimento string `json:"nascimento,omitempty"`
}

/*
type Usuarios struct {
	Usuarios []Usuario `json:"usuarios"`
}
*/

// Preparar vai chamar ps métodos para validar e formatar o usuário
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validateUser(""); erro != nil {
		return erro
	}

	return nil
}

func (u *Usuario) validateUser(action string) error {

	dataUser := u.Nascimento

	parsed, err := time.Parse("02/01/2006", dataUser)
	if err != nil {
		log.Fatal(err)
	}
	beforeYear := parsed.AddDate(18, 0, 0) //Somar 18 anos a data de nascimento do usuario
	today := time.Now()

	compareted := beforeYear.Before(today)
	if !compareted {
		return errors.New("o usuario requer ser maior de 18 anos")
	}
  // " " test
	if u.Cpf == 0 { 
		return errors.New("requer o número do CPF do usuario")
	}
	if u.Nome == "" {
		return errors.New("requer o nome do usuario")
	}
	if u.Nascimento == "" {
		return errors.New("requer a data de nascimento do usuario")
	}

	return nil
}
