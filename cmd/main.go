package main

import (
	"fmt"
	"log"

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

	// Create Product \\

	// product, err := con.ProductCreate(&models.CreateProduct{
	// 	Name: "Nan",
	// 	Price: "213141",
	// })
	// if err != nil {
	// 	log.Println("product create err:", err)
	// 	return
	// }
	// fmt.Printf("%+v\n", product)

	// Get By Id Product \\

	// byid, err := con.ProductGetById(&models.ProductPrimaryKey{Id: "d87345a9-595d-47eb-bb14-378dcffc362e"})
	// if err != nil {
	// 	log.Printf("Error while GetById: %+v\n", err)
	// 	return
	// }
	// fmt.Println(byid)

	// Get List Product

	// var dataLimit int
	// var page int
	// var answer string
	// fmt.Printf("Input Data limit: ")
	// fmt.Scan(&dataLimit)
	// for {
	// 	fmt.Printf("Press e if you want end. Press n to continue: ")
	// 	fmt.Scan(&answer)
	// 	if answer == "n" {
	// 		fmt.Println("Input page:")
	// 		fmt.Scan(&page)
	// 		respProduct, err := con.ProductGetList(&models.ProductGetListRequest{
	// 			Offset: (page - 1) * dataLimit,
	// 			Limit:  dataLimit,
	// 		})

	// 		if err != nil {
	// 			fmt.Println(err)
	// 			continue
	// 		}

	// 		for _, product := range respProduct.Products {
	// 			fmt.Println(product)
	// 		}
	// 	} else if answer == "e" {
	// 		break
	// 	}
	// }

	// Product Update \\

	// product, err := con.ProductUpdate(&models.UpdateProduct{
	// 	Id:        "d87345a9-595d-47eb-bb14-378dcffc362e",
	// 	Name: "updatedname",
	// 	Price:  "updatedlname",
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Printf("%+v\n", product)

	// Product Delete \\

	// err = con.ProductDelete(&models.ProductPrimaryKey{Id: "d87345a9-595d-47eb-bb14-378dcffc362e"})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// USER \\

	// for {
	// 	var (
	// 		answer string
	// 	)
	// 	fmt.Printf("1.Create User\n2.Get By User Id\n3.Get User List\n4.Update User Info\n5.Delete User\n6.Exit\nEnter: ")
	// 	fmt.Scan(&answer)
	// 	if answer == "1" {
	// 		var (
	// 			name     string
	// 			lastname string
	// 		)
	// 		fmt.Printf("First Name: ")
	// 		fmt.Scan(&name)
	// 		fmt.Printf("Last Name: ")
	// 		fmt.Scan(&lastname)
	// 		user, err := con.UserCreate(&models.CreateUser{
	// 			FirstName: name,
	// 			LastName:  lastname,
	// 		})

			// if err != nil {
			// 	log.Println("user create err:", err)
			// 	continue
			// }

			// fmt.Printf("%+v\n", user)
	// 		fmt.Println("User Created successfully")
	// 	} else if answer == "2" {
	// 		var id string
	// 		fmt.Printf("Enter id of user: ")
	// 		fmt.Scan(&id)
			// byid, err := con.GetById(&models.UserPrimaryKey{Id: id})
			// if err != nil {
			// 	log.Printf("Error while GetById: %+v\n", err)
			// 	continue
			// }
			// fmt.Println(byid)
	// 	} else if answer == "3" {
			// var dataLimit int
			// var page int
			// var answer string
			// fmt.Printf("Input Data limit: ")
			// fmt.Scan(&dataLimit)
			// for {
			// 	fmt.Printf("Press e if you want end. Press n to continue: ")
			// 	fmt.Scan(&answer)
			// 	if answer == "n" {
			// 		fmt.Println("Input page:")
			// 		fmt.Scan(&page)
			// 		respUser, err := con.UserGetList(&models.UserGetListRequest{
			// 			Offset: (page - 1) * dataLimit,
			// 			Limit:  dataLimit,
			// 		})

			// 		if err != nil {
			// 			fmt.Println(err)
			// 			continue
			// 		}

			// 		for _, user := range respUser.Users {
			// 			fmt.Println(user)
			// 		}
			// 	} else if answer == "e" {
			// 		break
			// 	}
			// }
	// 	} else if answer == "4" {
	// 		var (
	// 			id           string
	// 			updatedname  string
	// 			updatedlname string
	// 		)
	// 		fmt.Printf("Enter id of user: ")
	// 		fmt.Scan(&id)
	// 		fmt.Printf("Enter New Name: ")
	// 		fmt.Scan(&updatedname)
	// 		fmt.Printf("Enter new Last Name: ")
	// 		fmt.Scan(&updatedlname)
			// user, err := con.UserUpdate(&models.UpdateUser{
			// 	Id:        id,
			// 	FirstName: updatedname,
			// 	LastName:  updatedlname,
			// })

			// if err != nil {
			// 	fmt.Println(err)
			// 	continue
			// }

			// fmt.Printf("%+v\n", user)
	// 		fmt.Println("Users information successfully updated !")
	// 	} else if answer == "5" {
	// 		var (
	// 			id string
	// 		)
	// 		fmt.Printf("Enter id of User: ")
	// 		fmt.Scan(&id)
			// err = con.UserDelete(&models.UserPrimaryKey{Id: id})
			// if err != nil {
			// 	fmt.Println(err)
			// 	continue
			// }
	// 		fmt.Println("User successfully deleted !")
	// 	} else {
	// 		break
	// 	}
	// }
}
