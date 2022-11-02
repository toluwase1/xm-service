package db

import (
	"company-service/models"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// DB provides access to the different db

//go:generate mockgen -destination=../mocks/auth_repo_mock.go -package=mocks github.com/decagonhq/meddle-api/db AuthRepository
type AuthRepository interface {
	CreateCompany(user *models.Company) (*models.Company, error)
	IsCompanyNameExist(name string) error
	FindCompanyByName(email string) (*models.Company, error)
	UpdateCompany(user *models.Company) error
}

type authRepo struct {
	DB *gorm.DB
}

func NewAuthRepo(db *GormDB) AuthRepository {
	return &authRepo{db.DB}
}

func (a *authRepo) CreateCompany(company *models.Company) (*models.Company, error) {
	err := a.DB.Create(company).Error
	if err != nil {
		return nil, fmt.Errorf("could not create company: %v", err)
	}
	return company, nil
}

func (a *authRepo) FindCompanyByName(username string) (*models.Company, error) {
	db := a.DB
	user := &models.Company{}
	err := db.Where("email = ? OR username = ?", username, username).First(user).Error
	if err != nil {
		return nil, fmt.Errorf("could not find user: %v", err)
	}
	return user, nil
}

func (a *authRepo) IsCompanyNameExist(name string) error {
	var count int64
	err := a.DB.Model(&models.Company{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return errors.Wrap(err, "gorm.count error")
	}
	if count > 0 {
		return fmt.Errorf("name already in use")
	}
	return nil
}

func (a *authRepo) UpdateCompany(user *models.Company) error {
	return nil
}
