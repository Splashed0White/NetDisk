package models

import "time"

type Repository_pool struct {
	Id         int
	Identity   string
	Hash       string
	Name       string
	Ext        string
	Size       int64
	Path       string
	Created_at time.Time  `gorm:"created"`
	Updated_at time.Time  `gorm:"updated"`
	Deleted_at *time.Time `gorm:"deleted"`
}

func (table Repository_pool) TableName() string {
	return "repository_pool"
}
