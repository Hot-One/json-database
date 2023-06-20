package main

import (
	"fmt"

	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
)

func main() {
	cfg := config.Load()
	strg, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)

	con.TransferCash(7200000, &models.UserPrimaryKey{
		Id: "26fc445f-901e-453f-ae1b-81539e746601",
	}, &models.UserPrimaryKey{
		Id: "059f9715-a13b-4eb7-b2c0-eb969782cc08",
	})
	// active, err := con.ActiveUser()
	// if err != nil {
	// 	return err
	// }

	// s, err := con.HistoryUser(&models.OrderPrimaryKey{
	// 	Id: "27f2d317-7554-4285-8484-fdcdea004664",
	// })
	// fmt.Println(s.FirstName, s.LastName)
	// for k, v := range s.Products {
	// 	fmt.Println(k, v[0], v[1])
	// }
	// User(con)
	// Category(con)
	// Product(con)
	// Order(con)
	// OrderItem(con)
	// OrderPayment(con)
}

func User(con *controller.Controller) {
	con.UserCreate(&models.UserCreate{
		FirstName: "Xumoyun",
		LastName:  "Komilov",
		Balance:   3_000_000,
	})
}

func Category(con *controller.Controller) {
	con.CategoryCreate(&models.CategoryCreate{
		Name: "Ilmiy",
	})
}

func Product(con *controller.Controller) {

	var categoryId = "209237bc-aab9-47a9-aebd-58a0dcaa63f1"
	var products = []models.ProductCreate{
		{
			Name:         "Matematika",
			Price:        15_000,
			Discount:     0,
			DiscountType: "",
			CategoryID:   categoryId,
		},
		{
			Name:         "Learning Golang",
			Price:        200_000,
			Discount:     30,
			DiscountType: config.PercentDiscountType,
			CategoryID:   categoryId,
		},
		{
			Name:         "Clean Code",
			Price:        350_000,
			Discount:     40_000,
			DiscountType: config.FixDiscountType,
			CategoryID:   categoryId,
		},
	}

	for _, product := range products {
		con.ProductCreate(&product)
	}
}

func Order(con *controller.Controller) {
	con.OrderCreate(&models.OrderCreate{
		UserId: "c01f0fb9-2531-425f-85fb-480b8824f170",
	})
}

func OrderItem(con *controller.Controller) {
	var (
		orderId   = "8294f2dc-88f7-4794-aa2a-3f898486f555"
		productId = "d80ac6f0-42e5-4d82-aaf5-b9f28fa62ee3"
		count     = 7
	)

	con.AddOrderItem(&models.OrderItemCreate{
		OrderId:   orderId,
		ProductId: productId,
		Count:     count,
	})
}

func OrderPayment(con *controller.Controller) {

	var orderId = "27f2d317-7554-4285-8484-fdcdea004664"

	err := con.OrderPayment(&models.OrderPayment{
		OrderId: orderId,
	})

	fmt.Println(err)
}
