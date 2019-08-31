package service

import (
	. "sql2md/pkg/markdown"
	"sql2md/pkg/sql"
)

type TransferService struct {
}

func (srv *TransferService) Sql2md(tables []sql.Table) string {
	doc := NewDocument()

	for _, table := range tables {
		one := NewDocument()
		//数据表信息
		one.Insert(NewTitleRow(table.TableName))
		one.Insert(NewRemarkRow(table.Comment))

		//字段信息
		title := []string{"字段", "类型", "必填", "默认", "注释"}
		mdTable := NewTable(title)
		for _, field := range table.Fields {
			allowNull := "否"
			if field.AllowNull {
				allowNull = "是"
			}
			mdTable.Insert(NewTableRow(field.Name, field.Type, allowNull, field.Default, field.Comment))
		}
		one.Insert(mdTable)

		//索引信息
		one.Insert(NewUnorderedRow("索引:"))
		one.Insert(NewEmptyRow())
		title = []string{"名字", "字段", "类型"}
		indexTable := NewTable(title)
		for _, index := range table.Index {
			indexTable.Insert(NewTableRow(index.Name, index.Val, index.Type))
		}
		one.Insert(indexTable)
		one.Insert(NewRow())

		//备注信息
		one.Insert(NewUnorderedRow("备注:"))
		one.Insert(NewEmptyRow())
		title = []string{"属性", "值"}
		remarkTable := NewTable(title)
		for field, val := range table.ExtraInfo {
			remarkTable.Insert(NewTableRow(field, val))
		}
		one.Insert(remarkTable)

		one.Insert(NewRow())
		doc.Insert(one)
	}

	return doc.Draw()
}
