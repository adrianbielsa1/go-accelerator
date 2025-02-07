package mysql

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Code        string
	Name        string
	Description string
	Price       float64
}

type ProductUpdate struct {
	gorm.Model

	Code        *string
	Name        *string
	Description *string
	Price       *float64
}
