package db

import (
	"fmt"
	"time"

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

	// Mengakses objek *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Mengatur connection pooling
	sqlDB.SetMaxIdleConns(10)                  // Jumlah koneksi idle maksimum
	sqlDB.SetMaxOpenConns(100)                 // Jumlah koneksi terbuka maksimum
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Masa hidup koneksi

	return db, nil
}
