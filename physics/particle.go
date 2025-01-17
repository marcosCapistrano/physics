package physics

type Particle struct {
	Position     Vec2
	Velocity     Vec2
	Acceleration Vec2

	Mass   float32
	Radius float32
}

func NewParticle(x, y, mass float32, radius float32) Particle {
	return Particle{
		Position:     NewVec2(x, y),
		Velocity:     NewVec2(0, 0),
		Acceleration: NewVec2(0, 0),
		Mass:         mass,
		Radius:       radius,
	}
}
