package dbgo

import (
	"database/sql"
	"errors"
)

type SqlSession interface {
	Query(id string, data interface{}) (*sql.Rows, error)
	Exec(id string, data interface{}) (sql.Result, error)
	BeginTx()
}

type DefaultSqlSession struct {
	config *Config
	dbPool *sql.DB
}

func (t *DefaultSqlSession) Query(id string, data interface{}) (*sql.Rows, error) {
	ms, ok := t.config.getMappedStatement(id)
	if !ok {
		return nil, errors.New("no mapped statement for id:" + id)
	}

	sql := ms.getSql(data)
	return t.dbPool.Query(sql)
}

func (t *DefaultSqlSession) Exec(id string, data interface{}) (sql.Result, error) {
	ms, ok := t.config.getMappedStatement(id)
	if !ok {
		return nil, errors.New("no mapped statement for id:" + id)
	}

	sql := ms.getSql(data)
	return t.dbPool.Exec(sql)
}

func (t *DefaultSqlSession) BeginTx() {

}

func NewSqlSession(config *Config) (SqlSession, error) {
	if config == nil {
		return nil, errors.New("db config is null")
	}

	dbPool, err := openDBPool(config.driverName, config.dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DefaultSqlSession{
		config: config,
		dbPool: dbPool,
	}, nil
}

func openDBPool(driverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}

	return db, err
}
