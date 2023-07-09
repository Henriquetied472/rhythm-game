package main

import (
	"log"

	"github.com/Henriquetie472/rhythm-game/graphics"
	"github.com/tfriedel6/canvas/glfwcanvas"
)

func main() {
	wnd, cv, err := glfwcanvas.CreateWindow(graphics.Width, graphics.Height, "Rhythm Game")
	if err != nil {
		log.Fatal(err)
	}
	defer wnd.Close()

	graphics.Init(wnd, cv)

	wnd.MainLoop(func() {
		graphics.Render()
	})

}