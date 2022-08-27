package models

type Vendas struct {
	Id                 int     `json:"id,omitempty"`
	Id_jogo            int     `json:"id_jogo,omitempty"`
	Campeonato         string  `json:"campeonato,omitempty"` // pode ser um id_campeonato
	Data_jogo          string  `json:"data_jogo,omitempty"`
	Opcao_aposta       string  `json:"opcao_aposta,omitempty"`
	Valor_aposta       float64 `json:"valor_aposta,omitempty"`
	Cliente_nome       string  `json:"cliente_nome,omitempty"`
	Cliente_cpf        string  `json:"cliente_cpf,omitempty"`
	Cliente_nascimento string  `json:"cliente_nascimento,omitempty"`
}
