package models

import "github.com/google/uuid"

type Company struct {
	Id                uuid.UUID
	Name              string `json:"name"  gorm:"unique;default:null" binding:"required"`
	Description       string `json:"description,omitempty" binding:"max=3000"`
	AmountOfEmployees int    `json:"amount_of_employees"  gorm:"unique" binding:"required"`
	IsRegistered      bool   `json:"is_registered" gorm:"unique;not null" binding:"required,is_registered"`
	Type              string `json:"type" binding:"required"`
}

/*
• ID (uuid) required
• Name (15 characters) required - unique
• Description (3000 characters) optional
• Amount of Employees (int) required
• Registered (boolean) required
• Type (Corporations | NonProfit | Cooperative | Sole Proprietorship) required
*/
