package models

import "time"

type UserRepository struct {
	Id                 int
	Identity           string
	UserIdentity       string
	ParentId           int64
	RepositoryIdentity string
	Name               string
	Ext                string
	CreatedAt          time.Time  `gorm:"created"`
	UpdatedAt          time.Time  `gorm:"updated"`
	DeletedAt          *time.Time `gorm:"deleted"`
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
