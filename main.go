package main

import (
	"fmt"
)

func main() {
	us := GetUsers()

	for _, uID := range us.Users {
		fmt.Printf("UserId: %v\n", uID)

		u := GetUser(uID)

		GetUserImage(u.Profile)
	}
}
