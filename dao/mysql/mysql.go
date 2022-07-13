package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"systemMoniter-Server/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var sqlDB *sql.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	// 也可以使用MustConnect连接不成功就panic

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed!", zap.Error(err))
		return
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return
}

func Migrate() {
	db.AutoMigrate(&models.Info{})
	db.AutoMigrate(&models.Node{})
	db.AutoMigrate(&models.User{})
	zap.L().Info("Database Migration Completed!")
}

func InsertUser(user *models.User) error {
	result1 := db.Model(&models.User{}).Where("user = ?", user.User).First(&user)
	if result1.RowsAffected > 0 {
		err := errors.New("Duplicate user find")
		return err
	}

	result := db.Model(&models.User{}).Create(user)
	return result.Error
}

func FindUser(username string, user *models.User) error {
	result := db.Model(&models.User{}).Where("user = ?", username).First(&user)
	if result.RowsAffected <= 0 {
		err := errors.New("No user find")
		return err
	}
	return result.Error
}

func FindUserByUid(uid string, user *models.User) error {
	result := db.Model(&models.User{}).Where("id = ?", uid).First(&user)
	if result.RowsAffected <= 0 {
		err := errors.New("No user find")
		return err
	}
	return result.Error
}

func InsertNode(node *models.Node) error {
	result1 := db.Model(&models.Node{}).Where("name = ? and host = ?", node.Name, node.Host).First(&node)
	fmt.Println(result1.RowsAffected)
	if result1.RowsAffected > 0 {
		err := errors.New("Duplicate node find")
		return err
	}
	result := db.Create(node)
	return result.Error
}

func FindNodeByNameAndHost(name string, host string, node *models.Node) error {
	result := db.Model(&models.Node{}).Where("name = ? and host = ?", name, host).First(&node)
	if result.RowsAffected <= 0 {
		err := errors.New("No node find")
		return err
	}
	if node.Disabled == true {
		err := errors.New("Node is disabled")
		return err
	}
	return result.Error
}

func FindNodeByID(id string, node *models.Node) error {
	result := db.Model(&models.Node{}).Where("id = ? ", id).First(&node)
	if result.RowsAffected <= 0 {
		err := errors.New("No node find")
		return err
	}
	if node.Disabled == true {
		err := errors.New("Node is disabled")
		return err
	}
	return result.Error
}

func InsertInfoByNameAndHost(name string, host string, info *models.Info) error {
	var node models.Node
	result := db.Model(&models.Node{}).Where("name = ? and host = ?", name, host).First(&node)
	if result.RowsAffected <= 0 {
		err := errors.New("No node find")
		return err
	}
	if node.Id == "" {
		err := errors.New("No node find")
		return err
	}
	if node.Disabled == true {
		err := errors.New("Node is disabled")
		return err
	}
	db.Model(&models.Info{}).Where("node_id = ?", node.Id).Update("is_latest", false)
	info.NodeId = node.Id
	result2 := db.Create(info)
	return result2.Error
}

func InsertInfoByNodeID(nodeID string, info *models.Info) error {
	var node models.Node
	result := db.Model(&models.Node{}).Where("node_id = ?").First(&node)
	if result.RowsAffected <= 0 {
		err := errors.New("No node find")
		return err
	}
	if node.Id == "" {
		err := errors.New("No node find")
		return err
	}
	if node.Disabled == true {
		err := errors.New("Node is disabled")
		return err
	}
	db.Model(&models.Info{}).Where("node_id = ?", nodeID).Update("is_latest", false)
	result2 := db.Create(info)
	return result2.Error
}

func FindNodeIDByNameAndHost(name string, host string) (error, string) {
	var node models.Node
	result := db.Model(&models.Node{}).Where("name = ? and host = ?", name, host).First(&node)
	if result.RowsAffected <= 0 {
		err := errors.New("No node find")
		return err, ""
	}
	if node.Id == "" {
		err := errors.New("No node find")
		return err, ""
	}
	if node.Disabled == true {
		err := errors.New("Node is disabled")
		return err, ""
	}
	return nil, node.Id
}

func FindAllNodeStatus() (error, []models.Status) {
	var allNodeStatus []models.Status
	result := db.Raw("SELECT n.name,n.type,n.host,n.location,n.disabled, i.* FROM `nodes` as n left JOIN `infos` as i on n.id = i.node_id ORDER BY lastest_time DESC").Scan(&allNodeStatus)
	return result.Error, allNodeStatus
}
