package physics

import "math"

type Vec2 struct {
	X, Y float32
}

func NewVec2(x, y float32) Vec2 {
	return Vec2{
		X: x,
		Y: y,
	}
}

func (v *Vec2) Magnitude() float32 {
	mag := math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2))

	return float32(mag)
}

func (v *Vec2) MagnitudeSquared() float32 {
	magSqr := math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2)

	return float32(magSqr)
}

func (v *Vec2) Normalized() Vec2 {
	mag := v.Magnitude()
	normalized := *v

	normalized.X /= mag
	normalized.Y /= mag

	return normalized
}
