
## Docker
`docker-compose up -d`

### Acessar o banco pelo terminal
`docker-compose exec mysql bash`
`mysql -uroot -p orders`

## Rodar a aplicação
No diretório cmd/ordersystem, rodar:
`go run main.go wire_gen.go`

- O servidor web rodará na porta: 8000
- O servidor grpc rodará na porta: 50051
- O servidor GraphQL rodará na porta: 8080

### web
Utilizar arquivo api/api.http

### grpc
Após iniciar a aplicação, utilizar os seguintes comandos no terminal para utilizar o grpc:
1. `evans -r repl`
2. `package pb`
3. `service OrderService`
4. `call ListOrders` ou `call CreateOrder`

### graphQL
http://localhost:8080/

``` mutation createOrder {
  createOrder(input: {
    id: "d",
 		Price	: 320,
    Tax: 30,
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}

query queryOrders {
  orders {
    id
    Price
    FinalPrice
    Tax
  }
} 