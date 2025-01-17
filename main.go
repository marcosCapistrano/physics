// main.go
package main

import (
	"fmt"
	"physics/physics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth = 800
const screenHeight = 600

var particle physics.Particle
var fps int32

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Physics")
	rl.SetTargetFPS(61)

	particle = physics.NewParticle(screenWidth/2, screenHeight/2, 5, 5)
	particle.Velocity.X = 4 * PIXELS_PER_METER

	for !rl.WindowShouldClose() {
		update()
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.Black)
			rl.DrawText(fmt.Sprintf("FPS: %d", fps), 10, 0, 16, rl.White)
			rl.DrawCircle(int32(particle.Position.X), int32(particle.Position.Y), particle.Radius, rl.White)
		}
		rl.EndDrawing()
	}
}

func update() {
	deltaTime := rl.GetFrameTime()
	fps = int32(1 / deltaTime)

	particle.Acceleration = physics.NewVec2(0, 90.8*PIXELS_PER_METER)
	particle.Velocity.X += particle.Acceleration.X * deltaTime
	particle.Velocity.Y += particle.Acceleration.Y * deltaTime
	particle.Position.X += particle.Velocity.X * deltaTime
	particle.Position.Y += particle.Velocity.Y * deltaTime

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
