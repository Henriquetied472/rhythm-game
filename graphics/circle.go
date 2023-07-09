package graphics

func Circle(x, y, radius float64, fill, stroke bool, color ...string) {
	cv.BeginPath()
	cv.Arc(x, y, radius, 0, 2*Pi, false)
	if fill {
		cv.SetFillStyle(color[0])
		cv.Fill()
	}
	if stroke {
		if len(color) == 1 {
			cv.SetStrokeStyle(color[0])
		} else {
			cv.SetStrokeStyle(color[1])
		}
		cv.Stroke()
	}
	cv.ClosePath()
}