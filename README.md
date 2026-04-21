
# Como rodar o projeto?
go run ./cmd/main.go

## Qual a funcao de "nil"?

Em Go, nil é o valor zero para tipos que podem ser "vazios": ponteiros, slices, maps, funções, canais e interfaces.

No seu código aparecem dois usos:

1. Verificar erro — linha 29


if err != nil {
err é do tipo error (uma interface). Quando nenhum erro ocorre, seu valor é nil. Quando algo deu errado, ele contém um objeto com a mensagem de erro. Checar != nil é o jeito idiomático de dizer "se deu erro".

2. Retornar slice vazio em caso de erro — linha 41


return nil, err
Aqui nil está sendo retornado no lugar de []model.Product. Um slice nil em Go é funcional e equivalente a um slice vazio para a maioria dos usos — só significa "não há dados". É diferente do []model.Product{} da linha 31 (que aloca um slice vazio explícito), mas o efeito prático é o mesmo no contexto de retorno de erro.

Resumo prático:

Tipo	nil significa
error	sem erro
[]T (slice)	slice vazio / sem dados
*T (ponteiro)	ponteiro não inicializado
interface	sem implementação atribuída
Em outras linguagens seria equivalente a null, mas em Go nil é tipado — um nil de ponteiro é diferente de um nil de interface.


# Qual a funcao desse "*"?

● O * é usado para declarar e acessar ponteiros.
  ---                                                                                                     Declaração

  type ProductUsecase struct { /* campos */ }
  var p *ProductUsecase  // declara p como ponteiro para ProductUsecase

  ---                                                                                                     No seu código
  
  productController recebe um ponteiro para ProductUsecase.
  Ponteiro (*) garante que todos compartilham o mesmo usecase,
  sem criar cópias a cada requisição.

  type productController struct {
	  productUsecase *usecase.ProductUsecase
  }


  NewProductController recebe e retorna ponteiros (*),
  conectando o controller ao usecase sem duplicar objetos na memória.

  func NewProductController(productUsecase *usecase.ProductUsecase) *productController {
	  return &productController{
		  productUsecase: productUsecase,
	  }
  }

  GetProducts é o handler HTTP GET /product.
  Chama o usecase e retorna os produtos em JSON,
  ou 500 se ocorrer algum erro.

  func (p *productController) GetProducts(ctx *gin.Context) {
	  products, err := p.productUsecase.GetProducts()
	  if err != nil {
		 ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		 return
	  }
	  ctx.JSON(http.StatusOK, products)
}



# Qual a funcao desse "&"?

● O & retorna o endereço de memória de um valor, transformando-o em ponteiro.
  ---                                                                                                     Analogia

  Imagine que um struct é uma casa e o ponteiro é o endereço da casa:

  ProductUsecase{}   // a casa em si
  &ProductUsecase{}  // o endereço da casa (ex: "Rua A, nº 42")

  Quando você passa o endereço, qualquer um que receber sabe exatamente onde a casa está e pode mexer
  nela diretamente.

  ---
  No seu código
  
  func NewProductUsecase() *ProductUsecase {
      return &ProductUsecase{}  // cria o struct E retorna seu endereço
  }

  Sem o &:
  return ProductUsecase{}  // retorna uma CÓPIA do struct

  Com o &:                                                                                                return &ProductUsecase{}  // retorna o ENDEREÇO do struct original

  ---                                                                                                     Resumo da dupla & e *

  ┌─────────┬───────────────┬─────────────────────────────────────────────┐
  │ Símbolo │     Nome      │                  O que faz                  │
  ├─────────┼───────────────┼─────────────────────────────────────────────┤
  │ &       │ "endereço de" │ Pega um valor e retorna seu ponteiro        │
  ├─────────┼───────────────┼─────────────────────────────────────────────┤
  │ *       │ "valor em"    │ Pega um ponteiro e retorna o valor apontado │
  └─────────┴───────────────┴─────────────────────────────────────────────┘

  u := ProductUsecase{}   // valor
  p := &u                 // ponteiro (endereço de u)
  v := *p                 // volta ao valor (valor em p)
  
  Eles são opostos — & vai do valor para o ponteiro, * vai do ponteiro para o valor.

❯ // NewProductController recebe e retorna ponteiros (*),                                                 // conectando o controller ao usecase sem duplicar objetos na memória.                                  func NewProductController(productUsecase*usecase.ProductUsecase) *productController {                      return &productController{                                                                                  productUsecase: productUsecase,                                                                     }                                                                                                }
● Ótima pergunta — vamos destrinchar linha por linha:
  
  ---                                                                                                     func NewProductController(productUsecase *usecase.ProductUsecase)*productController {

- Recebe um ponteiro para ProductUsecase (o endereço do usecase criado no main.go)                      - Retorna um ponteiro para productController (o endereço do controller que será criado)
                                                                                                          ---                                                                                                     return &productController{
- productController{ } — cria o struct
- & — pega o endereço desse struct e retorna (por isso o retorno é *productController)

  ---
      productUsecase: productUsecase,                                                                     - Preenche o campo productUsecase do struct com o ponteiro recebido como parâmetro                    

- É como dizer: "guarda o endereço do usecase aqui dentro"
                                                                                                          ---                                                                                                     Fluxo no seu main.go
                                                                                                          ProductUsecase := usecase.NewProductUsecase()       // cria usecase, retorna &ProductUsecase
  ProductController := controller.NewProductController(ProductUsecase) // passa o endereço adiante
- Preenche o campo productUsecase do struct com o ponteiro recebido como parâmetro                      - É como dizer: "guarda o endereço do usecase aqui dentro"                                                                                                                                                      ---                                                                                                     Fluxo no seu main.go                                                                                                                                                                                            ProductUsecase := usecase.NewProductUsecase()       // cria usecase, retorna &ProductUsecase
  ProductController := controller.NewProductController(ProductUsecase) // passa o endereço adiante                                                                                                                memória:                                                                                                ┌─────────────────────┐                                                                                 │   ProductUsecase    │ ← endereço: 0xc0001                                                             └─────────────────────┘                                                                                         ↑                                                                                               ┌─────────────────────┐                                                                                 │  productController  │
  │  productUsecase ────┼──→ 0xc0001                                                                      └─────────────────────┘                                                                                                                                                                                         O controller não tem uma cópia do usecase — ele tem o endereço, então ambos apontam para o mesmo        objeto na memória.
                                                                                                        ────────────────────────────────────────────────────────────────────────────────────────────────────────❯  
────────────────────────────────────────────────────────────────────────────────────────────────────────  ? for shortcuts

  ---
      productUsecase: productUsecase,

- Preenche o campo productUsecase do struct com o ponteiro recebido como parâmetro
- É como dizer: "guarda o endereço do usecase aqui dentro"

  ---
  Fluxo no seu main.go

  ProductUsecase := usecase.NewProductUsecase()       // cria usecase, retorna &ProductUsecase
  ProductController := controller.NewProductController(ProductUsecase) // passa o endereço adiante

  memória:
  ┌─────────────────────┐
  │   ProductUsecase    │ ← endereço: 0xc0001
  └─────────────────────┘
          ↑
  ┌─────────────────────┐
  │  productController  │
  │  productUsecase ────┼──→ 0xc0001
  └─────────────────────┘

  O controller não tem uma cópia do usecase — ele tem o endereço, então ambos apontam para o mesmo
  objeto na memória.

───────────────────────────────────────────────────────────────────────────────────────────────────────❯  
───────────────────────────────────────────────────────────────────────────────────────────────────────  ? for shortcuts

  ---
  func NewProductController(productUsecase *usecase.ProductUsecase)*productController {

- Recebe um ponteiro para ProductUsecase (o endereço do usecase criado no main.go)
- Retorna um ponteiro para productController (o endereço do controller que será criado)

  ---
  return &productController{

- productController{ } — cria o struct
- & — pega o endereço desse struct e retorna (por isso o retorno é *productController)

  ---
      productUsecase: productUsecase,

- Preenche o campo productUsecase do struct com o ponteiro recebido como parâmetro
- É como dizer: "guarda o endereço do usecase aqui dentro"

  ---
  Fluxo no seu main.go

  ProductUsecase := usecase.NewProductUsecase()       // cria usecase, retorna &ProductUsecase
  ProductController := controller.NewProductController(ProductUsecase) // passa o endereço adiante

  memória:
  ┌─────────────────────┐
  │   ProductUsecase    │ ← endereço: 0xc0001
  └─────────────────────┘
          ↑
  ┌─────────────────────┐
  │  productController  │
  │  productUsecase ────┼──→ 0xc0001
  └─────────────────────┘

  O controller não tem uma cópia do usecase — ele tem o endereço, então ambos apontam para o mesmo objeto
  na memória.
