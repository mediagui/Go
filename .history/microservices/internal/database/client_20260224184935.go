package database

import (
	"fmt"

	"gorm.io/gorm"
)

type DatabaseClient interface {
	Ready() bool
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient() (DatabaseClient, error) {
	dsn := fmt.Sprintf("host:%s user=%s password=%s dmbname=%s sslmode=%s")
}
