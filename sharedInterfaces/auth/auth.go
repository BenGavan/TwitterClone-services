package auth

type Routes struct {
	NewAuthProfile string
}

func NewRoutes() Routes {
	return Routes{
		NewAuthProfile: "/new-auth-profile",
	}
}

type Service struct {
	routes Routes
}

func NewService() Service {
	return Service{routes: NewRoutes()}
}

//func init() {
//	routes = NewRoutes()
//}

func (s *Service) makeJsonJsonRequest(url string, reqData interface{}, respData interface{}) error {
	return nil
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

func (s *Service) NewAuthProfile(email string, barePassword string) NewAuthProfileResponse {
	reqData := NewAuthProfileRequest{
		Email:    email,
		Password: barePassword,
	}

	var respData NewAuthProfileResponse

	// make json request
	err := s.makeJsonJsonRequest(s.routes.NewAuthProfile, reqData, &respData)
	if err != nil {

	}
	return respData
}
