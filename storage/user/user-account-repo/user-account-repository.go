package user_account_repo

import (
	"fmt"
	"github.com/google/uuid"
	"timtubeApi/config"
	"timtubeApi/domain"
)

func CreateUserAccountTable() bool {
	var tableData = &domain.User{}
	err := config.GetDatabase().AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetUserAccountDatabase() {
	err := config.GetDatabase().Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.User{})
	if err != nil {
		fmt.Println("User Account database config not set")
	} else {
		fmt.Println("User Account database config set successfully")
	}
}
func CreateUserAccount(entity domain.UserAccount) *domain.UserAccount {
	var tableData = &domain.UserAccount{}
	id := "UA-" + uuid.New().String()
	user := domain.UserAccount{id, entity.Email, entity.Password, entity.Date}
	config.GetDatabase().Create(user).Find(&tableData)
	return tableData
}
func GetUserAccount(customerId string) domain.UserAccount {
	entity := domain.UserAccount{}
	config.GetDatabase().Where("CustomerId = ?", customerId).Find(&entity)
	return entity
}
func GetUserAccounts() []domain.UserAccount {
	var entity []domain.UserAccount
	config.GetDatabase().Find(&entity)
	return entity
}
func DeleteUserAccount(customerId string) bool {
	entity := domain.User{}
	config.GetDatabase().Where("CustomerId = ?", customerId).Delete(&entity)
	if entity.Email == "" {
		return true
	}
	return false
}
