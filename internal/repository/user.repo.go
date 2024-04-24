package repository

import (
	"coffee-shop/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) ReadUser(data *models.User, page int, limit int, category string, search string) ([]models.User, error) {
	var filterQuery string
	var metaQuery string
	offset := (page - 1) * limit

	if page != 0 && limit != 0 {
		metaQuery += fmt.Sprintf(`limit %d offset %d`, limit, offset)
	}
	if search != "" {
		filterQuery += fmt.Sprintf(`AND LOWER(first_name) like LOWER('%s')`, "%"+search+"%")
	}

	if category != "" {
		filterQuery += fmt.Sprintf(`AND LOWER(display_name) like LOWER('%s')`, "%"+category+"%")
	}
	query := fmt.Sprintf("SELECT id, displayname, first_name, last_name, gender, phone, email, birth_date, role, image, created_at, updated_at from users WHERE true %s %s", filterQuery, metaQuery)

	users := []models.User{}
	err := r.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *RepoUser) CreateUser(data *models.User) (string, error) {
	q := `insert into users (email, password, phone) 
	VALUES(
		:email,
		:password,
		:phone
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 data User created", nil

}

func (r *RepoUser) UpdateUser(data *models.User) (string, error) {
	query := `UPDATE public.users SET
			displayname = COALESCE(NULLIF(:displayname, ''), displayname), 
			first_name = COALESCE(NULLIF(:first_name, ''), first_name), 
			last_name = COALESCE(NULLIF(:last_name, ''), last_name),
			updated_at = now()
			WHERE id = :id`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "update user data successful", nil
}
func (r *RepoUser) DeleteUser(data *models.User) (string, error) {
	query := `DELETE FROM public.users WHERE id=:id`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete user data successful", nil
}
