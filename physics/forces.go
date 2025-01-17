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
