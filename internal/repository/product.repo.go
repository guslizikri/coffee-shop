package repository

import (
	"coffee-shop/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r *RepoProduct) ReadProduct(data *models.Product, page int, limit int, category string, search string) ([]models.Product, error) {

	var filterQuery string
	var metaQuery string
	offset := (page - 1) * limit

	if page != 0 && limit != 0 {
		metaQuery += fmt.Sprintf(`limit %d offset %d`, limit, offset)
	}
	if search != "" {
		filterQuery += fmt.Sprintf(`AND LOWER(name_product) like LOWER('%s')`, "%"+search+"%")
	}

	if category != "" {
		filterQuery += fmt.Sprintf(`AND LOWER(category) like LOWER('%s')`, "%"+category+"%")
	}

	query := fmt.Sprintf("SELECT * from products WHERE true %s %s", filterQuery, metaQuery)

	Products := []models.Product{}
	err := r.Select(&Products, query)
	if err != nil {
		return nil, err
	}
	return Products, nil
}

func (r *RepoProduct) CreateProduct(data *models.Product) (string, error) {
	q := `insert into products (name_product, description, image, category) 
	VALUES(
		:name_product,
		:description,
		:image, 
		:category
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 data product created", nil

}

func (r *RepoProduct) UpdateProduct(data *models.Product) (string, error) {
	query := `UPDATE public.products SET
			name_product = COALESCE(NULLIF(:name_product, ''), name_product), 
			description = COALESCE(NULLIF(:description, ''), description), 
			image = COALESCE(NULLIF(:image, ''), image), 
			category = COALESCE(NULLIF(:category, ''), category), 
			updated_at = now()
			WHERE id = :id`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "update product data successful", nil
}
func (r *RepoProduct) DeleteProduct(data *models.Product) (string, error) {
	query := `DELETE FROM public.products WHERE id=:id`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete Product data successful", nil
}
