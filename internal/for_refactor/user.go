package main

import (
	"fmt"
	"strconv"
	"time"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type Shoes struct {
	ID    int
	Name  string
	Price float64
}

var quantityOfUsers = 1000
var quantityOfShoes = 1000

func main() {
	start := time.Now()
	usersList := createUsers()

	shoesList := createShoes()

	for _, user := range usersList {
		targetClient := usersList[len(usersList)-1].Name
		if user.Name == targetClient {
			fmt.Printf("Target client found: %s\n", targetClient)
		}
		for _, shoes := range shoesList {
			offersShoesToUser(user, shoes)
		}
	}
	elapsedNonOptimized := time.Since(start)

	//map to optimize
	start = time.Now()
	shoesMap := make(map[int]Shoes)
	usersmap := make(map[int]User)

	// Populando o mapa de sapatos
	for _, shoes := range shoesList {
		shoesMap[shoes.ID] = shoes
	}

	// Populando o mapa de usuários
	for _, user := range usersList {
		usersmap[user.ID] = user
	}
	shoes := shoesMap[quantityOfShoes]
	user := usersmap[quantityOfUsers]
	elapsedOptimized := time.Since(start)
	fmt.Printf("Elapsed time with two loops: %v\n", elapsedNonOptimized)
	fmt.Printf("Elapsed time with map optimization: %v\n", elapsedOptimized)
	fmt.Printf("User %d bought shoes %d\n", user.ID, shoes.ID)
}

func offersShoesToUser(user User, shoes Shoes) {
	fmt.Printf("Offering %s to %s\n", shoes.Name, user.Name)
	if user.Name == "Alice" && shoes.ID == quantityOfShoes {
		buyShoes(user, shoes)
	}
}

func buyShoes(user User, shoes Shoes) {
	println(user.Name, "bought", shoes.Name)
}

func createUsers() []User {
	var users []User
	for i := 1; i <= quantityOfUsers; i++ {
		user := User{
			ID:    i,
			Name:  "User" + strconv.Itoa(i),
			Email: "user" + strconv.Itoa(i) + "@example.com",
		}
		users = append(users, user)
	}
	return users
}

func createShoes() []Shoes {
	var shoesList []Shoes
	for i := 1; i <= quantityOfShoes; i++ {
		shoes := Shoes{
			ID:    i,
			Name:  "Shoes" + strconv.Itoa(i),
			Price: float64(i) * 10.0,
		}
		shoesList = append(shoesList, shoes)
	}
	return shoesList
}
