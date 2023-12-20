package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// func Test_Drive_RecordsParameter(t *testing.T) {
// 	speed := 30
// 	degree := 45

// 	// c := mockController{
// 	// 	driveMock: func(r *Robot, degree, speed int) {
// 	// 		r.moveDegree = degree
// 	// 		r.moveSpeed = speed
// 	// 	},
// 	// }
// 	c:=controller{}
// 	r := Robot{controller: c}
// 	r.Drive(degree, speed)

// 	t.Run("degree", func(t *testing.T) {
// 		want := degree
// 		got := r.moveDegree
// 		if got != want {
// 			t.Errorf("Drive(%v, %v) moveDegree = %v, want %v", degree, speed, got, want)
// 		}
// 	})
// 	t.Run("speed", func(t *testing.T) {
// 		want := speed
// 		got := r.moveSpeed
// 		if got != want {
// 			t.Errorf("Drive(%v, %v) moveSpeed = %v, want %v", degree, speed, got, want)
// 		}
// 	})
// }
// func Test_Controller_MoveRobot(t *testing.T) {
// 	c := controller{}
// 	r := Robot{}
// 	c.MoveRobot(r)
// }

func Test_Controller_RobotMovement_SetsNewPositionAfter(t *testing.T) {
	acceleration := 2.0
	maxSpeed := 15.0
	testCases := []struct {
		desc          string
		currentSpeedX int
		currentSpeedY int
		desiredSpeedX, desiredSpeedY int
		finalSpeedX   int
		finalSpeedY   int
		moveTime      int
	}{
		{
			desc:          "0 seconds",
			currentSpeedX: 10.0,
			currentSpeedY: 5.0,
			finalSpeedX:   10.0,
			finalSpeedY:   5.0,
			moveTime:      0.0,
		},
		{
			desc:          "1 seconds",
			currentSpeedX: 10.0,
			currentSpeedY: 5.0,
			desiredSpeedX: 100,
			desiredSpeedY: 100,
			finalSpeedX:   10+,
			finalSpeedY:   5.0 + 1*math.Sqrt2,
			moveTime:      1.0,
		},
		// {
		// 	desc:          "2 seconds",
		// 	currentSpeedX: 10.0,
		// 	currentSpeedY: 5.0,
		// 	finalSpeedX:   10.0 + 2*math.Sqrt2,
		// 	finalSpeedY:   5.0 + 2*math.Sqrt2,
		// 	moveTime:      2.0,
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := Robot{
				currentSpeedX: tC.currentSpeedX,
				currentSpeedY: tC.currentSpeedY,
				acceleration:  acceleration,
				maxSpeed:      maxSpeed,
			}
			r.UpdatePosition(tC.moveTime)
			opts := cmpopts.EquateApprox(0, 0.001)
			if !cmp.Equal(r.currentSpeedX, tC.finalSpeedX, opts) || !cmp.Equal(r.currentSpeedY, tC.finalSpeedY, opts) {
				t.Errorf("UpdatePosition(%d) expected current speed to be (%d, %d), got (%d, %d)",
					tC.moveTime, tC.finalSpeedX, tC.finalSpeedY, r.currentSpeedX, r.currentSpeedY)
			}
		})
	}
}

// func TestObjectMovement(t *testing.T) {
// 	robot := Robot{
// 		currentSpeedX: 10,
// 		currentSpeedY: 5,
// 		moveAngle:     45,
// 		acceleration:  2,
// 		maxSpeed:      15,
// 	}

// 	// Test the initial position
// 	robot.UpdatePosition(0)
// 	if robot.currentSpeedX != 10 || robot.currentSpeedY != 5 {
// 		t.Errorf("Expected current speeds to be (10, 5), got (%f, %f)", robot.currentSpeedX, robot.currentSpeedY)
// 	}

// 	// Test the position after 2 seconds
// 	robot.UpdatePosition(2)
// 	if robot.currentSpeedX != 14 || robot.currentSpeedY != 7 {
// 		t.Errorf("Expected current speeds to be (14, 7), got (%f, %f)", robot.currentSpeedX, robot.currentSpeedY)
// 	}

// 	// // Test the position after 5 seconds (should reach maxSpeed)
// 	// robot.UpdatePosition(5)
// 	// if robot.currentSpeedX != robot.maxSpeed*math.Cos(robot.angle*math.Pi/180) ||
// 	// 	robot.currentSpeedY != robot.maxSpeed*math.Sin(robot.angle*math.Pi/180) {
// 	// 	t.Errorf("Expected current speeds to be (%f, %f), got (%f, %f)",
// 	// 		robot.maxSpeed*math.Cos(robot.angle*math.Pi/180), robot.maxSpeed*math.Sin(robot.angle*math.Pi/180),
// 	// 		robot.currentSpeedX, robot.currentSpeedY)
// 	// }
// }

// func TestObjectMovement2(t *testing.T) {
// 	robot := Robot{
// 		currentSpeedX: 0,
// 		currentSpeedY: 0,
// 		moveAngle:     45,
// 		acceleration:  2,
// 		maxSpeed:      15,
// 	}

// 	// Test the initial position
// 	robot.UpdatePosition(0)
// 	if robot.currentSpeedX != 10 || robot.currentSpeedY != 10 {
// 		t.Errorf("Expected current speeds to be (10, 5), got (%f, %f)", robot.currentSpeedX, robot.currentSpeedY)
// 	}

// 	// Test the position after 2 seconds
// 	robot.UpdatePosition(2)
// 	if robot.currentSpeedX != 14 || robot.currentSpeedY != 14 {
// 		t.Errorf("Expected current speeds to be (14, 14), got (%f, %f)", robot.currentSpeedX, robot.currentSpeedY)
// 	}
// }
