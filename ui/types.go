package ui

import (
	"fyne.io/fyne/v2"
	"github.com/c0nfsd/pixDraw/apptype"
	"github.com/c0nfsd/pixDraw/pxcanvas"
	"github.com/c0nfsd/pixDraw/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
