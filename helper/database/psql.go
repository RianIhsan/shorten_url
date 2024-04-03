package database

import (
	"fmt"
	"log"

	"github.com/RianIhsan/shorten_url/config"
	"github.com/RianIhsan/shorten_url/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectionDB(cnf config.APPConfig) *gorm.DB {
	psqlConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		cnf.DBConf.DbHost, cnf.DBConf.DbUser, cnf.DBConf.DbPass, cnf.DBConf.DbName, cnf.DBConf.DbPort)
	db, err := gorm.Open(postgres.Open(psqlConn), &gorm.Config{})
	if err != nil {
		log.Fatal("error open connection")
		return nil
	}
	log.Println("database connected")
	return db
}

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&entities.MstURL{}); err != nil {
		log.Fatal("error migrate table")
	}
	log.Println("migrate succesfully")
}
