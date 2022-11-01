package main

type User struct {
	Model
	Name           string `json:"name" binding:"required,min=2"`
	Email          string `json:"email" gorm:"unique;not null" binding:"required,email"`
	PhoneNumber    string `json:"phone_number" gorm:"unique;default:null" binding:"required,e164"`
	Password       string `json:"password,omitempty" gorm:"-" binding:"required,min=8,max=15"`
	HashedPassword string `json:"-" gorm:"password"`
	IsEmailActive  bool   `json:"-"`
}

type Model struct {
	ID        uint  `json:"id" gorm:"primaryKey,autoIncrement"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeletedAt int64 `json:"deleted_at"`
}
