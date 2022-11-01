package services

import (
	"company-service/config"
	"company-service/db"
	apiError "company-service/errors"
	"company-service/models"
	"company-service/server/jwt"
	"log"
	"net/http"
)

//go:generate mockgen -destination=../mocks/auth_mock.go -package=mocks github.com/decagonhq/meddle-api/services AuthService
// AuthService interface
type AuthService interface {
	SignupUser(request *models.Company) (*models.Company, *apiError.Error)
}

// authService struct
type authService struct {
	Config   *config.Config
	authRepo db.AuthRepository
}

// NewAuthService instantiate an authService
func NewAuthService(authRepo db.AuthRepository, conf *config.Config) AuthService {
	return &authService{
		Config:   conf,
		authRepo: authRepo,
	}
}

func (a *authService) SignupUser(user *models.User) (*models.User, *apiError.Error) {
	err := a.authRepo.IsEmailExist(user.Email)
	if err != nil {
		return nil, apiError.New("email already exist", http.StatusBadRequest)
	}

	err = a.authRepo.IsPhoneExist(user.PhoneNumber)
	if err != nil {
		return nil, apiError.New("phone already exist", http.StatusBadRequest)
	}

	user.HashedPassword, err = GenerateHashPassword(user.Password)
	if err != nil {
		log.Printf("error generating password hash: %v", err.Error())
		return nil, apiError.New("internal server error", http.StatusInternalServerError)
	}

	_, err = jwt.GenerateToken(user.Email, a.Config.JWTSecret)
	if err != nil {
		return nil, apiError.New("internal server error", http.StatusInternalServerError)
	}
	//if err := a.sendVerifyEmail(token, user.Email); err != nil {
	//	return nil, err
	//}

	user.Password = ""
	user.IsEmailActive = false
	user, err = a.authRepo.CreateUser(user)

	if err != nil {
		log.Printf("unable to create user: %v", err.Error())
		return nil, apiError.New("internal server error", http.StatusInternalServerError)
	}

	return user, nil
}
