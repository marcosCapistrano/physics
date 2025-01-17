package main

import (
	"physics/graphics"
	"physics/physics"

	"github.com/veandco/go-sdl2/sdl"
)

type application struct {
	window            *sdl.Window
	renderer          *sdl.Renderer
	isRunning         bool
	particle          physics.Particle
	timePreviousFrame uint64
}

func (app *application) start() error {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("Physics", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		return err
	}

	app.window = window
	app.renderer = renderer
	app.isRunning = true
	app.particle = physics.NewParticle(400, 300, 5)
	app.timePreviousFrame = 0

	return nil
}

func (app *application) run() {
	for app.isRunning {
		app.handleInput()
		app.update()
		app.render()
		sdl.Delay(33)
	}
}

func (app *application) handleInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			app.isRunning = false

		case *sdl.KeyboardEvent:
			keyCode := t.Keysym.Sym
			if keyCode == sdl.K_ESCAPE {
				app.isRunning = false
			}
		}
	}
}

func (app *application) update() {
	timeToWait := int32(MILLISECS_PER_FRAME) - int32((sdl.GetTicks64() - app.timePreviousFrame))
	if timeToWait > 0 {
		sdl.Delay(uint32(timeToWait))
	}

	// deltaTime in seconds
	deltaTime := float32(sdl.GetTicks64()-app.timePreviousFrame) / 1000

	app.timePreviousFrame = sdl.GetTicks64()

	app.particle.Velocity = physics.NewVec2(100*deltaTime, 0.0)
	app.particle.Position.X += app.particle.Velocity.X
	app.particle.Position.Y += app.particle.Velocity.Y
}

func (app *application) render() {
	app.renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)
	app.renderer.Clear()

	app.renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)
	graphics.DrawCircle(app.renderer, app.particle.Position.X, app.particle.Position.Y, 5)

	app.renderer.Present()
}

func (app *application) quit() {
	app.renderer.Destroy()
	app.window.Destroy()
	sdl.Quit()
}
