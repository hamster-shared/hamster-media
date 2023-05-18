package initialization

import (
	"fmt"
	"getNews/application"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

func InitDB() {
	user := "root"
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "t_",
		},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err))
	}
	application.SetBean[*gorm.DB]("db", db)
}
