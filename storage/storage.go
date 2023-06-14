package storage

import "app/models"

type StorageI interface {
	User() UserRepoI
}

type UserRepoI interface {
	Create(req *models.CreateUser) (*models.User, error)
	GetById(id string) (*models.User, error)
	GetList() []*models.User
	Update(req *models.User,) (*models.User, error)
	Delete(id string)(*models.User, error)
}
