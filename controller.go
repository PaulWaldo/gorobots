package main

type Controller interface {
	// driver
}

type control struct{}

// func (c controller) Drive(r *Robot, degree, speed int){

// }

// type driver interface {
// 	Drive(r *Robot, degree, speed int)
// }

// UpdatePosition updates the position of the object based on its current speed, angle, acceleration, and maximum speed.
func (r *Robot) UpdatePosition(timeInSeconds int) {
	// // Convert angle to radians
	// radians := r.moveAngle * (math.Pi / 180)

	// Update speeds based on acceleration
	r.currentSpeedX += r.acceleration * timeInSeconds
	r.currentSpeedY += r.acceleration * timeInSeconds

	// Limit speeds to the maximum speed
	if r.currentSpeedX > r.maxSpeed {
		r.currentSpeedX = r.maxSpeed
	}
	if r.currentSpeedY > r.maxSpeed {
		r.currentSpeedY = r.maxSpeed
	}

	// speedMagnitude := math.Sqrt(math.Pow(r.currentSpeedX, 2) + math.Pow(r.currentSpeedY, 2))
	// if speedMagnitude > r.maxSpeed {
	// 	scaleFactor := r.maxSpeed / speedMagnitude
	// 	r.currentSpeedX *= scaleFactor
	// 	r.currentSpeedY *= scaleFactor
	// }

	// Update position based on speeds
	r.posX = r.currentSpeedX * timeInSeconds
	r.posY = r.currentSpeedY * timeInSeconds
}
