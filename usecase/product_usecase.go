package usecase

import (
	"api-go-lang/model"
	"api-go-lang/repository"
)

// ProductUsecase contém as regras de negócio relacionadas a produtos.
// O * indica que os métodos operam sobre o ponteiro (endereço original),
// evitando cópias desnecessárias do struct.
type ProductUsecase struct {
	//Repository
	productRepository *repository.ProductRepository
}

// NewProductUsecase retorna um ponteiro (*) para ProductUsecase.
// Usando & criamos o struct e já retornamos seu endereço de memória,
// assim todos que receberem esse valor apontam para o mesmo objeto.
func NewProductUsecase(productRepository *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepository: productRepository}
}

// GetProducts busca a lista de produtos.
// O receiver (p *ProductUsecase) usa ponteiro para trabalhar no objeto original,
// não em uma cópia — necessário pois o tipo foi declarado com receivers de ponteiro.
func (p *ProductUsecase) GetProducts() ([]model.Product, error) {
	return p.productRepository.GetProducts()
}
