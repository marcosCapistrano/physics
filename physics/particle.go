package physics

type Particle struct {
	Position     Vec2
	velocity     Vec2
	acceleration Vec2

	mass float32
}

func NewParticle(x, y, mass float32) Particle {
	return Particle{
		Position:     NewVec2(x, y),
		velocity:     NewVec2(0, 0),
		acceleration: NewVec2(0, 0),
		mass:         mass,
	}
}
