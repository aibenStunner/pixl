package pxcanvas

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"thegadri.io/pixl/apptype"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type PXCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer    *PxCanvasRenderer
	PixelData   image.Image
	mouseState  PxCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (pxCanvas *PXCanvas) Bounds() image.Rectangle {
	x0 := int(pxCanvas.CanvasOffset.X)
	y0 := int(pxCanvas.CanvasOffset.Y)
	x1 := int(pxCanvas.PxCols*pxCanvas.PxSize + int(pxCanvas.CanvasOffset.X))
	y1 := int(pxCanvas.PxRow*pxCanvas.PxSize + int(pxCanvas.CanvasOffset.Y))

	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) &&
		pos.X < float32(bounds.Max.X) &&
		pos.Y >= float32(bounds.Min.Y) &&
		pos.Y < float32(bounds.Max.Y) {
		return true
	}
	return false
}

func NewBlankImage(cols int, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))
	for y := 0; y < 0; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, c)
		}
	}

	return img
}

func NewPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PXCanvas {
	pxCanvas := &PXCanvas{
		PxCanvasConfig: config,
		appState:       state,
	}

	pxCanvas.PixelData = NewBlankImage(pxCanvas.PxCols, pxCanvas.PxRow, color.NRGBA{128, 128, 128, 128})
	pxCanvas.ExtendBaseWidget(pxCanvas)
	return pxCanvas
}

func (pxCanvas *PXCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &PxCanvasRenderer{
		pxCanvas:     pxCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}

	pxCanvas.renderer = renderer

	return renderer
}

func (pxCanvas *PXCanvas) TryPan(previousCoord *fyne.PointEvent, ev *desktop.MouseEvent) {
	if previousCoord != nil && ev.Button == desktop.MouseButtonTertiary {
		pxCanvas.Pan(*previousCoord, ev.PointEvent)
	}
}

// Brushable interface
func (pxCanvas *PXCanvas) SetColor(c color.Color, x, y int) {
	if nrgba, ok := pxCanvas.PixelData.(*image.NRGBA); ok {
		nrgba.Set(x, y, c)
	}

	if rgba, ok := pxCanvas.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, c)
	}

	pxCanvas.Refresh()
}

func (pxCanvas *PXCanvas) MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int) {
	bounds := pxCanvas.Bounds()
	if !InBounds(ev.Position, bounds) {
		return nil, nil
	}

	pxSize := float32(pxCanvas.PxSize)
	xOffset := pxCanvas.CanvasOffset.X
	yOffset := pxCanvas.CanvasOffset.Y

	x := int((ev.Position.X - xOffset) / pxSize)
	y := int((ev.Position.Y - yOffset) / pxSize)

	return &x, &y
}

func (pxCanvas *PXCanvas) LoadImage(img image.Image) {
	dimensions := img.Bounds()

	pxCanvas.PxCanvasConfig.PxCols = dimensions.Dx()
	pxCanvas.PxCanvasConfig.PxRow = dimensions.Dy()

	pxCanvas.PixelData = img
	pxCanvas.reloadImage = true
	pxCanvas.Refresh()
}

func (pxCanvas *PXCanvas) NewDrawing(cols, rows int) {
	pxCanvas.appState.SetFilePath("")
	pxCanvas.PxCols = cols
	pxCanvas.PxRow = rows
	pixelData := NewBlankImage(cols, rows, color.NRGBA{128, 128, 128, 255})
	pxCanvas.LoadImage(pixelData)
}
