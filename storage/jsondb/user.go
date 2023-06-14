package jsondb

import (
	"log"
	"errors"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/google/uuid"

	"app/models"
)

type UserRepo struct {
	fileName string
	file     *os.File
}

func NewUserRepo(fileName string, file *os.File) *UserRepo {
	return &UserRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *UserRepo) Create(req *models.CreateUser) (*models.User, error) {

	var users []*models.User
	err := json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return nil, err
	}

	id := uuid.New()
	var user = &models.User{
		Id:        id.String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	users = append(users, user)

	body, err := json.MarshalIndent(users, "", "	")
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) GetById(id string) (*models.User, error){
	var users []*models.User
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil{
		log.Printf("Error while Read data: %+v", err)
		return nil, nil
	}
	err1 := json.Unmarshal(data, &users)
	if err1 != nil{
		log.Printf("Error while Unmarshal data: %+v", err1)
		return nil, nil
	}
	for _, user := range users{
		if user.Id == id{
			return user, nil
		}
	}
	
	return nil, errors.New("In database doest exist user with this id")
}

func (u *UserRepo) GetList() []*models.User{
	var (
		users []*models.User
	)
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil{
		log.Printf("Error while Read data: %+v", err)
		return nil
	}
	err1 := json.Unmarshal(data, &users)
	if err1 != nil{
		log.Printf("Error while Unmarshal data: %+v", err1)
		return nil
	}
	return users
}

func (u *UserRepo) Update(req *models.User) (*models.User, error){
	var users []*models.User
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil{
		log.Printf("Error while Read data: %+v", err)
		return nil, nil
	}
	err1 := json.Unmarshal(data, &users)
	if err1 != nil{
		log.Printf("Error while Unmarshal data: %+v", err1)
		return nil, nil
	}
	var updatedUser *models.User
	for _, user := range users{
		if user.Id == req.Id{
			user.FirstName = req.FirstName
			user.LastName = req.LastName
			updatedUser = user
		}
	}
	body, err := json.MarshalIndent(users, "", "	")

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *UserRepo) Delete(id string) (*models.User, error){
	var users []*models.User
	var newusers []*models.User
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil{
		log.Printf("Error while Read data: %+v", err)
		return nil, nil
	}
	err1 := json.Unmarshal(data, &users)
	if err1 != nil{
		log.Printf("Error while Unmarshal data: %+v", err1)
		return nil, nil
	}
	for _, user := range users{
		if user.Id != id{
			newusers = append(newusers, user)
		}
	}

	body, err := json.MarshalIndent(newusers, "", "	")
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

