package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github/joaltoroc/storicard/config"
	"github/joaltoroc/storicard/internal/transaction/entities"
)

func NewDatabase(cfg config.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Name)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	// AutoMigrate
	db.AutoMigrate(&entities.Transaction{})

	return db, nil
}
