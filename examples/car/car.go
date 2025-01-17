package main

import (
	"math"
	"physics/physics"
)

type car struct {
	Position     physics.Vec2
	Velocity     physics.Vec2
	Acceleration physics.Vec2
	sumOfForces  physics.Vec2

	Mass    float32
	InvMass float32

	Rotation float32
}

func (p *car) AddForce(force physics.Vec2) {
	p.sumOfForces.X += force.X
	p.sumOfForces.Y += force.Y
}

func (p *car) ClearForces() {
	p.sumOfForces.X = 0
	p.sumOfForces.Y = 0
}

func (p *car) Integrate(deltaTime float32) {
	/* acceleration */
	p.Acceleration.X = p.sumOfForces.X * physics.PIXELS_PER_METER * p.InvMass
	p.Acceleration.Y = p.sumOfForces.Y * physics.PIXELS_PER_METER * p.InvMass

	/* velocity */
	p.Velocity.X += p.Acceleration.X * deltaTime
	p.Velocity.Y += p.Acceleration.Y * deltaTime

	/* position */
	p.Position.X += p.Velocity.X * deltaTime
	p.Position.Y += p.Velocity.Y * deltaTime

	p.ClearForces()
}

func (c *car) turnLeft() {
	c.Rotation -= 1.0
	if c.Rotation < 0 {
		c.Rotation += 360
	}
}

func (c *car) turnRight() {
	c.Rotation += 1.0
	if c.Rotation > 360 {
		c.Rotation -= 360
	}
}

func (c *car) accelerate() {
	angleRad := c.Rotation * math.Pi / 180.0
	c.AddForce(physics.NewVec2(float32(math.Cos(float64(angleRad))*50), float32(math.Sin(float64(angleRad))*50)))
}
