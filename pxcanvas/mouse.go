package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (pxCanvas *PXCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

func (pxCanvas *PXCanvas) MouseMoved(ev *desktop.MouseEvent) {
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()

	pxCanvas.mouseState.previousCoord = &ev.PointEvent
}

func (pxCanvas *PXCanvas) MouseIn(ev *desktop.MouseEvent) {}

func (pxCanvas *PXCanvas) MouseOut(ev *desktop.MouseEvent) {}
