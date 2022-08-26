package models

type Aposta struct {
	Id         int     `json:"id,omitempty"`
	Id_jogo    int     `json:"id_jogo,omitempty"`
	Id_usuario int     `json:"id_usuario,omitempty"`
	Opcao      string  `json:"opcao_aposta,omitempty"`
	Valor      float64 `json:"valor_aposta,omitempty"`
}
