package entity

import (
	"example/fiber/utils"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               uint   `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Name             string `json:"name" gorm:"type:varchar(60)"`
	Email            string `json:"email" gorm:"type:varchar(255);unique; NOT NULL"`
	Password         string
	Phone            string         `json:"phone" gorm:"type:varchar(60);unique;NULL"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	EmailConfirmedAt time.Time      `json:"email_confirmed_at,omitempty" gorm:"autoCreateTime:false"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

func (u *User) EncryptPassword() {
	u.Password = utils.HashPassword(u.Password)
}
