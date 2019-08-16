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

//强调格式
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

func NewTableRow(cols ...string) Glyph {
	row := NewRow()
	row.Insert(TableSeparatorGlyph)

	for _, val := range cols {
		row.Insert(NewCol(val))
		row.Insert(TableSeparatorGlyph)
	}

	return row
}
