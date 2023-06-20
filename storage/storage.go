package storage

import (
	"app/models"
)

type StorageI interface {
	User() UserRepoI
	Product() ProductRepoI
	Order() OrderRepoI
	Category() CategoryRepoI
}

type UserRepoI interface {
	Create(*models.UserCreate) (*models.User, error)
	GetById(*models.UserPrimaryKey) (*models.User, error)
	GetList(*models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(*models.UserUpdate) (*models.User, error)
	Delete(*models.UserPrimaryKey) error
}

type CategoryRepoI interface {
	Create(*models.CategoryCreate) (*models.Category, error)
	GetById(*models.CategoryPrimaryKey) (*models.Category, error)
	GetList(*models.CategoryGetListRequest) (*models.CategoryGetListResponse, error)
	Update(*models.CategoryUpdate) (*models.Category, error)
	Delete(*models.CategoryPrimaryKey) error
}

type ProductRepoI interface {
	Create(*models.ProductCreate) (*models.Product, error)
	GetById(*models.ProductPrimaryKey) (*models.Product, error)
	GetList(*models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(*models.ProductUpdate) (*models.Product, error)
	Delete(*models.ProductPrimaryKey) error
}

type OrderRepoI interface {
	Create(*models.OrderCreate) (*models.Order, error)
	GetById(*models.OrderPrimaryKey) (*models.Order, error)
	GetList(*models.OrderGetListRequest) (*models.OrderGetListResponse, error)
	Update(*models.OrderUpdate) (*models.Order, error)
	Delete(*models.OrderPrimaryKey) error
	AddOrderItem(*models.OrderItemCreate) error
	RemoveOrderItem(*models.OrderItemPrimaryKeyRemove) error
}
