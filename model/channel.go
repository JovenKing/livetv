package model

type Channel struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	URL   string
	Proxy int // 1 代理 2 不代理
}
