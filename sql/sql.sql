CREATE TABLE IF NOT EXISTS jogos(
    id int primary key auto_increment,
    titulo varchar(50) not null,
    id_campeonato int not null,
    data date not null
);

CREATE TABLE IF NOT EXISTS campeonatos(
    id int primary key auto_increment,
    titulo varchar(100) not null
);