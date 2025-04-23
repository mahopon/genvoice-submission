package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primarykey" json:"id"`
	CreatedDate time.Time `gorm:"not null" json:"created_date"`
	Name        string    `gorm:"not null" json:"name"`
	Username    string    `gorm:"not null;unique;index" json:"username"`
	Password    string    `gorm:"not null" json:"-"`
	Role        string    `gorm:"not null" json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	u.CreatedDate = time.Now()
	return nil
}
