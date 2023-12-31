package main

// From https://crobots.deepthought.it/html/manual.html
// 8. CROBOTS C Intrinsic Function Library

// The intrinsic function library provides machine level control and certain arithmetic functions. These functions do not consume any of the program code space or data stack, except for the three words for call/return sequences. No explicit linking is required to use any intrinsic function.

// scan (degree,resolution)

// The scan() function invokes the robot's scanner, at a specified degree and resolution. scan() returns 0 if no robots are within the scan range or a positive integer representing the range to the closest robot. Degree should be within the range 0-359, otherwise degree is forced into 0-359 by a modulo 360 operation, and made positive if necessary. Resolution controls the scanner's sensing resolution, up to +/- 10 degrees.

// Examples:
//    range = scan(45,0); /* scan 45, with no variance */
//    range = scan(365,10); /* scans the range from 355 to 15 */
// cannon (degree,range)

// The cannon() function fires a missile heading a specified range and direction. cannon() returns 1 (true) if a missile was fired, or 0 (false) if the cannon is reloading. Degree is forced into the range 0-359 as in scan(). Range can be 0-700, with greater ranges truncated to 700.

// Examples:
//    degree = 45;    /* set a direction to test */
//    if ((range=scan(degree,2)) > 0) /* see if a target is there */
//      cannon(degree,range);  /* fire a missile */
// drive (degree,speed)

// The drive() function activates the robot's drive mechanism, on a specified heading and speed. Degree is forced into the range 0-359 as in scan(). Speed is expressed as a percent, with 100 as maximum. A speed of 0 disengages the drive. Changes in direction can be negotiated at speeds of less than 50 percent.

// Examples:
//    drive(0,100);  /* head due east, at maximum speed */
//    drive(90,0);   /* stop motion */
// damage ()

// The damage() function returns the current amount of damage incurred. damage() takes no arguments, and returns the percent of damage, 0-99. (100 percent damage means the robot is completely disabled, thus no longer running!)

// Examples:
//    d = damage();       /* save current state */
//    ; ; ;               /* other instructions */
//    if (d != damage())  /* compare current state to prior state */
//    {
//      drive(90,100);    /* robot has been hit, start moving */
//      d = damage();     /* get current damage again */
//    }
// speed ()

// The speed() function returns the current speed of the robot. speed() takes no arguments, and returns the percent of speed, 0-100. Note that speed() may not always be the same as the last drive(), because of acceleration and deceleration.

// Examples:
//    drive(270,100);   /* start drive, due south */
//    ; ; ;             /* other instructions */
//    if (speed() == 0) /* check current speed */
//    {
//      drive(90,20);   /* ran into the south wall, or another robot */
//    }
// loc_x ()
// loc_y ()

// The loc_x() function returns the robot's current x axis location. loc_x() takes no arguments, and returns 0-999. The loc_y() function is similar to loc_x(), but returns the current y axis position.

// Examples:
//    drive (180,50);  /* start heading for west wall */
//    while (loc_x() > 20)
//      ;              /* do nothing until we are close */
//    drive (180,0);   /* stop drive */
// rand (limit)

// The rand() function returns a random number between 0 and limit, up to 32767.

// Examples:
//    degree = rand(360);     /* pick a random starting point */
//    range = scan(degree,0); /* and scan */
// sqrt (number)

// The sqrt() returns the square root of a number. Number is made positive, if necessary.

// Examples:
//    x = x1 - x2;     /* compute the classical distance formula */
//    y = y1 - y2;     /* between two points (x1,y1) (x2,y2) */
//    distance = sqrt((x*x) - (y*y));
// sin (degree)
// cos (degree)
// tan (degree)
// atan (ratio)

// These functions provide trigonometric values. sin(), cos(), and tan(), take a degree argument, 0-359, and returns the trigonometric value times 100,000. The scaling is necessary since the CROBOT cpu is an integer only machine, and trig values are between 0.0 and 1.0. atan() takes a ratio argument that has been scaled up by 100,000, and returns a degree value, between -90 and +90. The resulting calculation should not be scaled to the actual value until the final operation, as not to lose accuracy. See programming examples for usage.

type Robot struct {
	controller    Controller
	damage        int
	moveAngle     float64
	moveSpeed     float64
	currentSpeedX float64 // in units per second (x-axis)
	currentSpeedY float64 // in units per second (y-axis)
	// angle         float64 // in degrees
	acceleration float64 // in units per second squared
	maxSpeed     float64 // in units per second
	posX         float64
	posY         float64
}

// Damage returns the current percent of damage incurred, 0-99.
// 100 percent damage means the robot is completely disabled, thus no longer running
func (r Robot) Damage() int {
	return r.damage
}

// Scan invokes the robot's scanner, at a specified degree and resolution.
// Returns 0 if no robots are within the scan range or a positive integer representing the
// range to the closest robot.
// Degree should be within the range 0-359, otherwise degree is forced into 0-359 by a
// modulo 360 operation, and made positive if necessary.
// Resolution controls the scanner's sensing resolution, up to +/- 10 degrees.
func (r Robot) Scan(degree, resolution int) int {
	return 0
}

// Canon fires a missile heading a specified direction and exploding at a given distance.
// Returns true if a missile was fired, or false if the cannon is reloading.
// Degree is forced into the range 0-359 as in scan().
// Distance can be 0-700, with greater ranges truncated to 700.
func (r Robot) Canon(degree, distance int) bool {
	return true
}

// Drive activates the robot's drive mechanism, on a specified heading and speed.
// Degree is forced into the range 0-359.
// Speed is expressed as a percent, with 100 as maximum. A speed of 0 disengages the drive.
// Changes in direction can be negotiated at speeds of less than 50 percent.
func (r *Robot) Drive(degree, speed int) {
	// r.controller.Drive(r, degree, speed)
	r.moveAngle = float64(degree)
	r.moveSpeed = float64(speed)
}
