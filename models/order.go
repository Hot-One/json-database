package models

type OrderPrimaryKey struct {
	Id string `json:"id"`
}

type Order struct {
	Id         string             `json:"id"`
	UserId     string             `json:"user_id"`
	Sum        int                `json:"sum"`
	SumCount   int                `json:"sum_count"`
	DateTime   string             `json:"date_time"`
	Status     string             `json:"status"`
	OrderItems []*OrderItemCreate `json:"order_items"`
}

type OrderGetListResponse struct {
	Count  int
	Orders []*Order
}
type OrderCreate struct {
	UserId   string `json:"user_id"`
	Sum      int    `json:"sum"`
	SumCount int    `json:"sum_count"`
	DateTime string `json:"date_time"`
	Status   string `json:"status"`
}

type OrderGetListRequest struct {
	Offset int
	Limit  int
}

type OrderUpdate struct {
	Id         string             `json:"id"`
	UserId     string             `json:"user_id"`
	Sum        int                `json:"sum"`
	SumCount   int                `json:"sum_count"`
	DateTime   string             `json:"date_time"`
	Status     string             `json:"status"`
	OrderItems []*OrderItemCreate `json:"order_items"`
}

type OrderItemCreate struct {
	Id         string `json:"id"`
	ProductId  string `json:"product_id"`
	OrderId    string `json:"order_id"`
	Count      int    `json:"count"`
	TotalPrice int    `json:"total_price"`
}

type OrderItemPrimaryKeyRemove struct {
	Id      string `json:"id"`
	OrderId string `json:"order_id"`
}

type OrderPayment struct {
	OrderId string `json:"order_id"`
}

type OrderItem struct {
	Id         string `json:"id"`
	ProductId  string `json:"product_id"`
	OrderId    string `json:"order_id"`
	Count      int    `json:"count"`
	TotalPrice int    `json:"total_price"`
}
