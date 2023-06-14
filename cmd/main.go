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
	for{
		var (
			answer	  string
		)
		fmt.Printf("1.Create User\n2.Get By User Id\n3.Get User List\n4.Update User Info\n5.Delete User\n6.Exit\nEnter: ")
		fmt.Scan(&answer)
		if answer == "1"{
			var (
				name string
				lastname string
			)
			fmt.Printf("First Name: ")
			fmt.Scan(&name)
			fmt.Printf("Last Name: ")
			fmt.Scan(&lastname)
			con.UserCreate(&models.CreateUser{
				FirstName: name,
				LastName:  lastname,
			})
		
			if err != nil {
				log.Println("user create err:", err)
				return
			}
		
			fmt.Println("User Created successfully")
		}else if answer == "2"{
			var id string
			fmt.Printf("Enter id of user: ")
			fmt.Scan(&id)
			byid, err := con.GetById(id)
			if err != nil{
				log.Println("Error while GetById: %+v", err)
			}
			fmt.Println(byid)
		}else if answer == "3"{
			var dataLimit int
			var page      int
			var answer	  string 
			fmt.Printf("Input Data limit: ")
			fmt.Scan(&dataLimit)
			for {
				fmt.Printf("Press e if you want end. Press n to continue: ")
				fmt.Scan(&answer)
				if answer == "n"{
					fmt.Println("Input page:")
					fmt.Scan(&page)
					respUser := con.UserGetList(&models.UserGetListRequest{
						Offset: (page - 1) * dataLimit,
						Limit:  dataLimit,
					})
			
					for _, user := range respUser.Users {
						fmt.Println(user)
					}
				}else if answer == "e"{
					break
				}
			}
		}else if answer == "4"{
			var (
				id string
				updatedname string
				updatedlname string
			)
			fmt.Printf("Enter id of user: ")
			fmt.Scan(&id)
			fmt.Printf("Enter New Name: ")
			fmt.Scan(&updatedname)
			fmt.Printf("Enter new Last Name: ")
			fmt.Scan(&updatedlname)
			con.UserUpdate(&models.User{
				Id: id,
				FirstName: updatedname,
				LastName: updatedlname,
			})
			fmt.Println("Users information successfully updated !")
		}else if answer == "5"{
			var (
				id string
			)
			fmt.Printf("Enter id of User: ")
			fmt.Scan(&id)
			con.UserDelete(id)
			fmt.Println("User successfully deleted !")
		}else{
			break
		}
	}
}
