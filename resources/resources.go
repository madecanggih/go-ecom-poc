package resources

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type RegisterResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    RegisterData `json:"data"`
}

type RegisterData struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Image    string `json:"image"`
}

type LoginResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    LoginData `json:"data"`
}

type LoginData struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Image    string `json:"image"`
	Token    string `json:"token"`
}
