package models

import "time"

type Favorite struct {
	Id           string     `db:"id" form:"id" json:"id" uri:"id"`
	Product_id   string     `db:"product_id" form:"product_id" json:"product_id"`
	User_id      string     `db:"user_id" form:"user_id" json:"user_id"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	Name_product string     `db:"name_product" form:"name_product" json:"name_product"`
	Description  *string    `db:"description" form:"description" json:"description"`
	Image        *string    `db:"image" form:"image" json:"image"`
	Category     string     `db:"category" form:"category" json:"category"`
	Email        *string    `db:"email" form:"email" json:"email"`
}
