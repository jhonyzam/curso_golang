# 8- Hora do famoso crud

Finalmente temos todos os conceitos, ou pelo menos quase todos, e agora chegou a hora
de consolidarmos isso. Vamos desenvolver um servidor HTTP que provê uma API, para 
criar, ler, editar e deletar produtos.

### 8.1- Estrutura do projeto

Até agora, todos os exemplos definiam um único pacote `main`. Projetos maiores e mais
complexos, porém, precisam de pacotes adicionais para organizar melhor o código.

Todo projeto em Go segue uma convenção semelhante. A estrutura de diretórios e pacotes
é definida de acordo com o caminho do repositório onde o código reside. Por exemplo projetos
hospedados no GitHub incluem `github.com/usuario/projeto/nome-do-pacote` no caminho
completo dos pacotes.

```bash
├── products
│   ├── domain/
│   │   ├── product/
│   │   ├── **/*.go
│   ├── internal/
│   ├── pkg/
│   ├── test/
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── main.go
└── README.md
```

A estrutura acima segue os padrões idiomáticos do Go e utiliza as práticas de 
Domain Driven Design (DDD) a medida que faça sentido para o nosso contexto.
Assim temos uma estrutura de pastas simples. Vale ressaltar que não existe um certo ou 
errado, existem outras estruturas e cada uma pode se adequar melhor para cada pessoa
ou time.

* `domain/` Esta pasta é onde vamos colocar a lógica de negócio. É nela onde adicionamos
no nosso caso as regras para manipular o produto.
  

* `internal/`
  * `infrastruct/`
      * `server/`: Destinado a estrutura do servidor, é onde vamos montar nossos endpoints.
      * `storage/`: Destinado ao banco de dados, caso houvesse mais de um database, adicionaríamos outra pasta além da `postgres`.

* `pkg/` Usado para manter pacotes reutilizáveis, no nosso caso apenas o `env` e `http_client`.

### 8.2- Tour pelo código

Como sempre, vamos iniciar pelo `main.go`. O código possui a mesma característica do
que foi apresentado no curso, com a diferença de ser pacotes diferentes.

~~~go
log.Printf("Products API starting - version: %s; date: %s\n", version, date)

env.CheckRequired(envPostgresURL)
~~~ 
A função main inicia imprimindo um log de inicio da API, usamos o pacote `log` padrão 
do Go. Na sequencia testamos as variáveis de ambiente obrigatórias, no nosso caso a
`envPostgresURL`. Isso serve para garantir que a api não irá subir em caso do dado não estar
preenchido.
~~~go
// Storages...
db, err := postgres.NewConnection(getPostgresURL())
if err != nil {
    log.Fatalf("ERROR: connecting database: %q\n", err)
    return
}
~~~
Na sequência criamos a conexão com o banco de dados, retornando a variável `db` para ser
usada em outro serviço, e `err` que é validada logo após execução, garantindo que 
não há nenhum erro.

Com a conexão pronta chamamos nosso primeiro serviço, o `productStorage`. Este é um serviço
que consideramos como interno, executando as _queries_ do CRUD.
~~~go
productStorage := postgres.NewProductStorage(db)
~~~
Quando executamos esse comando, estamos criando um serviço do tipo `domain/ProductStorage`, que 
como podemos ver é uma interface das `func` a serem usadas.
~~~go
type ProductStorage interface {
    Insert(product *Product) (*Product, error)
    FindByID(productID int) (*Product, error)
    Update(product *Product) error
    Delete(productID int) error
}
~~~
Estas funções estão implementadas dentro de `internal/infrastructure/storafe/postgres/product.go`
Todas utilizam do pacote `*sql.DB`, ou seja da variável `db` passada na criação do 
`NewProductStorage(db)`.

Seguindo com o nosso main, o próximo serviço criado é o:
~~~go
productService := product.NewService(productStorage)
~~~
Este serviço é onde vamos implementar as regras de negócio do nosso código. Quando 
olhamos o código dentro de `domain/product`, podemos ver que a maioria das funções apenas 
chama o serviço de storage correspondente. Porém, vamos usar como exemplo a função `GET`:
~~~go
func (s *service) Get(productID int) (*domain.Product, error) {
    if productID == 0 {
        return nil, domain.ErrInvalidProductID
    }
    return s.productStorage.FindByID(productID)
}
~~~
Como mencionado, neste serviço colocamos as regras de negócio. Nesse exemplo, podemos ver
que tem uma validação: "se o `productID` for igual a `0` ele vai retornar um erro". Sendo assim
podemos colocar qualquer outra regra neste código que seja útil para o produto final.

O próximo passo do main é a criação do serviço de `handler`. Este serviço é responsável
pela criação dos nossos endpoints:
~~~go
handler := http.NewHandler(productService)
~~~
Função implementada:
~~~go
func NewHandler(productService domain.ProductService) http.Handler {
    handler := &handler{
        productService: productService,
    }
  
    gin.SetMode(gin.ReleaseMode)
  
    router := gin.New()
    v1 := router.Group("/v1")
  
    v1.POST("/products", handler.postProduct)
    v1.GET("/products/:id", handler.getProduct)
    v1.PUT("/products/:id", handler.updateProduct)
    v1.DELETE("/products/:id", handler.deleteProduct)
  
    return router
}
~~~
Note que estamos usando o pacote do `gin` para acessar o serviço `http`, este é apenas
um facilitador, pode ser usado o pacote `http` puro, ou qualquer outra lib a sua escolha.
Nosso serviço terá quatro endpoints, que vão estar implementados dentro de 
`internal/infrastructure/server/http/product.go`.

As funções do handler, são como gatilhos, responsáveis por usar os serviços que obtém o 
resultado final:
~~~go
func (h *handler) getProduct(c *gin.Context) {
    productID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return
    }
    product, err := h.productService.Get(productID)
    if err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }
  
    c.JSON(http.StatusOK, product)
}
~~~
A função `getProduct()` por exemplo, recebe como parâmetro do `gin` um `id` de produto, 
converte para `int`, chama o método `h.productService.Get(productID)` que vai buscar esse produto
na base de dados. Se tudo der certo ele retorna usando o comando ` c.JSON(http.StatusOK, product)`.
Esse comando devolve o resultado no formato para o `http`, com status OK ou 200, e no
response os dados do produto.

Por fim, após implementar todos os serviços, executamos o server:
~~~go
server := http.New(getApplicationPort(), handler)
server.ListenAndServe()
~~~
Criamos um `http.Server`, e o executamos, semelhante ao que ja vimos, com a diferença
de estar separado em métodos.