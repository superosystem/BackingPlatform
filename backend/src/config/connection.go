package config

import (
	"log"

	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connect struct {
	MySQL *gorm.DB
}

func NewDBConn(cfg *Config) *Connect {
	return &Connect{
		MySQL: initMySQL(cfg),
	}
}

func initMySQL(cfg *Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.DB.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&entity.User{}, &entity.Campaign{}, &entity.CampaignImage{}, &entity.Transaction{}); err != nil {
		log.Fatal("error auto-migrations database", err.Error())
	}

	return db
}
