package jsondb

import (
	"os"

	"app/config"
	"app/storage"
)

type StoreJSON struct {
	user *UserRepo
	product *ProductRepo
}

func NewConnectionJSON(cfg *config.Config) (storage.StorageI, error) {

	userFile, err := os.Open(cfg.UserPath + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	productFile, err := os.Open(cfg.ProductPath + cfg.ProductFileName)
	if err != nil {
		return nil, err
	}

	return &StoreJSON{
		user: NewUserRepo(cfg.UserPath+cfg.UserFileName, userFile),
		product: NewProductRepo(cfg.ProductPath+cfg.ProductFileName, productFile),
	}, nil
}

func (u *StoreJSON) User() storage.UserRepoI {
	return u.user
}

func (u *StoreJSON) Product() storage.ProductRepoI {
	return u.product
}
