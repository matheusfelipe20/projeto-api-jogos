-- tabela de apostas 

create table if not exists venda(
	id bigserial primary key not null,
	id_jogo bigint not null,
    titulo_jogo varchar(50) not null,
	campeonato varchar(100) not null,
    data_jogo varchar(50) not null,
	opcao_aposta varchar(11) not null,
	valor_aposta float not null,
    limite_aposta float not null,
    cliente_nome varchar(100) not null,
    cliente_cpf varchar(50) not null,
    cliente_nascimento varchar(50) not null,
    check (opcao_aposta in('casa', 'empate', 'fora'))
);
