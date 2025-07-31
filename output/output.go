package output

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/viper"
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
			if ellipsis {
				toAdd = toAdd[:len(toAdd)-1] + "…"
			}
			clipped = true
		} else {
			toAdd = span.text
		}
		if withColor && (span.color != nil) {
			// color library adds a verbose reset code (eg `\xb[0;22;0;0;0m`) which seems to trip up some renderers, so
			// we add a more standard reset code to the end
			s += span.color.Sprint(toAdd) + "\x1b[0m"
		} else {
			s += toAdd
		}
		rawLen += len(toAdd)
		if clipped {
			break
		}
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
	rows   []Row
	header *Row
	footer string
}

func (t *Table) AddRow(row Row) {
	t.rows = append(t.rows, row)
}

func (t *Table) SetHeader(header Row) {
	t.header = &header
}

func (t *Table) SetFooter(footer string) {
	t.footer = footer
}

func (t *Table) Timestamp() {
	timestamp := time.Now().Format(time.RFC822)
	t.SetFooter(fmt.Sprintf("Last updated: %s", timestamp))
}

func (t *Table) Print(sep string, padded bool, withColor bool, emptyMsg string) {
	var s string
	var rows []Row
	if t.header != nil {
		rows = append(rows, *t.header)
	}
	rows = append(rows, t.rows...)
	maxRowLen := 0
	for _, row := range rows {
		maxRowLen = max(maxRowLen, len(row.cells))
	}
	maxCellLens := make([]int, maxRowLen)
	for _, row := range rows {
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
	for _, row := range rows {
		for i, cell := range row.cells {
			s += cell.Sprint(withColor, -1, false)
			if padded {
				padding := maxCellLens[i] - cell.rawLen
				s += strings.Repeat(" ", padding)
			}
			if i < len(row.cells)-1 {
				s += sep
			} else {
				s += "\n"
			}
		}
	}
	if len(t.rows) == 0 {
		s += fmt.Sprintln(emptyMsg)
	}
	if t.footer != "" {
		s += t.footer + "\n"
	}
	fmt.Print(s)
	//fmt.Printf("%q", s)
}

type Options struct {
	Color     bool
	Header    bool
	Timestamp bool
}

func OptionsFromConfig() Options {
	withColor := viper.GetBool("color")
	withHeader := viper.GetBool("header")
	withTimestamp := viper.GetBool("timestamp")
	color.NoColor = !withColor
	return Options{withColor, withHeader, withTimestamp}
}
