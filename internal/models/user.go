package models

import "time"

var schema = `
CREATE TABLE public.users (
	id uuid primary key NOT NULL DEFAULT uuid_generate_v4(),
  	displayname varchar(100) NULL,
	first_name varchar(100) NULL ,
	last_name varchar(100) NULL,
    gender varchar(15) NULL,
	phone varchar(15) not NULL,
	email text NOT NULL,
	password text NOT NULL,
  	birth_date timestamp NULL,
	role varchar(6) NOT NULL DEFAULT 'user'::character varying,
	image text NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL
);
`

type User struct {
	Id          string     `db:"id" form:"id" json:"id" uri:"id"`
	Displayname *string    `db:"displayname" form:"displayname" json:"displayname"`
	First_name  *string    `db:"first_name" form:"first_name" json:"first_name"`
	Last_name   *string    `db:"last_name" form:"last_name" json:"last_name"`
	Gender      *string    `db:"gender" form:"gender" json:"gender"`
	Phone       string     `db:"phone" form:"phone" json:"phone"`
	Email       string     `db:"email" form:"email" json:"email"`
	Password    string     `db:"password" form:"password" json:"password"`
	Birth_date  *time.Time `db:"birth_date" form:"birth_date" json:"birth_date"`
	Role        string     `db:"role" form:"role" json:"role"`
	Image       *string    `db:"image" form:"image" json:"image"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}