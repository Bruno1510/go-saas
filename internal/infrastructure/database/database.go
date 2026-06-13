package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(
	host string,
	port string,
	user string,
	pass string,
	dbName string,
) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		pass,
		dbName,
	)

	return gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
}