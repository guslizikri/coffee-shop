package repository

import (
	"coffee-shop/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RepoFavorite struct {
	*sqlx.DB
}

func NewFavorite(db *sqlx.DB) *RepoFavorite {
	return &RepoFavorite{db}
}

func (r *RepoFavorite) ReadFavorite(data *models.Favorite, page int, limit int, category string, search string) ([]models.Favorite, error) {
	var filterQuery string
	var metaQuery string
	offset := (page - 1) * limit

	if page != 0 && limit != 0 {
		metaQuery += fmt.Sprintf(`limit %d offset %d`, limit, offset)
	}
	if search != "" {
		filterQuery += fmt.Sprintf(`AND LOWER(p.name_product) like LOWER('%s')`, "%"+search+"%")
	}

	if category != "" {
		filterQuery += fmt.Sprintf(`AND LOWER(p.category) like LOWER('%s')`, "%"+category+"%")
	}

	query := fmt.Sprintf(`SELECT f.id as id_product p.name_product, p.description, p.image, p.category, u.email from public.products p 
	inner JOIN public.favorites f ON f.product_id = p.id
	inner JOIN public.users u ON f.user_id = u.id WHERE true %s %s`, filterQuery, metaQuery)

	Favorites := []models.Favorite{}
	err := r.Select(&Favorites, query)
	if err != nil {
		return nil, err
	}
	return Favorites, nil
}

func (r *RepoFavorite) CreateFavorite(data *models.Favorite) (string, error) {
	q := `insert into Favorites (product_id, user_id) 
	VALUES(
		:product_id,
		:user_id
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 data Favorite created", nil

}

func (r *RepoFavorite) DeleteFavorite(data *models.Favorite) (string, error) {
	query := `DELETE FROM public.favorites WHERE id=:id`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete Favorite data successful", nil
}
