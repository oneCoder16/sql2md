package markdown

import (
	"errors"
)

//简单的一列
type Glyph interface {
	Draw() string
	Insert(glyph Glyph) error
}

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
//规定行的列数
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
