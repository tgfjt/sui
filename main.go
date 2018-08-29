package main

import (
	"time"
)

func main() {
	us := GetUsers()

	f := func(uID string) {
		GetUser(uID).Profile.GetUserImage()
	}

	for _, uID := range us.Users {
		go f(uID)
	}

	time.Sleep(time.Second)
}
