package sql

type Field struct {
	FieldName string `sql:"field_name"`
	FieldType string `sql:"field_type"`
	IsNull    bool   `sql:"is_null"`
	Default   string `sql:"default"`
	Comment   string `sql:"comment"`
}

type Table struct {
	TableName string   `sql:"table"`
	Fields    []Field  `sql:"fields"`
	ExtraInfo []string `sql:"extra_info"`
}
