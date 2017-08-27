package models

import "time"

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}

type User struct {
	Model
	Email    string   `validator:"email" json:"email"`
	Password string   `json:"-"`
	Records  []Record `json:"records"`
}

type PublicUser struct {
	*User
}

type Record struct {
	Model
	UserID uint    `json:"userId"`
	Bmi    float32 `json:"bmi"`
}
