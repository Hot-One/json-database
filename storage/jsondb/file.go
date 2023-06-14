package jsondb

import (
	"os"

	"app/config"
	"app/storage"
)

type StoreJSON struct {
	user *UserRepo
}

func NewConnectionJSON(cfg *config.Config) (storage.StorageI, error) {

	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	return &StoreJSON{
		user: NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
	}, nil
}

func (u *StoreJSON) User() storage.UserRepoI {

	return u.user
}
