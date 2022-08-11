CREATE TABLE IF NOT EXISTS jogos(
    id int primary key not null,
    titulo varchar(50) not null,
    id_campeonato int not null,
    data date not null
);

CREATE TABLE IF NOT EXISTS campeonatos(
    id int primary key not null,
    titulo varchar(100) not null
);

CREATE TABLE IF NOT EXISTS usuarios(
    id int primary key not null,
    cpf int not null,
    nome varchar(100) not null,
    nascimento varchar(100) not null
);