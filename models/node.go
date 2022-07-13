package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	Id       string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name     string `json:"name" gorm:"unique"`
	Uid      string `sql:"type:uuid`
	Type     string `json:"type"`
	Host     string `json:"host" gorm:"unique"`
	Location string `json:"location"`
	Disabled bool   `json:"disabled"`
	Custom   string `json:"custom"`
}

func (node *Node) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	node.Id = uuid.New().String()
	node.Disabled = false
	return
}

type NodeData struct {
	Name     string `form:"name" json:"name" xml:"name" binding:"required"`
	Uid      string `form:"uid" json:"uid" xml:"uid" binding:"required"`
	Type     string `form:"type" json:"type" xml:"type" binding:"required"`
	Host     string `form:"host" json:"host" xml:"host" binding:"required"`
	Location string `form:"location" json:"location" xml:"location" binding:"required"`
	Custom   string `form:"custom" json:"custom" xml:"custom" `
}

// func NewDefaultNode() Node {
// 	return Node{
// 		Disabled: false,
// 		Location: "Unknow",
// 		Host:     "Unknow",
// 		Type:     "Unknow",
// 	}
// }
