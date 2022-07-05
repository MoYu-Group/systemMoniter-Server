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
	result1 := db.Where("user = ?", user.User).First(&user)
	if result1.RowsAffected > 0 {
		err := errors.New("Duplicate user find")
		return err
	}

	result := db.Create(user)
	return result.Error
}

func FindUser(username string, user *models.User) error {
	result := db.Where("user = ?", username).First(&user)
	if result.RowsAffected <= 0 {
		err := errors.New("No user find")
		return err
	}
	return result.Error
}
