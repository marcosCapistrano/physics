package main

import (
	"math"
	"physics/physics"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

const screenWidth = 800
const screenHeight = 600

const FPS uint32 = 60
const MILLISECS_PER_FRAME uint32 = 1000 / FPS

var myCar *car

var camera rl.Camera2D

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Physics")
	rl.SetTargetFPS(61)

	myCar = &car{
		Position:     physics.NewVec2(400, 300),
		Velocity:     physics.NewVec2(0, 0),
		Acceleration: physics.NewVec2(0, 0),
		Mass:         5,
		InvMass:      1.0 / 5.0,
		sumOfForces:  physics.NewVec2(0, 0),
		Rotation:     0,
	}

	camera.Target = rl.NewVector2(float32(myCar.Position.X+20), float32(myCar.Position.Y+20))
	camera.Offset = rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))
	camera.Rotation = 0.0
	camera.Zoom = 1.0

	shapes := make([]WorldShape, 100) // Create 100 shapes
	for i := range shapes {
		shapes[i] = generateRandomShape()
	}

	for !rl.WindowShouldClose() {
		camera.Target = rl.NewVector2(myCar.Position.X, myCar.Position.Y)
		update()
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.Black)

			rl.BeginMode2D(camera)
			myCar.Draw()
			DrawWorldShapes(shapes)
			rl.EndMode2D()
		}
		rl.EndDrawing()
	}
}

const (
	CAR_WIDTH  float32 = 40
	CAR_HEIGHT float32 = 20
)

func (c *car) Draw() {
	// Calculate the center point of rotation
	centerX := c.Position.X
	centerY := c.Position.Y

	// Convert rotation to radians
	rotation := float32(c.Rotation * math.Pi / 180.0)

	// Calculate the rectangle's corners relative to center
	halfWidth := CAR_WIDTH / 2
	halfHeight := CAR_HEIGHT / 2

	// Create a Rectangle for the car
	rec := rl.Rectangle{
		X:      centerX - halfWidth,
		Y:      centerY - halfHeight,
		Width:  CAR_WIDTH,
		Height: CAR_HEIGHT,
	}

	// Create Vector2 for rotation origin (center of rectangle)
	origin := rl.Vector2{
		X: halfWidth,
		Y: halfHeight,
	}

	// Draw the rotated rectangle
	rl.DrawRectanglePro(
		rec,        // Rectangle
		origin,     // Origin of rotation (relative to rectangle)
		c.Rotation, // Rotation in degrees
		rl.Red,     // Color
	)

	// Optional: Draw a line indicating the car's direction
	lineEndX := centerX + float32(math.Cos(float64(rotation))*20)
	lineEndY := centerY + float32(math.Sin(float64(rotation))*20)
	rl.DrawLine(
		int32(centerX),
		int32(centerY),
		int32(lineEndX),
		int32(lineEndY),
		rl.Green,
	)
}

func update() {
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		myCar.turnLeft()
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		myCar.turnRight()
	}

	// Handle acceleration
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		myCar.accelerate()
	}

	// Optional: Handle reverse
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		// Apply force in opposite direction of car's rotation
		myCar.AddForce(physics.NewVec2(
			float32(math.Cos(float64(myCar.Rotation*math.Pi/180.0))*-5),
			float32(math.Sin(float64(myCar.Rotation*math.Pi/180.0))*-5),
		))
	}

	myCar.AddForce(physics.NewVec2(0, 0.08*myCar.Mass*physics.PIXELS_PER_METER))
	myCar.AddForce(physics.NewFrictionForce(myCar.Velocity, 0.05))
	myCar.Integrate(rl.GetFrameTime())
}

type WorldShape struct {
	X, Y     float32
	Size     float32
	Color    rl.Color
	Type     int // 0: Circle, 1: Rectangle, 2: Triangle
	Rotation float32
}

func generateRandomShape() WorldShape {
	return WorldShape{
		X:    float32(rand.Intn(1001) - 500), // -500 to +500
		Y:    float32(rand.Intn(1001) - 500), // -500 to +500
		Size: float32(rand.Intn(41) + 10),    // 10 to 50
		Color: rl.Color{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		},
		Type:     rand.Intn(3),
		Rotation: float32(rand.Intn(360)),
	}
}

func DrawWorldShapes(shapes []WorldShape) {
	// Generate shapes

	// Draw background grid
	gridSize := float32(100)
	for x := float32(-500); x <= 500; x += gridSize {
		rl.DrawLine(int32(x), -500, int32(x), 500, rl.Gray)
	}
	for y := float32(-500); y <= 500; y += gridSize {
		rl.DrawLine(-500, int32(y), 500, int32(y), rl.Gray)
	}

	// Draw shapes
	for _, shape := range shapes {
		switch shape.Type {
		case 0: // Circle
			rl.DrawCircle(
				int32(shape.X),
				int32(shape.Y),
				shape.Size,
				shape.Color,
			)

		case 1: // Rectangle
			rec := rl.Rectangle{
				X:      shape.X - shape.Size/2,
				Y:      shape.Y - shape.Size/2,
				Width:  shape.Size,
				Height: shape.Size,
			}
			rl.DrawRectanglePro(
				rec,
				rl.Vector2{X: shape.Size / 2, Y: shape.Size / 2},
				shape.Rotation,
				shape.Color,
			)

		case 2: // Triangle
			angleRad := shape.Rotation * math.Pi / 180.0
			radius := shape.Size

			// Calculate triangle points
			p1 := rl.Vector2{
				X: shape.X + radius*float32(math.Cos(float64(angleRad))),
				Y: shape.Y + radius*float32(math.Sin(float64(angleRad))),
			}
			p2 := rl.Vector2{
				X: shape.X + radius*float32(math.Cos(float64(angleRad+2.0944))),
				Y: shape.Y + radius*float32(math.Sin(float64(angleRad+2.0944))),
			}
			p3 := rl.Vector2{
				X: shape.X + radius*float32(math.Cos(float64(angleRad+4.18879))),
				Y: shape.Y + radius*float32(math.Sin(float64(angleRad+4.18879))),
			}
			rl.DrawTriangle(p1, p2, p3, shape.Color)
		}
	}

	// Draw axes
	rl.DrawLine(-500, 0, 500, 0, rl.Black) // X axis
	rl.DrawLine(0, -500, 0, 500, rl.Black) // Y axis
}
