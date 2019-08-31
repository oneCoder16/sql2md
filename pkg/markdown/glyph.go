package markdown

import (
	"errors"
)

var (
	RemarkGlyph         Glyph = NewCol("> ")
	TitleGlyph          Glyph = NewCol("---")
	TableSeparatorGlyph Glyph = NewCol("|")
	UnorderedGlyph      Glyph = NewCol("- ")
)

type Glyph interface {
	Draw() string
	Insert(glyph Glyph) error
}

//一列
type ColGlyph struct {
	content string
}

func NewCol(content string) Glyph {
	return &ColGlyph{
		content: content,
	}
}

func (this *ColGlyph) Draw() string {
	return this.content
}

func (this *ColGlyph) Insert(glyph Glyph) error {
	return errors.New("can not insert ")
}

//简单的一行
type RowGlyph struct {
	rows []Glyph
}

func NewRow() Glyph {
	return &RowGlyph{}
}

func (this *RowGlyph) Draw() string {
	content := ""
	for _, val := range this.rows {
		content += val.Draw()
	}

	content += "\n"

	return content
}

func (this *RowGlyph) Insert(glyph Glyph) error {

	this.rows = append(this.rows, glyph)
	return nil
}

//备注节点
type RemarkRow struct {
	content string
}

func NewRemarkRow(content string) Glyph {
	return &RemarkRow{
		content: content,
	}
}

func (this *RemarkRow) Draw() string {
	return RemarkGlyph.Draw() + this.content + "\n\n"
}

func (this *RemarkRow) Insert(glyph Glyph) error {
	return errors.New("cant not insrt")
}

//强调节点
type UnorderedRow struct {
	content string
}

func NewUnorderedRow(content string) Glyph {
	return &UnorderedRow{
		content: content,
	}
}

func (this *UnorderedRow) Draw() string {
	return UnorderedGlyph.Draw() + this.content + "\n"
}

func (this *UnorderedRow) Insert(glyph Glyph) error {
	return errors.New("cant not insrt")
}

//标题
type TitleRow struct {
	content string
}

func NewTitleRow(content string) Glyph {
	return &TitleRow{
		content: content,
	}
}

func (this *TitleRow) Draw() string {
	return this.content + "\n" + TitleGlyph.Draw() + "\n"
}

func (this *TitleRow) Insert(glyph Glyph) error {
	return errors.New("cant not insrt")
}

//表格
type Table struct {
	rows  []Glyph
	title []string
}

func NewTable(title []string) Glyph {
	return &Table{
		title: title,
	}
}

func (this *Table) Draw() string {
	var table string = ""
	var tableTitle string = TableSeparatorGlyph.Draw()
	var tableSeparator string
	var content string

	for _, val := range this.title {
		tableTitle += val + TableSeparatorGlyph.Draw()
		tableSeparator += "---" + TableSeparatorGlyph.Draw()
	}

	for _, val := range this.rows {
		content += val.Draw()
	}

	table += tableTitle + "\n" + tableSeparator + "\n" + content + "\n"

	return table
}

func (this *Table) Insert(row Glyph) error {
	this.rows = append(this.rows, row)
	return nil
}

//表格一行
func NewTableRow(cols ...string) Glyph {
	row := NewRow()
	row.Insert(TableSeparatorGlyph)

	for _, val := range cols {
		row.Insert(NewCol(val))
		row.Insert(TableSeparatorGlyph)
	}

	return row
}

//根节点
type Document struct {
	glyphs []Glyph
}

func NewDocument() Glyph {
	return &Document{}
}

func (this *Document) Draw() string {
	content := ""
	for _, val := range this.glyphs {
		content += val.Draw()
	}

	content += "\n"

	return content
}

func (this *Document) Insert(glyph Glyph) error {
	this.glyphs = append(this.glyphs, glyph)
	return nil
}

//空行
func NewEmptyRow() Glyph {
	return NewRow()
}
