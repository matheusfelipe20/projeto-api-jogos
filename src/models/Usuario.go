package models

type Usuario struct {
	ID         uint64 `json:"id,omitempty"`
	Cpf        uint64 `json:"cpf,omitempty"`
	Nome       string `json:"nome,omitempty"`
	Nascimento string `json:"nascimento,omitempty"`
}

type Usuarios struct {
	Usuarios []Usuario `json:"usuarios"`
}

// Preparar vai chamar ps métodos para validar e formatar o usuário
// func (usuario *Usuario) Preparar() error {
// 	if erro := validators.ValidadeDataNascimento(""); erro != nil {
// 		return erro
// 	}
// 	return nil
// }
