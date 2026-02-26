package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseClient interface {
	Ready() bool
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient() (DatabaseClient, error) {
	dsn := fmt.Sprintf("host:%s user=%s password=%s dbname=%s port=%d sslmode=%s", "localhost", "postgres", "postgres", "postgres", 5432, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}){
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom."
		}
	}
	if err != nil {
		return nil, err
	}
	return &Client{DB: db}, nil
}
