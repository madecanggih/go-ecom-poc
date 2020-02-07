package resources

type (
	ErrorResponse struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}

	RegisterResponse struct {
		Status  bool         `json:"status"`
		Message string       `json:"message"`
		Data    RegisterData `json:"data"`
	}

	RegisterData struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Name     string `json:"name"`
		Address  string `json:"address"`
		Phone    string `json:"phone"`
		Image    string `json:"image"`
	}

	LoginResponse struct {
		Status  bool      `json:"status"`
		Message string    `json:"message"`
		Data    LoginData `json:"data"`
	}

	LoginData struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Name     string `json:"name"`
		Address  string `json:"address"`
		Phone    string `json:"phone"`
		Image    string `json:"image"`
		Token    string `json:"token"`
	}
)
