package graphics

import (
	"math"

	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/glfwcanvas"
)

const (
	Width = 500
	Height = 700
	Pi = math.Pi
	Offset = 10
)

var (
	wnd *glfwcanvas.Window
	cv *canvas.Canvas
)