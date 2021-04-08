package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func Initialize_Display() {
	if error := sdl.Init(sdl.INIT_EVERYTHING); error != nil {
		panic(error)
	}
	defer sdl.Quit()

	window, error := sdl.CreateWindow("go-boy", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		160, 144, sdl.WINDOW_SHOWN)
	if error != nil {
		panic(error)
	}
	defer window.Destroy()

	surface, error := window.GetSurface()
	if error != nil {
		panic(error)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 160, 144}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
			}
		}
	}
}
