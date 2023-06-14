package storage

import (
	"app/models"
)

type StorageI interface {
	User() UserRepoI
	Product() ProductRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (*models.User, error)
	GetById(*models.UserPrimaryKey) (*models.User, error)
	GetList(*models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(*models.UpdateUser) (*models.User, error)
	Delete(*models.UserPrimaryKey) error
}

type ProductRepoI interface {
	Create(req *models.CreateProduct) (*models.Product, error)
	GetById(req *models.ProductPrimaryKey) (*models.Product, error)
	GetList(req *models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(req *models.UpdateProduct) (*models.Product, error)
	Delete(req *models.ProductPrimaryKey) error
}
