package renderer

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/paon/internal/geometry"
	"github.com/negrel/paon/internal/render"
)

var _ render.Renderer = block{}

type block struct{}

func makeBlock() render.Renderer {
	return block{}
}

func (b block) Layout(ctx *render.Context) {
	width := ctx.Object.Width()
	height := ctx.Object.Height()

	assert.GreaterOrEqual(width, 0, "widget width can't be a negative number")
	assert.GreaterOrEqual(height, 0, "widget height can't be a negative number")

	ctx.Layer.Resize(geometry.NewSize(width, height))
}

func (b block) Draw(ctx *render.Context) {
	panic("implement me")
}
