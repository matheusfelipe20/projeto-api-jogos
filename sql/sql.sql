-- tabela de campeonatos

CREATE TABLE IF NOT EXISTS campeonatos(
    id bigserial primary key not null,
    titulo varchar(100) not null
);

-- tabela de jogos

CREATE TABLE IF NOT EXISTS jogos(
	id bigserial primary key not null,
	titulo varchar(100) not null,
	id_campeonato bigint not null,
	data_jogo varchar(50) not null,
    FOREIGN KEY (id_campeonato) REFERENCES campeonatos(id)
);

-- tabela de usuarios

CREATE TABLE IF NOT EXISTS usuarios(
    id bigserial primary key not null,
    cpf bigint unique not null,
    nome varchar(100) not null,
    nascimento varchar(50) not null
);