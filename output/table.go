package output

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// FormattedText represents a piece of formatted text.
type FormattedText struct {
	richText string
	rawText  string
}

// Add adds some text. The `text` argument should be the *unformatted* text, and the `color` argument should specify
// how that text should be formatted.
func (ft *FormattedText) Add(text string, color *color.Color) {
	ft.rawText += text
	if color != nil {
		ft.richText += color.Sprint(text)
	} else {
		ft.richText += text
	}
}

func NewFormattedText(text string, color *color.Color) FormattedText {
	var richText string
	if color != nil {
		richText = color.Sprint(text)
	} else {
		richText = text
	}
	return FormattedText{
		rawText:  text,
		richText: richText,
	}
}

// FormattedRow is a row of `FormattedText` objects (each of which represents a column).
type FormattedRow struct {
	columns []FormattedText
}

func (row *FormattedRow) NewCol() *FormattedText {
	ft := FormattedText{}
	row.columns = append(row.columns, ft)
	return &row.columns[len(row.columns)-1]
}

func (row *FormattedRow) AddCol(col FormattedText) {
	row.columns = append(row.columns, col)
}

func (row *FormattedRow) GetCol(i int) *FormattedText {
	if i < len(row.columns) {
		return &row.columns[i]
	} else {
		return nil
	}
}

func NewFormattedRow(cols ...FormattedText) FormattedRow {
	return FormattedRow{columns: cols}
}

type Table struct {
	rows []FormattedRow
}

func (t *Table) NewRow() *FormattedRow {
	r := FormattedRow{}
	t.rows = append(t.rows, r)
	return &t.rows[len(t.rows)-1]
}

func (t *Table) AddRow(row FormattedRow) {
	t.rows = append(t.rows, row)
}

func (t *Table) Print(sep string, padded bool, color bool) {
	maxRowLen := 0
	for _, row := range t.rows {
		maxRowLen = max(maxRowLen, len(row.columns))
	}
	maxColLens := make([]int, maxRowLen)
	for _, row := range t.rows {
		for i := range maxRowLen {
			col := row.GetCol(i)
			var colLen int
			if col == nil {
				colLen = 0
			} else {
				colLen = len(col.rawText)
			}
			maxColLens[i] = max(maxColLens[i], colLen)
		}
	}
	for _, row := range t.rows {
		for i, col := range row.columns {
			if color {
				fmt.Print(col.richText)
			} else {
				fmt.Print(col.rawText)
			}
			if padded {
				padding := maxColLens[i] - len(col.rawText)
				fmt.Print(strings.Repeat(" ", padding))
			}
			if i < len(row.columns)-1 {
				fmt.Print(sep)
			} else {
				fmt.Println()
			}
		}
	}
}
