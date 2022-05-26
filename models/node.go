package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	Id       string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name     string `json:"username"`
	User     string `json:"user"`
	Password string `json:"password"`
	Type     string `json:"types"`
	Host     string `json:"host" gorm:"unique"`
	Location string `json:"location"`
	Disabled bool   `json:"disabled"`
}

func (node *Node) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	node.Id = uuid.New().String()
	return
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
