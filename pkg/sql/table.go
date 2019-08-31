package sql

type Index struct {
	Name string
	Val  string
	Type string
}

type Field struct {
	Name      string
	Type      string
	AllowNull bool
	Default   string
	Comment   string
}

type Table struct {
	TableName string
	Comment   string
	Fields    []Field
	ExtraInfo map[string]string
	Index     []Index
}
