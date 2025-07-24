package output

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type span struct {
	text  string
	color *color.Color
}

type Cell struct {
	spans  []span
	rawLen int
}

func NewCell(text string, color *color.Color) Cell {
	cell := Cell{}
	cell.AddText(text, color)
	return cell
}

func (cell *Cell) AddText(text string, color *color.Color) {
	cell.spans = append(cell.spans, span{text, color})
	cell.rawLen += len(text)
}

func (cell *Cell) Sprint(withColor bool, maxLen int, ellipsis bool) string {
	var s string
	clipped := false
	rawLen := 0
	for _, span := range cell.spans {
		var toAdd string
		if (maxLen >= 0) && (rawLen+len(span.text) > maxLen) {
			toAdd = span.text[:(maxLen - rawLen)]
			clipped = true
		} else {
			toAdd = span.text
		}
		if withColor && (span.color != nil) {
			s += span.color.Sprint(toAdd)
		} else {
			s += toAdd
		}
		rawLen += len(toAdd)
		if clipped {
			break
		}
	}
	if clipped && ellipsis {
		s = s[:len(s)-1] + "…"
	}
	return s
}

// Row is a row of `FormattedText` objects (each of which represents a column).
type Row struct {
	cells []Cell
}

func (row *Row) AddCell(cell Cell) {
	row.cells = append(row.cells, cell)
}

func (row *Row) GetCell(i int) *Cell {
	if i < len(row.cells) {
		return &row.cells[i]
	} else {
		return nil
	}
}

func NewRow(cells ...Cell) Row {
	return Row{cells: cells}
}

type Table struct {
	rows []Row
}

func (t *Table) NewRow() *Row {
	r := Row{}
	t.rows = append(t.rows, r)
	return &t.rows[len(t.rows)-1]
}

func (t *Table) AddRow(row Row) {
	t.rows = append(t.rows, row)
}

func (t *Table) Print(sep string, padded bool, withColor bool) {
	maxRowLen := 0
	for _, row := range t.rows {
		maxRowLen = max(maxRowLen, len(row.cells))
	}
	maxCellLens := make([]int, maxRowLen)
	for _, row := range t.rows {
		for i := range maxRowLen {
			cell := row.GetCell(i)
			var cellLen int
			if cell == nil {
				cellLen = 0
			} else {
				cellLen = cell.rawLen
			}
			maxCellLens[i] = max(maxCellLens[i], cellLen)
		}
	}
	for _, row := range t.rows {
		for i, cell := range row.cells {
			fmt.Print(cell.Sprint(withColor, -1, false))
			if padded {
				padding := maxCellLens[i] - cell.rawLen
				fmt.Print(strings.Repeat(" ", padding))
			}
			if i < len(row.cells)-1 {
				fmt.Print(sep)
			} else {
				fmt.Println()
			}
		}
	}
}
