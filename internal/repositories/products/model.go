package products

type Product struct {
	Code        string
	Name        string
	Description string
	Price       float64
}

type ProductUpdate struct {
	Code        *string
	Name        *string
	Description *string
	Price       *float64
}
