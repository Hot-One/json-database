package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type ProductCreate struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	CategoryID   string `json:"category_id"`
}

type Product struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	CategoryID   string `json:"category_id"`
}

type ProductUpdate struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	CategoryID   string `json:"category_id"`
}

type ProductGetListRequest struct {
	Offset int
	Limit  int
}

type ProductGetListResponse struct {
	Count    int
	Products []*Product
}
