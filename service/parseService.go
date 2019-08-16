package service

import (
	"fmt"
	"regexp"
	"sql2md/pkg/sql"
)

type ParseService struct {
}

//从字符串中解析出 sql
func (srv *ParseService) Str2Sql(data string) []sql.Table {
	reg := regexp.MustCompile(`[:print:]`)
	all := reg.FindAllSubmatch([]byte(data), -1)

	fmt.Printf("%s", all)
	return nil
}
