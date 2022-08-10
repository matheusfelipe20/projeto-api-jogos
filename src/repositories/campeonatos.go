package repositories

import (
	"database/sql"

	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
)

type campeonatos struct {
	db *sql.DB
}

// NovoRepositorioDeCampeonatos cria um novo repositório de campeonatos
func NovoRepositorioDeCampeonatos(db *sql.DB) *campeonatos {
	return &campeonatos{db: db}
}

// AdicionarCampeonato irá adicionar um novo campeonato no banco de dados
func (repositorio campeonatos) AdicionarCampeonato(campeonato models.Campeonato) (uint64, error) {

	statement, err := repositorio.db.Prepare("insert into campeonatos (id, titulo) values ($1,$2) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id int
	err = statement.QueryRow(campeonato.ID, campeonato.Titulo).Scan(&id)
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

// BuscarCampeonatos irá fazer uma busca em todos os campeonatos no banco de dados
func (repositorio campeonatos) BuscarCampeonatos() ([]models.Campeonato, error) {

	linhas, err := repositorio.db.Query("select * from campeonatos")
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var camps []models.Campeonato

	for linhas.Next() {
		var campeonato models.Campeonato

		if err = linhas.Scan(&campeonato.ID, &campeonato.Titulo); err != nil {
			return nil, err
		}

		camps = append(camps, campeonato)
	}

	return camps, nil
}
