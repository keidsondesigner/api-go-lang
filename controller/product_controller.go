package controller

import (
	"net/http"

	"api-go-lang/usecase"

	"github.com/gin-gonic/gin"
)

// productController recebe um ponteiro para ProductUsecase.
// Ponteiro (*) garante que todos compartilham o mesmo usecase,
// sem criar cópias a cada requisição.
type productController struct {
	// Usecase
	productUsecase *usecase.ProductUsecase
}

// NewProductController recebe e retorna ponteiros (*),
// conectando o controller ao usecase sem duplicar objetos na memória.
func NewProductController(productUsecase *usecase.ProductUsecase) *productController {
	return &productController{
		productUsecase: productUsecase,
	}
}

// GetProducts é o handler HTTP GET /product.
// Chama o usecase e retorna os produtos em JSON,
// ou 500 se ocorrer algum erro.
func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
