package dbgo

import (
	"bytes"
	"text/template"
)

type MappedStatement struct {
	//namespace+id
	id        string
	sqlSource SqlSource
}

func (t MappedStatement) getSql(data interface{}) string {
	return t.sqlSource.getSql(data)
}

type SqlSource interface {
	getSql(data interface{}) string
}

type DynamicSqlSource struct {
	tmpl *template.Template
}

func (t DynamicSqlSource) getSql(data interface{}) string {
	var buff bytes.Buffer
	t.tmpl.Execute(&buff, data)
	return buff.String()
}

type StaticSqlSource struct {
	sql string
}

func (t StaticSqlSource) getSql(_ interface{}) string {
	return t.sql
}
