package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	AdminRole = "Admin"
	UserRole  = "User"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `gorm:"type:varchar(255)"`
	Email    string    `gorm:"type:varchar(320);unique"`
	Password string    `gorm:"type:varchar(255)"`
	Address  string    `json:"address"`
	Phone    string    `gorm:"type:varchar(16)"`
	Role     string    `gorm:"type:varchar(50);default:'user'"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	return
}

func MigrateUserEntity(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
