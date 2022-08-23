package repositories

import (
	"database/sql"

	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
)

type jogos struct {
	db *sql.DB
}

// NovoRepositorioDeJogos cria um novo repositório de jogos
func NovoRepositorioDeJogos(db *sql.DB) *jogos {
	return &jogos{db: db}
}

// Criar irá inserir um novo jogo no banco de dados
func (repositorio jogos) Criar(jogo models.Jogo) (uint64, error) {

	statement, err := repositorio.db.Prepare("insert into jogos (titulo, id_campeonato, data_jogo) values ($1,$2,$3) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id int
	err = statement.QueryRow(jogo.Titulo, jogo.ID_Campeonato, jogo.Data).Scan(&id)
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

// BuscarJogos irá listar todos os jogos do banco de dados
func (repositorio jogos) BuscarJogos() ([]models.Jogo, error) {

	linhas, err := repositorio.db.Query("select * from jogos")
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var jgs []models.Jogo

	for linhas.Next() {
		var jogo models.Jogo

		if err = linhas.Scan(&jogo.ID, &jogo.Titulo, &jogo.ID_Campeonato, &jogo.Data); err != nil {
			return nil, err
		}

		jgs = append(jgs, jogo)
	}

	return jgs, nil
}

// BuscarJogoByID irá buscar um jogo pelo seu ID
func (repositorio jogos) BuscarJogoByID(id int) (models.Jogo, error) {
	linhas, err := repositorio.db.Query("select * from jogos where id = $1", id)
	if err != nil {
		return models.Jogo{}, err
	}
	defer linhas.Close()

	var jgID models.Jogo

	if linhas.Next() {
		if err := linhas.Scan(&jgID.ID, &jgID.Titulo, &jgID.ID_Campeonato, &jgID.Data); err != nil {
			return models.Jogo{}, err
		}
	}
	return jgID, nil
}

// BuscarJogoByData irá listar todos os jogos que possuirem a data informada
func (repositorio jogos) BuscarJogosByData(data string) ([]models.Jogo, error) {

	linhas, err := repositorio.db.Query("select * from jogos where data_jogo = $1", data)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var jgs []models.Jogo

	for linhas.Next() {
		var jogo models.Jogo

		if err = linhas.Scan(&jogo.ID, &jogo.Titulo, &jogo.ID_Campeonato, &jogo.Data); err != nil {
			return nil, err
		}

		jgs = append(jgs, jogo)
	}

	return jgs, nil
}
