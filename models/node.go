package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	Id       string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name     string `json:"username"`
	Uid      string `sql:"type:uuid`
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
