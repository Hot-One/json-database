package controller

import (
	"errors"
	"fmt"
	"log"
	"sort"

	"app/models"
)

func (c *Controller) HistoryUser(req *models.OrderPrimaryKey) (*models.UserHistory, error) {

	log.Printf("Order id req: %+v", req)
	productsinfo := []int{}
	product := make(map[string][]int)

	// Order \\

	getorder, err := c.OrderGetById(req)
	if err != nil {
		log.Printf("Error while HistoryUser => GetByIdOrder: %+v", err)
		return nil, err
	}

	// User \\

	userId := getorder.UserId
	getUser, err := c.UserGetById(&models.UserPrimaryKey{
		Id: userId,
	})
	if err != nil {
		log.Printf("Error while HistoryUser => UserGetById: %+v", err)

		return nil, err
	}

	// Product \\

	getorderitems := getorder.OrderItems
	for _, items := range getorderitems {
		getProduct, err := c.PoductGetById(&models.ProductPrimaryKey{
			Id: items.ProductId,
		})
		productsinfo = append(productsinfo, items.Count, getProduct.Price)
		product[getProduct.Name] = productsinfo
		productsinfo = []int{}
		if err != nil {
			log.Printf("Error while HistoryUser => GetByIdPoduct: %+v", err)
			return nil, err
		}
	}
	return &models.UserHistory{
		FirstName: getUser.FirstName,
		LastName:  getUser.LastName,
		Products:  product,
	}, nil
}

func (c *Controller) ActiveUser() (string, error) {
	users := make(map[string]int)
	getorder, err := c.OrderGetList(&models.OrderGetListRequest{})
	if err != nil {
		return "", err
	}
	for _, value := range getorder.Orders {
		users[value.UserId] += value.Sum
	}
	user, sum := "", 0
	for key, value := range users {
		if sum < value {
			user = key
			sum = value
		}
	}
	getuser, err := c.UserGetById(&models.UserPrimaryKey{
		Id: user,
	})
	if err != nil {
		return "", err
	}
	return getuser.FirstName, nil
}

func (c *Controller) TopProducts() error {
	prodctsMap := make(map[string]int)
	getOrder, err := c.OrderGetList(&models.OrderGetListRequest{})
	if err != nil {
		return err
	}

	for _, value := range getOrder.Orders {
		for _, items := range value.OrderItems {
			prodctsMap[items.ProductId] += items.Count
		}
	}
	keys := make([]string, 0, len(prodctsMap))
	for key := range prodctsMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return prodctsMap[keys[i]] > prodctsMap[keys[j]]
	})
	i := 1
	for _, k := range keys {
		getProduct, err := c.PoductGetById(&models.ProductPrimaryKey{
			Id: k,
		})
		if err != nil {
			return err
		}
		fmt.Println(i, getProduct.Name, prodctsMap[k])
		i += 1
	}
	return nil
}

func (c *Controller) TransferCash(cash int, from *models.UserPrimaryKey, to *models.UserPrimaryKey) error {
	userFrom, err := c.UserGetById(from)
	if err != nil {
		log.Printf("Error while TransferCash => userFrom: %+v", err)
		return err
	}

	userTo, err := c.UserGetById(to)
	if err != nil {
		log.Printf("Error while TransferCash => userTo: %+v", err)
		return err
	}

	if userFrom.Balance < cash {
		return errors.New("In balance not enough cash")
	}

	UpdateFrom := &models.UserUpdate{
		Id:        userFrom.Id,
		FirstName: userFrom.FirstName,
		LastName:  userFrom.LastName,
		Balance:   userFrom.Balance - cash,
	}
	c.UserUpdate(UpdateFrom)

	UpdateTo := &models.UserUpdate{
		Id:        userTo.Id,
		FirstName: userTo.FirstName,
		LastName:  userTo.LastName,
		Balance:   userTo.Balance + cash,
	}

	c.UserUpdate(UpdateTo)
	return nil
}
