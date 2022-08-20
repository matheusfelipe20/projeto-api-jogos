-- tabela de campeonatos

CREATE SEQUENCE campeonato_id_seq
    START WITH 1
    INCREMENT BY 1
    MINVALUE 1
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS campeonatos(
    id serial primary key not null,
    titulo varchar(100) not null
);

ALTER TABLE campeonatos ALTER COLUMN id SET DEFAULT nextval('campeonato_id_seq');

-- tabela de jogos

CREATE SEQUENCE jogo_id_seq
    START WITH 1
    INCREMENT BY 1
    MINVALUE 1
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS jogos(
	id bigint primary key,
	titulo varchar(100) not null,
	id_campeonato bigint not null,
	data_jogo varchar(50) not null,
    FOREIGN KEY (id_campeonato) REFERENCES campeonatos(id)
);

ALTER TABLE jogos ALTER COLUMN id SET DEFAULT nextval('jogo_id_seq');

-- tabela de usuarios

CREATE SEQUENCE usuario_id_seq
    START WITH 1
    INCREMENT BY 1
    MINVALUE 1
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS usuarios(
    id bigserial primary key not null,
    cpf int not null,
    nome varchar(100) not null,
    nascimento varchar(50) not null
);

ALTER TABLE usuarios ALTER COLUMN id SET DEFAULT nextval('usuario_id_seq');