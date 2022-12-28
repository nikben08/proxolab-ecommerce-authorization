package contracts

type LoginRequest struct {
	Email    string `json:"eamil"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Email          string `json:"eamil"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
	Name           string `json:"name"`
	Sername        string `json:"sername"`
}

type LoginResponse struct {
}

type SignupResponse struct {
}
