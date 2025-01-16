// main.go
package main

import (
	"log"
	"physics/physics"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Physics", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		log.Fatal(err)
	}
	defer renderer.Destroy()

	particle := physics.NewParticle(400, 300, 1)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false

			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym
				if keyCode == sdl.K_ESCAPE {
					running = false
				}
			}
		}

		renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)
		renderer.Clear()

		renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)
		renderer.DrawPoint(int32(particle.Position.X), int32(particle.Position.Y))

		renderer.Present()
		sdl.Delay(33)
	}
}
