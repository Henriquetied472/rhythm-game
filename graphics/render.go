package graphics

import (
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/glfwcanvas"
)

func Init(w *glfwcanvas.Window, c *canvas.Canvas) {
	wnd = w
	cv = c
}

func Render() {
	cv.ClearRect(0, 0, 700, 700)
	cv.SetFillStyle("#FFFFFF")
	cv.FillRect(0, 0, Width, Height)

	Circle(((Width/3)-(Offset*2)-(15*2))*1+Offset, Height-(30+15), 15, true, false, "#55D6C2")
	Circle(((Width/3)-(Offset*2)-(15*2))*2+Offset, Height-(30+15), 15, true, false, "#5506C2")
	Circle(((Width/3)-(Offset*2)-(15*2))*3+Offset, Height-(30+15), 15, true, false, "#55D602")
}
