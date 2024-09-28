# clean-architecture-challenge

## 1. Instanciando o mysql/rabbitmq

- Execute o comando `docker compose up --detach` para iniciar os servidores `mysql` e `rabbitmq`.
- Verifique o status do servidor com o comando `docker ps`
- Quando o servidor for instanciado, execute o comando `make migrate` para criar a tabela orders.

## 2. Iniciando os servidores http, gRPC e GraphQL

- Mude o diretório de trabalho para `cmd/ordersystem`
- Inicie os servidores `http`, `gRPC` e `graphQL` com o comando `go run main.go wire_gen.go`.

## 3. Adicionando itens na tabela

- Adicione itens na tabela orders utilizando o arquivo [api/create_order.http](api/create_order.http)

#### Exemplos:

```
POST http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
      "id": "aaaa",
      "price": 100.5,
      "tax": 0.5
}
```

```
POST http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
      "id": "bbbb",
      "price": 100.5,
      "tax": 0.5
}
```

## 4. HTTP - Listagem de orders

- Utilize o arquivo [api/list_orders.http](api/list_orders.http) para listar as ordens com o servidor http, a resposta deve ser uma lista, com todos os itens adicionados.

A resposta deve ser um `200 OK` como no exemplo a seguir:

```
HTTP/1.1 200 OK
Date: Sat, 28 Sep 2024 22:33:50 GMT
Content-Length: 125
Content-Type: text/plain; charset=utf-8
Connection: close

{
  "orders": [
    {
      "id": "aaaa",
      "price": 100.5,
      "tax": 0.5,
      "final_price": 101
    },
    {
      "id": "bbbb",
      "price": 100.5,
      "tax": 0.5,
      "final_price": 101
    }
  ]
}
```

## 5. gRPC - Listagem de orders

Em outro terminal, utilize o evans para listar as orders no gRPC, com os comandos a seguir:

```
evans -r repl

call ListOrders
```

A resposta do comando `call ListOrders` deve ser uma lista, como no exemplo a seguir:

```
{
  "orders": [
    {
      "finalPrice": 101,
      "id": "aaaa",
      "price": 100.5,
      "tax": 0.5
    },
    {
      "finalPrice": 101,
      "id": "bbbb",
      "price": 100.5,
      "tax": 0.5
    }
  ]
}
```

## 6. GraphQL - Listagem de orders

- No seu navegador de preferência, acesse a URL http://localhost:8080 para abrir o playgroud do GraphQL.

- Execute a seguinte `query` no terminal do GraphQL:

```
query Order {
  orders {
    id,
    Price,
    Tax,
    FinalPrice
  }
}
```

A resposta dessa `query` deve ser uma lista de orders, como no exemplo a seguir:

```
{
  "orders": [
    {
      "finalPrice": 101,
      "id": "aaaa",
      "price": 100.5,
      "tax": 0.5
    },
    {
      "finalPrice": 101,
      "id": "bbbb",
      "price": 100.5,
      "tax": 0.5
    }
  ]
}
```
