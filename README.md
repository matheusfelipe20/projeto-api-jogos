<h1 align="center">Projeto Api Jogos</h1>

<h4 align="center"> 
	:wrench:🚧 Aplicação em desenvolvimento... 🚧:wrench:
</h4>

## 🚀 Tecnologias
Esse projeto foi desenvolvido utilizando as seguintes tecnologias:

- Golang
- PostgreSQL
- Gorilla-mux

---

## 💻 Descrição

- O projeto tem como objetivo simular um ambiente de apostas esportivas com o uso de uma api com integração de terceiros, com as opções de fazer aposta e realizar consultas. Ele usa [gorilla-mux](https://github.com/gorilla/mux) para criar as rotas;

- A funcionalidade desse projeto consiste em realizar apostas: `venda`;
- Consultar todas as apostas realizadas: `venda`;
- Filtrar por: `campeonatos` e `jogos`;
- Filtrar usuários por: `cpf`.

---

## Consultas

#### Consultar Apostas
- GET `/venda`
- Output
```json
[
  {
    "id": 1,
    "id_jogo": 354858757161272,
    "titulo_jogo": "São Paulo x Flamengo",
    "campeonato": "Brasileirão - Serie A",
    "data_jogo": "2022-08-29",
    "opcao_aposta": "casa",
    "valor_aposta": 50,
    "limite_aposta": 500,
    "cliente_nome": "Jonas Victor Alves da Silva",
    "cliente_cpf": "23130011480",
    "cliente_nascimento": "02-03-2000"
  }
]
```

#### Consultar Campeonatos

- GET `/campeonatos`
- Output
```json
[
  {
    "id": 30,
    "titulo": "Brasileirão - Serie A"
  },
  {
    "id": 35,
    "titulo": "Copa América - Feminina"
  },
  {
    "id": 36,
    "titulo": "Uruguai - Primeira Divisão"
  }
]
```

#### Consultar Jogos
- GET `/jogos`
- Output
```json
[
  {
    "id": 354858757161272,
    "titulo": "São Paulo x Flamengo",
    "id_campeonato": 30,
    "data_jogo": "2022-08-31"
  },
  {
    "id": 354858757161273,
    "titulo": "Fluminense x Palmeiras",
    "id_campeonato": 30,
    "data_jogo": "2022-07-18"
  },
  {
    "id": 354858757161274,
    "titulo": "Botafogo x Santos",
    "id_campeonato": 30,
    "data_jogo": "2022-07-15"
  },
  {
    "id": 354858757161275,
    "titulo": "Vasco x Atlético",
    "id_campeonato": 30,
    "data_jogo": "2022-07-16"
  },
  {
    "id": 354858757161276,
    "titulo": "Ceará x Avaí",
    "id_campeonato": 30,
    "data_jogo": "2022-07-22"
  },
  {
    "id": 354858324654689,
    "titulo": "Colômbia x Chile",
    "id_campeonato": 35,
    "data_jogo": "2022-07-22"
  },
  {
    "id": 354858324654690,
    "titulo": "Equador x Paraguai",
    "id_campeonato": 35,
    "data_jogo": "2022-07-15"
  },
  {
    "id": 65489162165498,
    "titulo": "Liverpool FC x AlbionFC",
    "id_campeonato": 36,
    "data_jogo": "2022-07-15"
  },
  {
    "id": 65489162165499,
    "titulo": "Deportivo Maldonado x Torque da Cidade de Montevideu",
    "id_campeonato": 36,
    "data_jogo": "2022-07-18"
  }
]
```

#### Consultar Usuários por CPF
- GET `/usuarios/65910255452`
- CPFs disponíveis:
    - 65910255452
    - 84280875472
    - 43717395475
    - 23130011480
- Output
```json
null
```

---

### Como executar o teste do projeto:

- O teste se encontra na pasta: 'models', nome do arquivo: "Aplicacao_test.go"
- Diretorio: `projeto-api-jogos/src/models/Aplicacao_test.go`

