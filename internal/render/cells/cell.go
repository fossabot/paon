package cells

import (
	"github.com/gdamore/tcell"
)

// Cell define a terminal screen cell.
type Cell interface {
	Style() tcell.Style
	Content() rune
}
