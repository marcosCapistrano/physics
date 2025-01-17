package graphics

import "github.com/veandco/go-sdl2/sdl"

func DrawPoint(renderer *sdl.Renderer, x, y float32) {
	renderer.DrawPoint(int32(x), int32(y))
}

func DrawCircle(renderer *sdl.Renderer, centerX, centerY, radius float32) {
	x := radius
	var y float32 = 0
	var radiusError float32 = 1 - x

	for x >= y {
		// Draw 8 symmetrical points for each step
		DrawPoint(renderer, centerX+x, centerY-y)
		DrawPoint(renderer, centerX+y, centerY-x)
		DrawPoint(renderer, centerX-y, centerY-x)
		DrawPoint(renderer, centerX-x, centerY-y)
		DrawPoint(renderer, centerX-x, centerY+y)
		DrawPoint(renderer, centerX-y, centerY+x)
		DrawPoint(renderer, centerX+y, centerY+x)
		DrawPoint(renderer, centerX+x, centerY+y)

		y++

		if radiusError < 0 {
			radiusError += 2*y + 1
		} else {
			x--
			radiusError += 2 * (y - x + 1)
		}
	}
}
