package controller

import (
	"errors"
	"log"

	"app/models"
)

func (c *Controller) UserCreate(req *models.CreateUser) (*models.User, error) {

	log.Printf("User create req: %+v\n", req)

	resp, err := c.Strg.User().Create(req)
	if err != nil {
		log.Printf("error while user create: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) GetById(id string) (*models.User, error){
	return c.Strg.User().GetById(id)
}

func (c *Controller) UserGetList(req *models.UserGetListRequest) (*models.UserGetListResponse){
	data := c.Strg.User().GetList()
	var (
		response = &models.UserGetListResponse{}
		offset   = c.Cfg.DefaultOffset
		limit    = c.Cfg.DefaultLimit
	)

	if req.Offset > 0 {
		offset = req.Offset
	}

	if req.Limit > 0 {
		limit = req.Limit
	}

	response.Count = len(data)
	if len(data) < limit+offset {
		response.Users = data[offset:]
	} else {
		response.Users = data[offset : offset+limit]
	}

	return response
}

func (c *Controller) UserUpdate(req *models.User) (*models.User, error){
	return c.Strg.User().Update(req)
}

func (c *Controller) UserDelete(id string) (*models.User, error){
	return c.Strg.User().Delete(id)
}