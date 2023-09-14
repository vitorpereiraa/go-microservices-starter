package config

import (
	"fmt"

	"github.com/vitorpereiraa/go-microservices-starter/TodoList/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func LoadDB(settings Settings) (*gorm.DB, error)  {
	var (
		host= settings.DB_HOST 
		user= settings.DB_USER 
		password= settings.DB_PASSWORD 
		dbname= settings.DB_NAME 
		port= settings.DB_PORT 
		sslmode="disable" 
		timeZone="Europe/Lisbon"
	) 

	var connectionString = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", host, user, password, dbname, port, sslmode, timeZone) 

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, err
	}

	// Migrate the schema
	db.AutoMigrate(&domain.TodoList{})

	return db, nil
}