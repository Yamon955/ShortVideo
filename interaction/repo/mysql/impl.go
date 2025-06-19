package mysql

import (
	"gorm.io/gorm"
)

type dbClient struct {
	client *gorm.DB
}
