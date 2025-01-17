// main.go
package main

import (
	"fmt"
	"physics/physics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth = 800
const screenHeight = 600

var particle *physics.Particle
var fps int32

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Physics")
	rl.SetTargetFPS(61)

	particle = physics.NewParticle(screenWidth/2, 0, 5, 5)

	for !rl.WindowShouldClose() {
		update()
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.Black)
			rl.DrawText(fmt.Sprintf("FPS: %d", fps), 10, 0, 16, rl.White)
			rl.DrawRectangle(0, 400, 800, 200, rl.Blue)
			rl.DrawCircle(int32(particle.Position.X), int32(particle.Position.Y), particle.Radius, rl.White)
		}
		rl.EndDrawing()
	}
}

func update() {
	deltaTime := rl.GetFrameTime()
	fps = int32(1 / deltaTime)

	weightForce := physics.NewVec2(0, 9.8*particle.Mass*PIXELS_PER_METER)

	particle.AddForce(weightForce)
	if particle.Position.Y > 400 {
		particle.AddForce(physics.NewDragForce(particle.Velocity, 0.25))
	} else {
		particle.AddForce(physics.NewDragForce(particle.Velocity, 0.01))
	}
	particle.Integrate(deltaTime)

	x := particle.Position.X
	y := particle.Position.Y
	radius := particle.Radius

	if x+radius > screenWidth {
		particle.Position.X = screenWidth - radius
		particle.Velocity.X *= -1
	} else if x-radius < 0 {
		particle.Position.X = radius
		particle.Velocity.X *= -1
	}

	if y+radius > screenHeight {
		particle.Position.Y = screenHeight - radius
		particle.Velocity.Y *= -1
	} else if y-radius < 0 {
		particle.Position.Y = radius
		particle.Velocity.Y *= -1
	}
}
