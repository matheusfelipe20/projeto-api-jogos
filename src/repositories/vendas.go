package repositories
/*

import (
	"database/sql"

	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
)


type aposta struct{
	db *sql.DB
}

// NovoRepositorioDeAposta cria um novo repositório de aposta
func NovoRepositorioDeAposta(db *sql.DB) *aposta{
	return &aposta{db: db}
}

// CriarAposta irá inserir uma nova aposta no banco de dados
func (ap *aposta) CriarAposta(aposta models.Aposta) (int, error){

	stmt, err := ap.db.Prepare("insert into aposta (id_jogo, id_usuario, opcao_aposta, valor_aposta) values ($1,$2,$3,$4) RETURNING id");
	if err != nil{
		return 0, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(aposta.Id_jogo, aposta.Id_usuario, aposta.Opcao, aposta.Valor).Scan(&id)
	if err != nil{
		return 0, err
	}

	return id, nil
}

// BuscarApostaByNomeComprador irá buscar uma aposta pelo nome do comprador
func (ap *aposta) BuscarApostaByNomeComprador(nome string) ([]interface{}, error){

	
	linhas, err := ap.db.Query("select u.nome, j.titulo, a.opcao_aposta, a.valor_aposta from aposta a inner join jogos j on a.id_jogo = j.id inner join usuarios u  on a.id_usuario = u.id where u.nome  = $1", nome)
	if err != nil{
		return nil, err		
	}
	defer linhas.Close()

	var apst []interface{} 

	for linhas.Next(){
		var usr models.Usuario
		var jg models.Jogo
		var ap models.Aposta

		if err = linhas.Scan(&usr.Nome, &jg.Titulo, &ap.Opcao, &ap.Valor); err != nil{
			return nil, err
		}

		apst = append(apst, usr, jg, ap)

		
	}

	return apst, nil
}

// BuscarApostaByTituloJogo irá buscar uma aposta pelo titulo do jogo
func (ap *aposta) BuscarApostaByTituloJogo(titulo string) ([]models.Aposta, error){

	search := "select u.nome, j.titulo, j.data_jogo, a.opcao_aposta, a.valor_aposta from aposta ap inner join jogos jg on ap.id_jogo = jg.id inner join usuarios u on ap.id_usuario = u.id where jg.titulo = $1"
	
	linhas, err := ap.db.Query(search, titulo)
	if err != nil{
		return nil, err
	}
	defer linhas.Close()

	var aps []models.Aposta

	for linhas.Next(){
		var aposta models.Aposta

		if err = linhas.Scan(&aposta.Id, &aposta.Id_jogo, &aposta.Id_usuario, &aposta.Opcao, &aposta.Valor); err != nil{
			return nil, err
		}

		aps = append(aps, aposta)
	}

	return aps, nil
}

// BuscarApostaByCampeonato irá buscar uma aposta pelo campeonato */