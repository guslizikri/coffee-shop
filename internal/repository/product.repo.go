package repository

import (
	"coffee-shop/config"
	"coffee-shop/internal/models"
	"fmt"
	"math"

	"github.com/jmoiron/sqlx"
)

type RepoProductIF interface {
	CreateProduct(data *models.Product) (*config.Result, error)
	UpdateProduct(data *models.Product) (*config.Result, error)
	DeleteProduct(data *models.Product) (*config.Result, error)
	ReadProduct(params models.Query) (*config.Result, error)
}
type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r *RepoProduct) ReadProduct(params models.Query) (*config.Result, error) {
	var data models.Products
	var metas config.Metas
	var count int
	var filterQuery string
	var metaQuery string

	// untuk sql injection '?' rebind, dipakai setelah r.rebind
	// args gabungan meta dan filter ijection
	var args []interface{}
	// filter injection, digunakan saat query count total data, karena tidak butuh meta jadi dibikin var baru
	var filter []interface{}

	if params.Name != "%%" {
		filterQuery += "AND LOWER(name_product) like LOWER(?)"
		filter = append(filter, params.Name)
		args = append(args, params.Name)
	}

	if params.Category != "%%" {
		filterQuery += `AND LOWER(category) like LOWER(?)`
		filter = append(filter, params.Category)
		args = append(args, params.Category)
	}
	offset := (params.Page - 1) * params.Limit
	metaQuery = "LIMIT ? OFFSET ? "
	args = append(args, params.Limit, offset)
	// menghitung jumlah total data yg ada, untuk dimasukkan ke metas
	m := fmt.Sprintf(`SELECT COUNT(id) as count FROM products WHERE true %s`, filterQuery)
	err := r.Get(&count, r.Rebind(m), filter...)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * from products WHERE true %s %s", filterQuery, metaQuery)

	err = r.Select(&data, r.Rebind(query), args...)
	if err != nil {
		return nil, err
	}
	check := math.Ceil(float64(count) / float64(params.Limit))
	metas.Total = count
	if count > 0 && params.Page != int(check) {
		metas.Next = params.Page + 1
	}

	if params.Page != 1 {
		metas.Prev = params.Page - 1
	}

	return &config.Result{Data: data, Meta: metas}, nil
}

func (r *RepoProduct) CreateProduct(data *models.Product) (*config.Result, error) {
	q := `insert into products (name_product, description, image, category) 
	VALUES(
		:name_product,
		:description,
		:image, 
		:category
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return nil, err
	}

	return &config.Result{Message: "1 data product created"}, nil

}

func (r *RepoProduct) UpdateProduct(data *models.Product) (*config.Result, error) {
	query := `UPDATE public.products SET
			name_product = COALESCE(NULLIF(:name_product, ''), name_product), 
			description = COALESCE(NULLIF(:description, ''), description), 
			image = COALESCE(NULLIF(:image, ''), image), 
			category = COALESCE(NULLIF(:category, ''), category), 
			updated_at = now()
			WHERE id = :id`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 data product updated"}, nil

}
func (r *RepoProduct) DeleteProduct(data *models.Product) (*config.Result, error) {
	query := `DELETE FROM public.products WHERE id=:id`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 data product deleted"}, nil

}
