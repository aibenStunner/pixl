package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"thegadri.io/pixl/pxcanvas/brush"
)

func (pxCanvas *PXCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

func (pxCanvas *PXCanvas) MouseMoved(ev *desktop.MouseEvent) {
	if x, y := pxCanvas.MouseToCanvasXY(ev); x != nil && y != nil {
		brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
	}

	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()

	pxCanvas.mouseState.previousCoord = &ev.PointEvent
}

func (pxCanvas *PXCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (pxCanvas *PXCanvas) MouseOut()                      {}

func (pxCanvas *PXCanvas) MouseUp(ev *desktop.MouseEvent) {}

func (pxCanvas *PXCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
}
