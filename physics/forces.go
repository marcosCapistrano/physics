package physics

func NewDragForce(velocity Vec2, k float32) Vec2 {
	dragForce := NewVec2(0, 0)

	if velocity.MagnitudeSquared() > 0 {
		dragDirection := velocity.Normalized()
		dragDirection.X *= -1
		dragDirection.Y *= -1

		dragMag := k * velocity.MagnitudeSquared()

		dragForce = dragDirection
		dragForce.X *= dragMag
		dragForce.Y *= dragMag
	}

	return dragForce
}

func NewFrictionForce(velocity Vec2, k float32) Vec2 {
	fricDirection := velocity

	fricDirection.X *= -1
	fricDirection.Y *= -1

	fricForce := fricDirection
	fricForce.X *= k
	fricForce.Y *= k

	return fricForce
}
