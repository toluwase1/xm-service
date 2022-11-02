package services

import (
	"company-service/config"
	"company-service/db"
	apiError "company-service/errors"
	"company-service/models"
	"log"
	"net/http"
)

//go:generate mockgen -destination=../mocks/auth_mock.go -package=mocks github.com/decagonhq/meddle-api/services AuthService
// CompanyService interface
type CompanyService interface {
	CreateCompany(request *models.Company) (*models.Company, *apiError.Error)
}

// companyService struct
type companyService struct {
	Config   *config.Config
	authRepo db.AuthRepository
}

// NewCompanyService instantiate an companyService
func NewCompanyService(authRepo db.AuthRepository, conf *config.Config) CompanyService {
	return &companyService{
		Config:   conf,
		authRepo: authRepo,
	}
}

func (a *companyService) CreateCompany(company *models.Company) (*models.Company, *apiError.Error) {
	err := a.authRepo.IsCompanyNameExist(company.Name)
	if err != nil {
		return nil, apiError.New("company name already exist, please choose another name", http.StatusBadRequest)
	}
	validTypes := []string{"Corporations", "NonProfit", "Cooperative", "Sole Proprietorship"}

	for i := 0; i < len(validTypes); i++ {
		if company.Type != validTypes[i] {
			return nil, apiError.New("Please provide a valid Type", http.StatusBadRequest)
		}
	}
	company, err = a.authRepo.CreateCompany(company)
	if err != nil {
		log.Printf("unable to create company: %v", err.Error())
		return nil, apiError.New("internal server error", http.StatusInternalServerError)
	}
	return company, nil
}
