package auth

type Routes struct {
	NewAuthProfile string
}

func NewRoutes() Routes {
	return Routes{
		NewAuthProfile: "/new-auth-profile",
	}
}

type NewAuthProfileRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewAuthProfileResponse struct {
	UUID            string `json:"uuid"`
	IsEmailValid    bool   `json:"is_email_valid"`
	IsPasswordValid bool   `json:"is_password_valid"`
}
