package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/alexsasharegan/dotenv"
)

const (
	usersApi string = "https://slack.com/api/usergroups.users.list"
	userApi  string = "https://slack.com/api/users.profile.get"
)

type Users struct {
	Users []string `json:"users"`
}

type Profile struct {
	Image       string `json:"image_512"`
	DisplayName string `json:"display_name"`
}

type User struct {
	Profile Profile `json:"profile"`
}

func getToken() string {
	var once sync.Once
	var t string

	once.Do(func() {
		err := dotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		t = os.Getenv("SLACK_TOKEN")

		if t == "" {
			log.Fatalf("please set env:SLACK_TOKEN")
		}
	})

	return t
}

func getGroupId() string {
	var once sync.Once
	var gID string

	once.Do(func() {
		err := dotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		gID = os.Getenv("GROUP_ID")

		if gID == "" {
			log.Fatalf("please set env:GROUP_IDd")
		}
	})

	return gID
}

func getClient() *http.Client {
	var once sync.Once
	var c *http.Client

	once.Do(func() {
		c = &http.Client{}
	})

	return c
}

func GetUsers() Users {
	c := getClient()
	d := url.Values{}
	d.Set("token", getToken())
	d.Set("usergroup", getGroupId())

	req, _ := http.NewRequest("POST", usersApi, bytes.NewBufferString(d.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.Do(req)

	if err != nil {
		log.Fatalf("fail to get users: %v", getGroupId())
	}

	var users Users
	json.NewDecoder(res.Body).Decode(&users)

	return users
}

func GetUser(userId string) User {
	c := getClient()
	d := url.Values{}
	d.Set("token", getToken())
	d.Set("user", userId)

	req, _ := http.NewRequest("POST", userApi, bytes.NewBufferString(d.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.Do(req)

	if err != nil {
		log.Fatalf("fail to get user: %v", userId)
	}

	var user User
	json.NewDecoder(res.Body).Decode(&user)

	return user
}
