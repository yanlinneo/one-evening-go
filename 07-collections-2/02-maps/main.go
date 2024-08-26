package main

import "fmt"

var (
	Stats = map[string]int{}
)

func CreateUser(user string) {
	//Stats["create"]++
	Stats["create"] += 1
	fmt.Println("Creating user", user)
}

func UpdateUser(user string) {
	//Stats["update"]++
	Stats["update"] += 1
	fmt.Println("Updating user", user)
}

func UsersCreated() int {
	return Stats["create"]
}

func UsersUpdated() int {
	return Stats["update"]
}

func PurgeStats() {
	delete(Stats, "create")
	delete(Stats, "update")
	// Stats["create"] = 0
	// Stats["update"] = 0
}
