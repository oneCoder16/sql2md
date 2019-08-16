package markdown

import (
	"testing"
)

func TestMarkDown(t *testing.T) {
	doc := NewDocument()
	doc.Insert(NewTitleRow("pm_user_info"))
	doc.Insert(NewRemarkRow("存储用户信息"))

	title := []string{"字段", "类型", "必填", "默认", "注释"}
	table := NewTable(title)
	table.Insert(NewTableRow("id", "int(10)", "否", "", ""))
	table.Insert(NewTableRow("id", "int(10)", "否", "", ""))
	table.Insert(NewTableRow("id", "int(10)", "否", "", ""))
	table.Insert(NewTableRow("id", "int(10)", "否", "", ""))
	table.Insert(NewTableRow("id", "int(10)", "否", "", ""))

	doc.Insert(table)
	doc.Insert(NewRow())
	doc.Insert(NewUnorderedRow("备注:"))

	t.Log(doc.Draw())
}
