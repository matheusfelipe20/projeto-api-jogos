package models
type Jogo struct {
	ID            uint64               `json:"id,omitempty"`
	Titulo        string               `json:"titulo,omitempty"`
	ID_Campeonato uint64               `json:"id_campeonato,omitempty"`
	Data          string               `json:"data_jogo,omitempty"`
	Opcoes        []map[string]float64 `json:"opcoes,omitempty"`
	Limite        []map[string]int     `json:"limite,omitempty"`
}
