package models

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}
type CategoryCreate struct {
	Name string `json:"name"`
}
type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type CategoryUpdate struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type CategoryGetListRequest struct {
	Offset int
	Limit  int
}
type CategoryGetListResponse struct {
	Count     int
	Categorys []*Category
}
