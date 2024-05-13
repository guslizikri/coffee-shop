package repository

import (
	"coffee-shop/config"
	"coffee-shop/internal/models"
	"fmt"
	"math"

	"github.com/jmoiron/sqlx"
)

type RepoFavoriteIF interface {
	CreateFavorite(data *models.Favorite) (*config.Result, error)
	DeleteFavorite(data *models.Favorite) (*config.Result, error)
	ReadFavorite(params models.Query) (*config.Result, error)
}
type RepoFavorite struct {
	*sqlx.DB
}

func NewFavorite(db *sqlx.DB) *RepoFavorite {
	return &RepoFavorite{db}
}

func (r *RepoFavorite) ReadFavorite(params models.Query) (*config.Result, error) {
	var data models.Favorites
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
		filterQuery += `AND LOWER(name_product) like LOWER(?)`
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
	m := fmt.Sprintf(`SELECT  COUNT(distinct p.id) as count FROM products p 
	inner JOIN public.favorites f ON f.product_id = p.id
	inner JOIN public.users u ON f.user_id = u.id WHERE true %s`, filterQuery)
	err := r.Get(&count, r.Rebind(m), filter...)
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`SELECT f.id, p.name_product, p.description, p.image, p.category, u.email from public.products p 
	inner JOIN public.favorites f ON f.product_id = p.id
	inner JOIN public.users u ON f.user_id = u.id WHERE true %s %s`, filterQuery, metaQuery)

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

func (r *RepoFavorite) CreateFavorite(data *models.Favorite) (*config.Result, error) {
	q := `insert into Favorites (product_id, user_id) 
	VALUES(
		:product_id,
		:user_id
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 data favorite created"}, nil

}

func (r *RepoFavorite) DeleteFavorite(data *models.Favorite) (*config.Result, error) {
	query := `DELETE FROM public.favorites WHERE id=:id`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 data favorite deleted"}, nil

}
