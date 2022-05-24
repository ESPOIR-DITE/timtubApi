package domain

import (
	"errors"
	"net/http"
	"time"
)

type User struct {
	Email     string    `json:"email" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	BirthDate time.Time `json:"birthDate"`
	RoleId    string    `json:"roleId"`
}

type Role struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (role Role) Bind(r *http.Request) error {
	//TODO implement me
	if role.Name == "" {
		return errors.New("missing required fields")
	}
	return nil
}

type UserAccount struct {
	CustomerId string    `json:"customerId" gorm:"primaryKey"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Date       time.Time `json:"date"`
}

type UserSubscription struct {
	Id             string    `json:"id" gorm:"primaryKey"`
	UserId         string    `json:"userId"`
	Stat           string    `json:"stat"`
	SubscriptionId string    `json:"subscriptionId"`
	Date           time.Time `json:"date"`
}
type UserVideo struct {
	CustomerId string    `json:"customerId" gorm:"primaryKey"`
	VideoId    string    `json:"videoId"`
	Date       time.Time `json:"date"`
}
