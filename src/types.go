package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type Server struct {
	dbAccess *sql.DB
	router   *mux.Router
}

type UserID struct {
	UserID string `json:"id"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}

type getUser struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Message  string `json:"message"`
	GotUser  bool   `json:"gotuser"`
}

type updateUser struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}

type UpdateUserResult struct {
	UserUpdated bool   `json:"userupdated"`
	Message     string `json:"message"`
}

type DeleteUserResult struct {
	UserDeleted bool   `json:"userdeleted"`
	UserID      string `json:"id"`
	Message     string `json:"message"`
}

type LoginUserResult struct {
	UserID       string `json:"id"`
	Username     string `json:"username"`
	UserLoggedIn bool   `json:"userloggedin"`
	Message      string `json:"message"`
}

type RegisterUserResult struct {
	UserCreated string `json:"usercreated"`
	Username    string `json:"username"`
	UserID      string `json:"id"`
	Message     string `json:"message"`
}

type dbConfig struct {
	UserName        string
	Password        string
	DatabaseName    string
	Port            string
	PostgresHost    string
	PostgresPort    string
	ListenServePort string
}

//Forgot password
type UserEmail struct {
	Email string `json:"email"`
}

type ForgotPasswordResult struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Message  string `json:"message"`
}

//advert crud
type PostAdvertisement struct {
	UserID            string `json:"userid"`
	IsSelling         string `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type PostAdvertisementResult struct {
	AdvertisementPosted bool   `json:"advertisementposted"`
	ID                  string `json:"id"`
	Message             string `json:"message"`
}

type UpdateAdvertisement struct {
	AdvertisementID   string `json:"id"`
	UserID            string `json:"userid"`
	IsSelling         string `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type UpdateAdvertisementResult struct {
	AdvertisementUpdated bool   `json:"advertisementupdated"`
	Message              string `json:"message"`
}

type DeleteAdvertisementResult struct {
	AdvertisementDeleted bool   `json:"advertisementdeleted"`
	AdvertisementID      string `json:"id"`
	Message              string `json:"message"`
}

type AdvertisementID struct {
	AdvertisementID string `json:"id"`
}

type getAdvertisement struct {
	AdvertisementID   string `json:"id"`
	UserID            string `json:"userid"`
	IsSelling         bool   `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
	Message           string `json:"message"`
}

type getAdvertisements struct {
	AdvertisementID   string `json:"id"`
	UserID            string `json:"userid"`
	IsSelling         bool   `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type TypeAdvertisementList struct {
	TypeAdvertisements []getAdvertisements `json:"typeadvertisements"`
}

type AdvertisementList struct {
	Advertisements []getAdvertisements `json:"advertisements"`
}

type GetUserAdvertisementResult struct {
	AdvertisementID   string `json:"advertisementid"`
	IsSelling         bool   `json:"isselling"`
	AdvertisementType string `json:"advertisementtype"`
	EntityID          string `json:"entityid"`
	Price             string `json:"price"`
	Description       string `json:"description"`
}

type UserAdvertisementList struct {
	UserAdvertisements []GetUserAdvertisementResult `json:"useradvertisements"`
}

type DeleteAdvertisementsResult struct {
	AdvertisementsDeleted bool   `json:"advertisementsdeleted"`
	Message               string `json:"message"`
}
type Config struct {
	ListenServePort string
}
