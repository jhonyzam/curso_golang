# Products API

Este projeto tem como objetivo exemplificar como é o desenvolvimento de uma api em GO,
de forma simples e didática.

### Pré requisitos
    * Docker
    * Docker compose

Comando para subir a aplicação em segundo plano:
````
docker compose up -d
````

Método    | URL    | Body
--------- | ------ | ------
GET     | http://localhost:8080/v1/products/:id | `null`
POST    | http://localhost:8080/v1/products | `{ "name": "Produto post", "price": 1.00 }`
PUT     | http://localhost:8080/v1/products/:id | `{ "name": "Produto put", "price": 2.00 }`
DELETE  | http://localhost:8080/v1/products/:id | `null`

#### Collection postman: [Postman collection](https://github.com/jhonyzam/curso_golang/tree/aula_8/aula_8/products/postman_collection.json)

### Rodar Testes
````
POSTGRES_URL=postgres://postgres:admin@localhost:5432/products_test?sslmode=disable go test -v -cpu 1 -failfast -coverprofile=coverage.out -covermode=set ./...
````