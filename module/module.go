package module

import (
	"time"
)

//User是我们的用户表结构。
type User struct {
	ID        int64  // xorm默认自动递增
	Version   string `xorm:"varchar(200)"`
	Salt      string
	Username  string
	Password  string    `xorm:"varchar(200)"`
	Languages string    `xorm:"varchar(200)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

type Path struct {
	Id    int64  `json:"id" xorm:"pk autoincr"`
	Path  string `xorm:"path" json:"path"`
	Name  string `xorm:"name" json:"name"`
	Thumb string `xorm:"thumb" json:"thumb"`
}

type Top struct {
	Id    int64     `json:"id" xorm:"pk autoincr"`
	File  string    `json:"file" xorm:"file"`
	L     int       `json:"l" xorm:"int(11) not null default 0"`
	H     int       `json:"h" xorm:"int(11) not null default 0"`
	Time  time.Time `json:"time"`
	Thumb string    `xorm:"thumb" json:"thumb"`
	//Path  `xorm:"extends"`
}

type DirInfo struct {
	Name    string
	DirName string
}
