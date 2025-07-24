package output

import (
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestCellBasic(t *testing.T) {
	cell := Cell{}
	cell.AddText("Hello", nil)
	assert.Equal(t, "Hello", cell.Sprint(false, -1, false))
	assert.Equal(t, "Hel", cell.Sprint(false, 3, false))
	assert.Equal(t, "He…", cell.Sprint(false, 3, true))
	cell.AddText(" World", color.RGB(20, 20, 20))
	assert.Equal(t, "Hello World", cell.Sprint(false, -1, false))
	assert.Equal(t, "Hello Wor", cell.Sprint(false, 9, false))
	assert.Equal(t, "Hello Wo…", cell.Sprint(false, 9, true))
}
