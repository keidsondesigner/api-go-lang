package repository

import (
	"context"
	"fmt"

	"api-go-lang/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ProductRepository é responsável por acessar os dados de produtos no banco.
// Recebe um ponteiro para o pool de conexões do Neon DB.
type ProductRepository struct {
	connection *pgxpool.Pool
}

// NewProductRepository cria um novo repositório conectado ao pool do banco.
func NewProductRepository(connection *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{connection: connection}
}

// GetProducts executa a query no banco e retorna a lista de produtos.
func (p *ProductRepository) GetProducts() ([]model.Product, error) {
	query := `SELECT id, name, description, price FROM product`
	rows, err := p.connection.Query(context.Background(), query)

	// Se deu erro ao executar a query, retornar nil e o erro
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		// Se deu erro ao capturar os dados, retornar nil e o erro
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			fmt.Println(err)
			return nil, err
		}
		// Se não deu erro, adicionar o produto ao slice
		products = append(products, product)
	}

	// Se não deu erro, retornar todos os produtos
	return products, nil
}

// GetProductById executa a query no banco e retorna o produto pelo ID.
func (p *ProductRepository) GetProductById(id int) (model.Product, error) {
	query := `SELECT id, name, description, price FROM product WHERE id = $1`
	rows, err := p.connection.Query(context.Background(), query, id)

	// Se deu erro ao executar a query, retornar nil e o erro
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}
	defer rows.Close()

	var product model.Product
	for rows.Next() {
		// Se deu erro ao capturar os dados, retornar nil e o erro
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			fmt.Println(err)
			return model.Product{}, err
		}
	}

	// Se não deu erro, retornar o produto
	return product, nil
}

func (p *ProductRepository) CreateProduct(product *model.Product) (int, error) {
	query := `INSERT INTO product (name, description, price) VALUES ($1, $2, $3) RETURNING id`

	var id int
	err := p.connection.QueryRow(context.Background(), query, product.Name, product.Description, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Se não deu erro, retornar o id
	return id, nil
}