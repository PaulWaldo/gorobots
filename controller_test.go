package main

import (
	"testing"
)

func Test_Controller_RobotMovement_SetsNewPositionAfter(t *testing.T) {
	acceleration := 2
	maxSpeed := 15
	testCases := []struct {
		desc                         string
		currentSpeedX                int
		currentSpeedY                int
		desiredSpeedX, desiredSpeedY int
		finalSpeedX                  int
		finalSpeedY                  int
		moveTime                     int
	}{
		{
			desc:          "0 seconds",
			currentSpeedX: 10,
			currentSpeedY: 5,
			finalSpeedX:   10,
			finalSpeedY:   5,
			moveTime:      0,
		},
		{
			desc:          "1 seconds",
			currentSpeedX: 10,
			currentSpeedY: 5,
			desiredSpeedX: 100,
			desiredSpeedY: 100,
			finalSpeedX:   12, // 10 + 1*acceleration
			finalSpeedY:   7,  // 5 + 1*acceleration
			moveTime:      1,
		},
		{
			desc:          "3 seconds",
			currentSpeedX: 10,
			currentSpeedY: 5,
			desiredSpeedX: 100,
			desiredSpeedY: 100,
			finalSpeedX:   15, // 10 + 3*acceleration, but maxxed
			finalSpeedY:   11, // 5 + 3*acceleration
			moveTime:      3,
		},
		{
			desc:          "10 seconds",
			currentSpeedX: 10,
			currentSpeedY: 5,
			desiredSpeedX: 100,
			desiredSpeedY: 100,
			finalSpeedX:   15, // maxed
			finalSpeedY:   15, // maxed
			moveTime:      10,
		},
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
			if r.currentSpeedX != tC.finalSpeedX || r.currentSpeedY != tC.finalSpeedY {
				t.Errorf("UpdatePosition(%d) expected current speed to be (%d, %d), got (%d, %d)",
					tC.moveTime, tC.finalSpeedX, tC.finalSpeedY, r.currentSpeedX, r.currentSpeedY)
			}
		})
	}
}
