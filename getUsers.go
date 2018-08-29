package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"sync"
)

const (
	usersApi string = "https://slack.com/api/usergroups.users.list"
	userApi  string = "https://slack.com/api/users.profile.get"
)

type Users struct {
	Ok    bool
	Error string
	Users []string `json:"users"`
}

type User struct {
	Ok      bool
	Error   string
	Profile Profile `json:"profile"`
}

func getClient() *http.Client {
	var once sync.Once
	var c *http.Client

	once.Do(func() {
		c = &http.Client{}
	})

	return c
}

func GetUsers(t string, gID string) Users {
	c := getClient()
	d := url.Values{}
	d.Set("token", t)
	d.Set("usergroup", gID)

	req, _ := http.NewRequest("POST", usersApi, bytes.NewBufferString(d.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.Do(req)

	if err != nil {
		log.Fatalf("fail to get users: %v", gID)
	}

	var users Users
	json.NewDecoder(res.Body).Decode(&users)

	if !users.Ok {
		log.Fatalf("fail to get users: %v", users.Error)
	}

	return users
}

func GetUser(t string, uID string) User {
	c := getClient()
	d := url.Values{}
	d.Set("token", t)
	d.Set("user", uID)

	req, _ := http.NewRequest("POST", userApi, bytes.NewBufferString(d.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.Do(req)

	if err != nil {
		log.Fatalf("fail to get user: %v", uID)
	}

	var user User
	json.NewDecoder(res.Body).Decode(&user)

	if !user.Ok {
		log.Fatalf("fail to get user: %v", user.Error)
	}

	return user
}
