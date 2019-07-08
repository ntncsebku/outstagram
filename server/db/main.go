package db

import (
	"flag"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"outstagram/server/models"
)

var dbConn *gorm.DB

func New() (*gorm.DB, error) {
	if dbConn == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Cannot load .env file")
		}

		var err error

		// If in TEST mode
		if flag.Lookup("test.v") != nil {
			dbConn, err = gorm.Open("mysql", "root:root@tcp(172.24.21.56:33060)/outstagram?charset=utf8mb4&parseTime=True&loc=Local")
		} else {
			dbConn, err = gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_URL"))
		}

		if err != nil {
			log.Fatal(err.Error())
		}

		dbConn.SingularTable(true)

		if os.Getenv("ENV") != "production" {
			//dbConn.LogMode(true)
			//dbConn.Debug()
			dbConn.AutoMigrate(
				&models.Comment{},
				&models.Commentable{},
				&models.Image{},
				&models.Message{},
				&models.Notification{},
				&models.NotifBoard{},
				&models.Post{},
				&models.PostImage{},
				&models.React{},
				&models.Reactable{},
				&models.Reply{},
				&models.Room{},
				&models.StoryBoard{},
				&models.Story{},
				&models.User{},
				&models.Viewable{})
		}
	}

	return dbConn, nil
}
