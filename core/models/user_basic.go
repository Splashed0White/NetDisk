package models

import "time"

type UserBasic struct {
	Id       int
	Identity string

	Name        string
	Password    string
	Email       string
	Created_at  time.Time  `gorm:"created"`
	Updated_at  time.Time  `gorm:"updated"`
	Deleted_at  *time.Time `gorm:"deleted"` // *time.Time，那么它的默认值是 nil，这个值可以存储在数据库中，表示没有删除
	TotalVolume int64      `gorm:"total_volume"`
	NowVolume   int64      `gorm:"now_volume"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
