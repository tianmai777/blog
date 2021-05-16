package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/tianmai777/blog/global"
	"github.com/tianmai777/blog/pkg"

	_ "github.com/go-sql-driver/mysql"
)

type Model struct {
	ID         int64  `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	UpdatedBy  string `json:"updated_by"`
	CreatedOn  int64  `json:"created_on"`
	ModifiedOn int64  `json:"modified_on"`
	DeletedOn  int64  `json:"deleted_on"`
	IsDel      int8   `json:"is_del"`
}

func NewDBEngine(conf *pkg.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(conf.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			conf.Username,
			conf.Passport,
			conf.Host,
			conf.DBName,
			conf.Charset,
			conf.ParseTime))
	if err != nil {
		return nil, errors.Wrap(err, "new db engine failed")
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(conf.MaxIdleConns)
	db.DB().SetMaxOpenConns(conf.MaxOpenConns)
	return db, nil
}
