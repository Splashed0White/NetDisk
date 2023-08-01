package models

import "time"

type ShareBasic struct {
	Id                     int
	Identity               string
	UserIdentity           string
	UserRepositoryIdentity string
	RepositoryIdentity     string
	ExpiredTime            int
	Click_num              int
	CreatedAt              time.Time  `gorm:"created"`
	UpdatedAt              time.Time  `gorm:"updated"`
	DeletedAt              *time.Time `gorm:"deleted"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
