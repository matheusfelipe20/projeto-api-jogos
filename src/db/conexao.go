package db

import (
	"database/sql"
	"github.com/matheusfelipe20/projeto-api-jogos/src/configs"
	
	_ "github.com/lib/pq"	// driver de conexão do postgres
)

// Conectar abre a conexão com o postgres
func Conectar() (*sql.DB, error){

	db, err := sql.Open("postgres", configs.StringConexaoBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}