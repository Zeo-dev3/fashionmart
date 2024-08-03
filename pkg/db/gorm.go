package db

import (
	"fmt"

	"github.com/Zeo-dev3/fashionmart/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	port := cfg.Postgres.Port
	host := cfg.Postgres.Host
	user := cfg.Postgres.User
	password := cfg.Postgres.Password
	name := cfg.Postgres.Name

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, name, port)
	dialect := postgres.Open(dsn)

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
