package validators

// Verifica se o ID é diferente de zero
func ValidadeID(id uint64) bool {
	idJogo := id

	return idJogo != 0
}

// Verifica se o campo do título do jogo está vazio
func ValidarTitulo(titulo string) bool {
	tituloJogo := titulo

	return tituloJogo != ""
}

// Verifica se o ID do campeonato é diferente de zero
func ValidarIDCamp(campeonato uint64) bool {
	idCamp := campeonato

	return idCamp != 0
}

// Verifica se a data do jogo está vazia
func ValidarData(data string) bool {
	dataJogo := data

	return dataJogo != ""
}
