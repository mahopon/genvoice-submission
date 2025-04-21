package model

import "time"

type User struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedDate time.Time `json:"created_date"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Password    string    `json:"-"`
}
