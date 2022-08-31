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

- O projeto tem como objetivo simular um ambiente de apostas esportivas com o uso de uma api com integração de terceiros, com as opções de fazer aposta e realizar consultas. Ele usa [gorilla-mux](https://github.com/gorilla/mux) para criar as rotas.

- A funcionalidade desse projeto consiste em realizar apostas: `venda`. 
- Consultar todas as apostas realizadas: `venda`.
- Filtrar por: `campeonatos`, `jogos`.
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
    "id": 1,
    "titulo": "Flamengo x Athletico-PR",
    "id_campeonato": 1,
    "data_jogo": "01-01-2022"
  },
  {
    "id": 2,
    "titulo": "Palmeiras x Corinthians",
    "id_campeonato": 1,
    "data_jogo": "02-02-2022"
  },
  {
    "id": 3,
    "titulo": "Internacional x Botafogo",
    "id_campeonato": 1,
    "data_jogo": "03-03-2022"
  },
  {
    "id": 4,
    "titulo": "Fluminense x River Plate",
    "id_campeonato": 2,
    "data_jogo": "10-04-2022"
  },
  {
    "id": 5,
    "titulo": "Santos x Banfield",
    "id_campeonato": 2,
    "data_jogo": "11-05-2022"
  },
  {
    "id": 6,
    "titulo": "Lanús x Melgar",
    "id_campeonato": 2,
    "data_jogo": "12-06-2022"
  },
  {
    "id": 7,
    "titulo": "Real Madrid x Liverpool",
    "id_campeonato": 3,
    "data_jogo": "29-07-2022"
  },
  {
    "id": 8,
    "titulo": "PSG x Manchester City",
    "id_campeonato": 3,
    "data_jogo": "29-08-2022"
  },
  {
    "id": 9,
    "titulo": "Atlético de Madrid x Porto",
    "id_campeonato": 3,
    "data_jogo": "29-09-2022"
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

