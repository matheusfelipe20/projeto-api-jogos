package repositories

import (
	"database/sql"

	"github.com/matheusfelipe20/projeto-api-jogos/src/models"
)

// usuarios representa repositório de usuarios
type usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um novo repositório
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// AdicionarUsuarios irá adicionar um usuário no banco de dados
func (repositorio usuarios) AdicionarUsuario(usuario models.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare("insert into usuarios (cpf, nome, nascimento) values ($1, $2, $3) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id int
	err = statement.QueryRow(usuario.Cpf, usuario.Nome, usuario.Nascimento).Scan(&id)
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

// BuscarUsuarioByCpf irá resgatar usuário pelo cpf
func (repositorio usuarios) BuscarUsuarioByCpf(cpf int) (models.Usuario, error) {
	linhas, err := repositorio.db.Query("select * from usuarios where cpf = $1", cpf)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linhas.Close()

	var usr models.Usuario

	if linhas.Next() {
		if err := linhas.Scan(&usr.ID, &usr.Cpf, &usr.Nome, &usr.Nascimento); err != nil {
			return models.Usuario{}, err
		}
	}
	return usr, nil
}
