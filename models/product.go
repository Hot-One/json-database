package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateProduct struct {
	Name string `json:"first_name"`
	Price  string `json:"last_name"`
}

type Product struct {
	Id        string `json:"id"`
	Name string `json:"first_name"`
	Price  string `json:"last_name"`
}

type UpdateProduct struct {
	Id        string `json:"id"`
	Name string `json:"first_name"`
	Price  string `json:"last_name"`
}

type ProductGetListRequest struct {
	Offset int
	Limit  int
}

type ProductGetListResponse struct {
	Count int
	Products []*Product
}
