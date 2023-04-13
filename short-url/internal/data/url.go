package data

import (
	"github.com/jinzhu/gorm"
)

type Url struct {
	Id          int64  `gorm:"column:id;default:0;primary_key"`
	OriginalUrl string `gorm:"column:original_url;not null;default:'';comment:'原始链接'"`
	ShortUrl    string `gorm:"column:short_url;not null; default:'';comment:'短链接'"`
	CreateTime  int64  `gorm:"column:create_time;not null;default:0;comment:'创建时间''"`
}

func (u *Url) TableName() string {
	return "url"
}

type UrlQuery struct {
	Url
	orm *gorm.DB
}

func New(orm *gorm.DB) *UrlQuery {
	u := &UrlQuery{}
	u.orm = orm
	return u
}

func (u *UrlQuery) Orm() *gorm.DB {
	if u.orm == nil {
		u.orm = Orm
	}
	return u.orm
}

func (u *UrlQuery) CreateTable() {
	if !u.Orm().HasTable(&Url{}) {
		u.Orm().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='链接表'").CreateTable(&Url{})
	}
}

func (u *UrlQuery) Create(url *Url) error {
	return u.Orm().Create(url).Error
}

// GetById 使用 id 获取
func (u *UrlQuery) GetById(id int64) (url *Url, err error) {
	url = &Url{}
	err = u.Orm().Where("id=?", id).Take(url).Error
	return
}
