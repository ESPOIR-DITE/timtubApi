package user_repo

import (
	"fmt"
	"testing"
	"time"
	"timtubeApi/domain"
)

func TestCreateUserTable(t *testing.T) {
	result := CreateUserTable()
	fmt.Println("result :", result)
}
func TestCreateUser(t *testing.T) {
	entity := domain.User{"dite@gmail.com", "espoir", "ditekemena", time.Now(), "0001"}
	result := CreateUser(entity)
	fmt.Println(result)
}
func TestGetUser(t *testing.T) {
	result := GetUser("espoir@gmail.com")
	fmt.Println(result)
}
func TestGetUsers(t *testing.T) {
	result := GetUsers()
	fmt.Println(result)
}
func TestSetDatabase(t *testing.T) {
	SetDatabase()
}
func TestDelete(t *testing.T) {
	result := Delete("espoir@gmail.com")
	fmt.Println(result)
}
