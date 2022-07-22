package db

import (
	"database/sql"
	"fmt"

	"github.com/matheusfelipe20/projeto-api-jogos/src/configs"
	_ "github.com/lib/pq"	// driver de conexão do postgres
)

// Conectar abre a conexão com o postgres
func Conectar() (*sql.DB, error){

	conf := configs.GetDB()

	stringConexao := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", stringConexao)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err

}