// Package mysql 数据库操作包
package mysql

// DBClient 数据库操作接口对象
type DBClient interface {
}

// NewDBClient 新建一个数据库操作接口对象
func NewDBClient() DBClient {
	return &dbClient{
		db,
	}
}
