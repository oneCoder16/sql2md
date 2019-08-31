package service

import (
	"regexp"
	"sql2md/pkg/sql"
	"strings"
)

type ParseService struct {
}

//从字符串中解析出默认值
func (srv *ParseService) parseDefaultVal(str string) string {
	reg := regexp.MustCompile(`DEFAULT '(.*?)'`)
	all := reg.FindAllSubmatch([]byte(str), -1)
	if len(all) > 0 {
		//获取最后一个
		return string(all[len(all)-1][1])
	} else {
		return ""
	}

}

//从字符串中解析出备注信息
func (srv *ParseService) parseCommentInfo(str string) string {
	reg := regexp.MustCompile(`COMMENT '(.*?)'`)
	all := reg.FindAllSubmatch([]byte(str), -1)
	if len(all) > 0 {
		//获取最后一个
		return string(all[len(all)-1][1])
	} else {
		return ""
	}

}

//从字符串中解析出 数据表字段信息
func (srv *ParseService) parseFieldInfo(str []byte) (fields []sql.Field) {
	all := strings.Split(string(str), "\n")

	for _, item := range all {
		if (strings.Index(item, "INDEX") != -1) || (strings.Index(item, "PRIMARY") != -1) {
			continue
		}

		fieldInfo := strings.TrimSpace(string(item))
		if fieldInfo == "" {
			continue
		}

		var field sql.Field
		field.Name = regexp.MustCompile("`(.*)`").FindString(fieldInfo)
		field.Type = regexp.MustCompile(`\s+.*?\s`).FindString(fieldInfo)
		field.AllowNull = strings.Contains(fieldInfo, "NOT NULL")
		field.Default = srv.parseDefaultVal(fieldInfo)
		field.Comment = srv.parseCommentInfo(fieldInfo)
		fields = append(fields, field)
	}

	return
}

//解析索引信息
func (src *ParseService) parseIndexdInfo(str []byte) (indexs []sql.Index) {
	all := strings.Split(string(str), "\n")

	for _, item := range all {
		if (strings.Index(item, "INDEX") == -1) && (strings.Index(item, "PRIMARY") == -1) {
			continue
		}

		indexInfo := strings.TrimSpace(string(item))
		index := sql.Index{}

		if strings.Contains(indexInfo, "PRIMARY KEY") {
			//主键信息
			index.Type = "PRIMARY KEY"
			index.Name = " "
		} else {
			//其他索引
			index.Type = regexp.MustCompile(`(\w*)\s`).FindString(indexInfo)
			index.Name = regexp.MustCompile("`(.*?)`").FindString(indexInfo)
		}

		index.Val = regexp.MustCompile(`\((.*)\)`).FindString(indexInfo)
		indexs = append(indexs, index)
	}

	return
}

//解析额外备注信息
func (src *ParseService) parseExtraInfo(str []byte) map[string]string {
	info := make(map[string]string)

	reg := regexp.MustCompile(`(\w+\s*\w+)\s*=\s*'?([\p{Han}\w]+)'?\s`)
	all := reg.FindAllSubmatch(str, -1)

	for _, item := range all {
		info[string(item[1])] = string(item[2])
	}

	return info
}

//从字符串中解析出 数据表结构
func (srv *ParseService) Str2Sql(data []byte) (tables []sql.Table) {
	reg := regexp.MustCompile(`CREATE TABLE ` + "`" + `(\w+)` + "`" + `\s*\(((?:.*\r\n)*?)\)(.*);`)
	all := reg.FindAllSubmatch(data, -1)

	for _, item := range all {
		table := sql.Table{}
		//获取数据表名
		table.TableName = string(item[1])

		//获取数据字段信息
		table.Fields = srv.parseFieldInfo(item[2])

		//获取索引信息
		table.Index = srv.parseIndexdInfo(item[2])

		//获取注释信息
		table.ExtraInfo = srv.parseExtraInfo(item[3])

		//获取数据表备注信息
		if val, ok := table.ExtraInfo["COMMENT"]; ok {
			table.Comment = val
			delete(table.ExtraInfo, "COMMENT")
		}

		tables = append(tables, table)
	}

	return tables
}
