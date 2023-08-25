<div align="center">
    <img alt="Golang Logo" title="Golang Logo" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="300" height="300">
</div>

# OPPORTUNITIES - REST API - GOLANG

## Libs utilizadas

- [GIN Framework HTTP](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [GORM Driver PostgreSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)
- [Swaggo](https://github.com/swaggo/swag)
- [Docker](https://www.docker.com/)

## Variáveis de Ambiente

Para rodar esse projeto, você vai precisar adicionar as seguintes variáveis de ambiente no seu .env

`PORT`

`DB_USER`

`DB_PASSWORD`

`DB_NAME`

## Funcionalidades

- [x] Cadastro de oportunidade;
- [x] Busca uma oportunidade pelo ID;
- [x] Busca todas oportunidades;
- [x] Atualiza uma oportunidade;
- [x] Deleta uma oportunidade;

## Endpoints

#### Swagger Documentation

```http
  GET /swagger/index.html
```

#### Criar uma vaga

```http
  POST /opportunity
```

| Parâmetro  | Tipo     | Descrição                           |
| :--------- | :------- | :---------------------------------- |
| `Role`     | `string` | **Obrigatório**. Nome da vaga       |
| `Company`  | `string` | **Obrigatório**. Empresa da vaga    |
| `Location` | `string` | **Obrigatório**. Local de trabalho  |
| `Remote`   | `bool`   | **Obrigatório**. Se é remoto ou não |
| `Link`     | `string` | **Obrigatório**. Link da vaga       |
| `Salary`   | `int64`  | **Obrigatório**. Salário            |

#### Buscar todas

```http
  GET /opportunities
```

#### Buscar uma pelo id

```http
  GET /opportunity
```

_Query_
| Parâmetro | Tipo | Descrição |
| :-------- | :----- | :-------- |
| `id` | `string` | id da vaga |

#### Atualizar uma vaga

```http
  GET /opportunity
```

_Query_
| Parâmetro | Tipo | Descrição |
| :-------- | :----- | :-------- |
| `id` | `string` | id da vaga |

_Body_
| Parâmetro | Tipo | Descrição |
| :-------- | :------- | :-------------------------------- |
| `Role` | `string` | Nome da vaga |
| `Company` | `string` | Empresa da vaga |
| `Location` | `string` | Local de trabalho |
| `Remote` | `bool` | Se é remoto ou não |
| `Link` | `string` | Link da vaga |
| `Salary` | `int64` | Salário |

#### Deletar uma vaga

```http
  GET /opportunity
```

_Query_
| Parâmetro | Tipo | Descrição |
| :-------- | :----- | :-------- |
| `id` | `string` | id da vaga |

- Exemplo de Payload:

```json
{
  "role": "Pleno Software Engineer",
  "company": "FC Lab",
  "location": "SP",
  "remote": true,
  "link": "vaga.com/vaga",
  "salary": 15000
}
```

- Exemplo de Requisição:

```sh
curl -X POST \
  http://localhost:8080/opportunity \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{
		"role": "Pleno Software Engineer",
	  "company": "FC Lab",
	  "location": "SP",
	  "remote": true,
	  "link": "vaga.com/vaga",
	  "salary": 15000
}'
```
