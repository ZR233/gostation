/*
@Time : 2019-06-26 15:22
@Author : zr
@File : base
@Software: GoLand
*/
package gorm

import (
	"github.com/jinzhu/gorm"
)

type BaseSqlDAO struct {
	db *gorm.DB
}

func newBase() *BaseSqlDAO {
	b := &BaseSqlDAO{}
	b.db = client
	return b
}

func (b *BaseSqlDAO) getDB() *gorm.DB {
	return b.db
}

func (b *BaseSqlDAO) getErr(db *gorm.DB) error {
	return DbError(db)
}
