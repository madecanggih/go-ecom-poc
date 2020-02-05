package services

import (
	"architect/saras-go-poc/config"
	"architect/saras-go-poc/models"
	"github.com/jinzhu/gorm"
	"strconv"
)

func UserLogin(userLogin *models.Users, users *models.Users) {
	config.DB.Where(&userLogin).First(&users)
}

func UserRegister(users *models.Users) {
	config.DB.Create(&users)
}

func GetUsers(users *[]models.Users, id string) {
	if user_id, err := strconv.Atoi(id); err != nil {
		config.DB.Find(&users)
	} else {
		config.DB.Find(&users, user_id)
	}
}

func UserList(id string) []models.Users {
	users := []models.Users{}
	var res *gorm.DB

	if user_id, err := strconv.Atoi(id); err != nil {
		res = config.DB.Find(&users)
	} else {
		res = config.DB.Find(&users, user_id)
	}

	if res.Error != nil {
		return nil
	}
	return users
}

// func GetCheckout(data *models.Carts) {
// 	config.DB.Create(&data)
// }

// func AddCommit() {
// 	// config.DB.Find()
// }

// func AddCartdata(data *models.Carts) {
// 	config.DB.First(&data)
// }

// func GetCartById() {

// }

// func UpdateCart() {

// }

// func DeleteCart() {

// }

// func getCategory() {

// }

// func addCategory() {

// }

// func getCategoryById() {

// }

// func getInvoiceById() {

// }

// func getInvoiceByUserId() {

// }

// func getProduct() {

// }

// func addProduct() {

// }

// func getProductById() {

// }

// func getPromo() {

// }

// func addPromo() {

// }

// func getPromoByCode() {

// }

// func updatePromo() {

// }

// func deletePromo() {

// }

// func getStore() {

// }

// func addStore() {

// }

// func getStoreById() {

// }

// func updateStore() {

// }

// func deleteStore() {

// }

// func getTrolleyById() {

// }

// func getUsers(users *models.Users) {
// config.DB.Find(&users)
// }

// func getUsersById() {

// }

// func addWishlist() {

// }

// func getWishlistById() {

// }

// func deleteWishlist() {

// }
