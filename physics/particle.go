package physics

type Particle struct {
	Position     Vec2
	Velocity     Vec2
	Acceleration Vec2

	sumOfForces Vec2

	Mass    float32
	InvMass float32
	Radius  float32
}

func NewParticle(x, y, mass float32, radius float32) *Particle {
	return &Particle{
		Position:     NewVec2(x, y),
		Velocity:     NewVec2(0, 0),
		Acceleration: NewVec2(0, 0),
		Mass:         mass,
		InvMass:      1 / mass,
		Radius:       radius,
		sumOfForces:  NewVec2(0, 0),
	}
}

func (p *Particle) AddForce(force Vec2) {
	p.sumOfForces.X += force.X
	p.sumOfForces.Y += force.Y
}

func (p *Particle) ClearForces() {
	p.sumOfForces.X = 0
	p.sumOfForces.Y = 0
}

func (p *Particle) Integrate(deltaTime float32) {
	/* acceleration */
	p.Acceleration.X = p.sumOfForces.X * PIXELS_PER_METER * p.InvMass
	p.Acceleration.Y = p.sumOfForces.Y * PIXELS_PER_METER * p.InvMass

	/* velocity */
	p.Velocity.X += p.Acceleration.X * deltaTime
	p.Velocity.Y += p.Acceleration.Y * deltaTime

	/* position */
	p.Position.X += p.Velocity.X * deltaTime
	p.Position.Y += p.Velocity.Y * deltaTime

	p.ClearForces()
}
