package config

import (
	cdm "api/src/config/db"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host          string
	Port          int
	Username      string
	Password      string
	DatabasesName string
	TimeZone      string
}

var DB = Init()

func BuildDBConfig() *Config {
	Config := Config{
		Host:          "satao.db.elephantsql.com",
		Port:          5432,
		Username:      "ixarmjil",
		Password:      "rHJyA70TEFmCOETrZ56kXAbFWyCciZiC",
		DatabasesName: "ixarmjil",
	}
	return &Config
}
func DbURL(conf *Config) string {
	// return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Bangkok",
		conf.Host,
		conf.Username,
		conf.Password,
		conf.DatabasesName,
		conf.Port,
	)
}

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(DbURL(BuildDBConfig())), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&cdm.User{})
	return db
}
