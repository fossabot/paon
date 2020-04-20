package widgets

import (
	"image"
	"log"

	"github.com/negrel/ginger/v2/render"
)

var _ Layout = &_column{}
var _ Widget = &_column{}

// _column is a layout that arrange widget vertically.
type _column struct {
	*CoreLayout
}

// Column return a layout that arrange widget vertically.
func Column(children []Widget) Layout {
	col := &_column{
		CoreLayout: NewCoreLayout("column", children),
	}

	col.Rendering = col.rendering

	return col
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Widget interface

func (c *_column) renderChilds(bounds image.Rectangle) ([]*render.Frame, image.Point) {
	childCount := len(c.Children)
	childrenFrames := make([]*render.Frame, childCount)
	size := image.Pt(0, 0)

	for i := 0; i < childCount; i++ {
		childFrame := c.Children[i].Render(bounds)
		childrenFrames[i] = childFrame

		if childWidth := childFrame.Patch.Width(); childWidth > size.X {
			size.X = childWidth
		}

		// updating bounds for next children
		bounds.Min.Y += childFrame.Patch.Height()
		// updating total size.Y
		size.Y += childFrame.Patch.Height()
	}

	return childrenFrames, size
}

// Rendering implements Widget interface.
func (c *_column) rendering(bounds image.Rectangle) *render.Frame {
	// children bounds are relative
	childBounds := image.Rectangle{
		Min: image.Pt(0, 0),
		Max: bounds.Max.Sub(bounds.Min),
	}

	childrenFrames, size := c.renderChilds(childBounds)

	frame := render.NewFrame(bounds.Min, size.X, size.Y)

	for i, childFrame := range childrenFrames {
		err := frame.Add(childFrame)

		if err != nil {
			log.Print("COLUMN: ", err)
			log.Printf("COLUMN FRAME: %+v %+v %+v", frame.Position, frame.Patch.Width(), frame.Patch.Height())
			log.Fatalf("CHILD n°%v FRAME: %+v %+v %+v", i, childFrame, childFrame.Patch.Width(), childFrame.Patch.Height())
		}
	}

	return frame
}