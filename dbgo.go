package dbgo

type Config struct {
	mappedStatements map[string]*MappedStatement
	driverName       string
	dataSourceName   string
}

func (t *Config) addMappedStatement(ms *MappedStatement) {
	t.mappedStatements[ms.id] = ms
}

func (t *Config) getMappedStatement(id string) (*MappedStatement, bool) {
	r, ok := t.mappedStatements[id]
	return r, ok
}
