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

- O projeto tem como objetivo simular um ambiente de apostas esportivas com o uso de uma api com integração de terceiros, com as opções de cadastrar e realizar consultas. Ele usa [gorilla-mux](https://github.com/gorilla/mux) para criar as rotas.

- A funcionalidade desse projeto consiste em cadastrar: `usuarios`, `campeonatos` e `jogos`. 
- Consultar todos os eventos cadastrados ou filtrar por: `id`, `campeonatos`e `datas`. 
- E filtrar usuários por: `cpf`.

---

## Consultas:

#### Consultar Campeonatos

- GET `/campeonatos`
```json
[
  {
    "id": 1,
    "titulo": "Brasileirão - Serie A"
  },
  {
    "id": 2,
    "titulo": "Copa Sul-Americana"
  },
  {
    "id": 3,
    "titulo": "Champions League"
  }
]
```

#### Consultar Jogos
- GET `/jogos`
```json
[
  {
    "id": 10,
    "titulo": "Flamengo x Athletico-PR",
    "id_campeonato": 1,
    "data_jogo": "01-01-2022"
  },
  {
    "id": 11,
    "titulo": "Palmeiras x Corinthians",
    "id_campeonato": 1,
    "data_jogo": "02-02-2022"
  },
  {
    "id": 12,
    "titulo": "Internacional x Botafogo",
    "id_campeonato": 1,
    "data_jogo": "03-03-2022"
  },
  {
    "id": 20,
    "titulo": "Fluminense x River Plate",
    "id_campeonato": 2,
    "data_jogo": "10-04-2022"
  },
  {
    "id": 21,
    "titulo": "Santos x Banfield",
    "id_campeonato": 2,
    "data_jogo": "11-05-2022"
  },
  {
    "id": 22,
    "titulo": "Lanús x Melgar",
    "id_campeonato": 2,
    "data_jogo": "12-06-2022"
  },
  {
    "id": 30,
    "titulo": "Real Madrid x Liverpool",
    "id_campeonato": 3,
    "data_jogo": "29-07-2022"
  },
  {
    "id": 31,
    "titulo": "PSG x Manchester City",
    "id_campeonato": 3,
    "data_jogo": "29-08-2022"
  },
  {
    "id": 32,
    "titulo": "Atlético de Madrid x Porto",
    "id_campeonato": 3,
    "data_jogo": "29-09-2022"
  }
]
```

#### Consultar Jogos por ID
- GET `/jogos/10`
- Jogos disponíveis:
    - `10`, `11`, `12`
    - `20`, `21`, `22`
    - `30`, `31`, `32`
```json
{
  "id": 10,
  "titulo": "Flamengo x Athletico-PR",
  "id_campeonato": 1,
  "data_jogo": "01-01-2022"
}
```

#### Consultar Jogos por datas
- GET `/jogos/datas/29-07-2022`
```json
null
```

#### Consultar Usuários por CPF
- GET `/usuarios/65910255452`
- CPFs disponíveis:
    - _
    - _
    - _
```json
null
```

---

### Como executar o teste do projeto:

- O teste se encontra na pasta: 'models', nome do arquivo: "Aplicacao_test.go"
- Diretorio: `projeto-api-jogos/src/models/Aplicacao_test.go`

