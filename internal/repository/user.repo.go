package repository

import (
	"coffee-shop/config"
	"coffee-shop/internal/models"
	"errors"
	"fmt"
	"math"

	"github.com/jmoiron/sqlx"
)

type RepoUserIF interface {
	CreateUser(data *models.User) (*config.Result, error)
	UpdateUser(data *models.User) (*config.Result, error)
	DeleteUser(data *models.User) (*config.Result, error)
	ReadUser(params models.Query) (*config.Result, error)
	GetAuthData(user string) (*models.User, error)
}
type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

// hanya untuk kebutuhan dev, seperti login, register dll
func (r *RepoUser) GetAuthData(user string) (*models.User, error) {
	var result models.User
	q := `SELECT id, email, role, password, phone FROM users WHERE email = ?`

	if err := r.Get(&result, r.Rebind(q), user); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("username not found")
		}

		return nil, err
	}

	return &result, nil
}

func (r *RepoUser) ReadUser(params models.Query) (*config.Result, error) {
	var data models.Users
	var metas config.Metas
	var count int
	var filterQuery string
	var metaQuery string

	// untuk sql injection '?' rebind, dipakai setelah r.rebind
	// args gabungan meta dan filter ijection
	var args []interface{}
	// filter injection, digunakan saat query count total data, karena tidak butuh meta jadi dibikin var baru
	var filter []interface{}

	if params.Name != "" {
		filterQuery += `AND LOWER(displayname) like LOWER(?)`
		filter = append(filter, params.Name)
		args = append(args, params.Name)
	}
	offset := (params.Page - 1) * params.Limit
	metaQuery = "LIMIT ? OFFSET ? "
	args = append(args, params.Limit, offset)

	// menghitung jumlah total data yg ada, untuk dimasukkan ke metas
	m := fmt.Sprintf(`SELECT COUNT(id) as count FROM users WHERE true %s`, filterQuery)
	err := r.Get(&count, r.Rebind(m), filter...)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT id, displayname, first_name, last_name, gender, phone, email, birth_date, role, image, created_at, updated_at from users WHERE true %s %s", filterQuery, metaQuery)

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

func (r *RepoUser) CreateUser(data *models.User) (*config.Result, error) {
	q := `insert into users (email, password, phone, role) 
	VALUES(
		:email,
		:password,
		:phone,
		:role
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return nil, err
	}

	return &config.Result{Message: "1 data user created"}, nil

}

func (r *RepoUser) UpdateUser(data *models.User) (*config.Result, error) {
	query := `UPDATE public.users SET
			displayname = COALESCE(NULLIF(:displayname, ''), displayname), 
			first_name = COALESCE(NULLIF(:first_name, ''), first_name), 
			last_name = COALESCE(NULLIF(:last_name, ''), last_name),
			image = COALESCE(NULLIF(:image, ''), image),
			updated_at = now()
			WHERE id = :id`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 data user updated"}, nil

}
func (r *RepoUser) DeleteUser(data *models.User) (*config.Result, error) {
	query := `DELETE FROM public.users WHERE id=:id`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 data user deleted"}, nil
}
