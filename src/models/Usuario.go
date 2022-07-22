package models

type Usuario struct {
	Cpf        string `json:"cpf,omitempty"`
	Nome       string `json:"nome,omitempty"`
	Nascimento string `json:"nascimento,omitempty"`
}

type Usuarios struct {
	Usuarios []Usuario `json:"usuarios"`
}
