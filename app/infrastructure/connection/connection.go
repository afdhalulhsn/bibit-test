package connection

import (
	"bibit/app/model"
	"bibit/internal/config"
	"bibit/internal/config/db"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DbConn *gorm.DB

func init() {
	cfg := config.GetConfig()
	var err error
	DbConn, err = Conn(cfg.Database.MovieImdb.Mysql)
	if err != nil {
		log.Fatal("Cannot Connect Database")
	}
	DbConn.Migrator().AutoMigrate(&model.LogApi{})
}

func Conn(d db.Database) (*gorm.DB, error) {
	pg := fmt.Sprintf("host= %v user=%v password=%v dbname=%v port=%v sslmode=%v", d.Host, d.Username, d.Password, d.Dbname, d.Port, "disable")
	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
