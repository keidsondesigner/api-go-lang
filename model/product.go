package model

// Product representa a estrutura de dados de um produto.
// As tags json definem o nome dos campos na resposta JSON da API.
type Product struct {
	ID          int     `json:"id_product"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
