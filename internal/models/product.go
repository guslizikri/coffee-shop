package models

import "time"

type Product struct {
	Id           string     `db:"id" form:"id" json:"id" uri:"id"`
	Name_product string     `db:"name_product" form:"name_product" json:"name_product"`
	Description  *string    `db:"description" form:"description" json:"description"`
	Image        *string    `db:"image" form:"image" json:"image"`
	Category     string     `db:"category" form:"category" json:"category"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
}
