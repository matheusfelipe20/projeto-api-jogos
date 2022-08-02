CREATE TABLE IF NOT EXISTS api_jogos(
    id int primary key auto_increment,
    titulo varchar(50) not null,
    id_campeonato int unique not null,
    data date not null,
);