package handlers

import (
	"architect/saras-go-poc/config"
	"architect/saras-go-poc/models"
	"architect/saras-go-poc/resources"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const InternalServerErrorMessage string = "Internal Server Error"
const RegistrationSuccessMessage string = "Registration Successful"
const LoginSuccessMessage string = "Login Successful"
const EmailRequiredMessage string = "Email is required"
const PasswordRequiredMessage string = "Password is required"
const LoginErrorMessage string = "Password is not correct or User is not registered"

type resultLists struct {
	Users []models.Users `json:"users"`
}

type handler struct {
	UserModel models.UserModelImpl
}

func NewHandler(u models.UserModelImpl) *handler {
	return &handler{u}
}

func (h *handler) GetIndex(c echo.Context) error {
	lists := h.UserModel.FindAll()
	u := &resultLists{lists}

	return c.JSON(200, u)
}

func (h *handler) GetDetail(c echo.Context) error {
	id := c.Param("id")
	user_id, _ := strconv.Atoi(id)
	u := h.UserModel.FindByID(user_id)
	return c.JSON(http.StatusOK, u)
}

func setErrorResponse(message string) resources.ErrorResponse {
	return resources.ErrorResponse{Status: false, Message: message}
}

func PostLogin(c echo.Context) error {
	login := new(models.Login)

	if err := c.Bind(&login); err != nil {
		res := setErrorResponse(InternalServerErrorMessage)
		return c.JSON(http.StatusInternalServerError, res)
	}

	if login.Email == "" {
		res := setErrorResponse(EmailRequiredMessage)
		return c.JSON(http.StatusBadRequest, res)
	}

	if login.Password == "" {
		res := setErrorResponse(PasswordRequiredMessage)
		return c.JSON(http.StatusBadRequest, res)
	}

	// TODO: Hash the pass
	//login.Password += "hash"

	userLogin := models.Users{Email: login.Email, Password: login.Password}
	var users models.Users

	config.DB.Where(&userLogin).First(&users)

	if users.ID == 0 {
		res := setErrorResponse(LoginErrorMessage)
		return c.JSON(http.StatusUnauthorized, res)
	}

	token, err := generateToken(users.ID, users.Name)
	if err != nil {
		res := setErrorResponse(InternalServerErrorMessage)
		return c.JSON(http.StatusInternalServerError, res)
	}

	data := resources.LoginData{ID: users.ID, Email: users.Email, Username: users.Username, Name: users.Name, Address: users.Address, Phone: users.Phone, Image: users.Image, Token: token}
	res := resources.LoginResponse{Status: true, Message: LoginSuccessMessage, Data: data}

	return c.JSON(http.StatusOK, res)
}

func PostRegister(c echo.Context) error {
	var register models.Users
	if err := c.Bind(&register); err != nil {
		res := setErrorResponse(InternalServerErrorMessage)
		return c.JSON(http.StatusInternalServerError, res)
	}
	//TODO Check unique
	config.DB.Create(&register)

	data := resources.RegisterData{ID: register.ID, Email: register.Email, Username: register.Username, Name: register.Name, Address: register.Address, Phone: register.Phone, Image: register.Image}
	res := resources.RegisterResponse{Status: true, Message: RegistrationSuccessMessage, Data: data}
	return c.JSON(http.StatusOK, res)
}

// func GetCheckout(c echo.Context) {
// 	userID := c.Param("user_id")

// }

// func PostCheckout(c echo.Context) {

// }

// func PostCart(c echo.Context) {

// }

// func GetCart(c echo.Context) {
// 	id := c.Param("id")
// }

// func PutCart(c echo.Context) {
// 	id := c.Param("id")
// }

// func DeleteCart(c echo.Context) {
// 	id := c.Param("id")
// }

// func GetCategory(c echo.Context) {
// 	id := c.Param("id")
// }

// func PutCategory(c echo.Context) {

// }

// func GetCategory(c echo.Context) {

// }

// func GetInvoice(c echo.Context) {
// 	id := c.Param("id")
// }

// func GetInvoiceHistory(c echo.Context) {
// 	userID := c.Param("user_id")
// }

// func GetProduct(c echo.Context) {
// 	id := c.Param("id")
// }

// func PostProduct(c echo.Context) {

// }

// func GetProduct(c echo.Context) {

// }

// func GetPromo(c echo.Context) {
// 	code := c.Param("code")
// }

// func PostPromo(c echo.Context) {

// }

// func PutPromo(c echo.Context) {
// 	code := c.Param("code")
// }

// func DeletePromo(c echo.Context) {
// 	code := c.Param("code")
// }

// func GetStore(c echo.Context) {
// 	id := c.Param("id")
// }

// func PostStore(c echo.Context) {

// }

// func PutStore(c echo.Context) {
// 	id := c.Param("id")
// }

// func DeleteStore(c echo.Context) {
// 	id := c.Param("id")
// }

// func GetTrolley(c echo.Context) {
// 	UserID := c.Param("user_id")
// }

func GetUsers(c echo.Context) error {
	id := c.Param("id")
	var users []models.Users

	if user_id, err := strconv.Atoi(id); err != nil {
		config.DB.Find(&users)
	} else {
		config.DB.Find(&users, user_id)
	}

	if len(users) == 0 {
		return c.JSON(http.StatusNoContent, "No Content")
	}
	if len(users) == 1 {
		return c.JSON(http.StatusOK, users) // Map to response & one data
	}
	return c.JSON(http.StatusOK, users) // Map to response & data
}

// func PostWishlist(c echo.Context) {

// }

// func GetWishlist(c echo.Context) {
// 	id := c.Param("id")
// }

// func DeleteWishlist(c echo.Context) {
// 	id := c.Param("id")
// }

func generateToken(ID uint, name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = ID
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}
