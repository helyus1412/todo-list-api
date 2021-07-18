package projectdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/helyus1412/todo-list-api/configs"
	"github.com/helyus1412/todo-list-api/database/schema"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	Client   *sql.DB
	Database *gorm.DB
)

func Connect() {
	_ = godotenv.Load()
	src := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		configs.EnvFile.DbUser, configs.EnvFile.DbHost, configs.EnvFile.DbPort, configs.EnvFile.DbName,
	)
	db, err := gorm.Open("mysql", src)
	if err != nil {
		log.Println("MySQL: ", err)
		panic(err)
	}
	db.Debug()

	Database = db
	Client = db.DB()

	log.Println("database successfully configured")
	autoCreate := configs.EnvFile.DbAutoMigrate
	if autoCreate {
		fmt.Println("dropping and recreating all tables....")
		schema.AutoMigrate(db)
		fmt.Println("all tables recreated successfully!")
	}
}
