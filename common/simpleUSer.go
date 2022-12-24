package common

import "time"

type SimpleUser struct {
	SQLModel  `json:",inline"`
	LastName  string     `json:"last_name" gorm:"column:last_name;"`
	FirstName string     `json:"first_name" gorm:"column:first_name;"`
	Role      string     `json:"role" gorm:"column:role;"`
	Avatar    *Image     `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	CreatedAt *time.Time `json:"-" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"-" gorm:"column:updated_at;"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)
}