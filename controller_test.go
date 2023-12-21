package main

import (
	"testing"
)

type initialState struct {
	currentSpeedX, currentSpeedY int
	posX, posY                   int
	acceleration                 int // in units per second squared
	maxSpeed                     int // in units per second
}

type endState struct {
	speedX, speedY int
	posX, posY     int
}

func Test_Controller_RobotMovement_SetsNewPositionAfter(t *testing.T) {
	acceleration := 2
	maxSpeed := 15
	testCases := []struct {
		desc     string
		start    initialState
		end      endState
		moveTime int
	}{
		{
			desc:     "0 seconds",
			moveTime: 0,
			start: initialState{
				posX:          0,
				posY:          0,
				currentSpeedX: 10,
				currentSpeedY: 5,
			},
			end: endState{
				speedX: 10,
				speedY: 5,
				posX:   0,
				posY:   0,
			},
		},
		{
			desc:     "1 seconds",
			moveTime: 1,
			start: initialState{
				posX:          0,
				posY:          0,
				currentSpeedX: 10,
				currentSpeedY: 5,
			},
			end: endState{
				speedX: 12, // 10 + 1*acceleration
				speedY: 7,  // 5 + 1*acceleration
				posX:   12,
				posY:   7,
			},
		},
		{
			desc:     "3 seconds",
			moveTime: 3,
			start: initialState{
				posX:          0,
				posY:          0,
				currentSpeedX: 10,
				currentSpeedY: 5,
			},
			end: endState{
				speedX: maxSpeed,
				speedY: 11, // 5 + 3*acceleration
				posX:   15,
				posY:   11,
			},
		},
		{
			desc:     "10 seconds",
			moveTime: 10,
			start: initialState{
				posX:          0,
				posY:          0,
				currentSpeedX: 10,
				currentSpeedY: 5,
			},
			end: endState{
				speedX: maxSpeed,
				speedY: maxSpeed,
				posX:   15,
				posY:   11,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := Robot{
				currentSpeedX: tC.start.currentSpeedX,
				currentSpeedY: tC.start.currentSpeedY,
				acceleration:  acceleration,
				maxSpeed:      maxSpeed,
			}
			r.UpdatePosition(tC.moveTime)
			if r.currentSpeedX != tC.end.speedX || r.currentSpeedY != tC.end.speedY {
				t.Errorf("UpdatePosition(%d) expected current speed to be (%d, %d), got (%d, %d)",
					tC.moveTime, tC.end.speedX, tC.end.speedY, r.currentSpeedX, r.currentSpeedY)
			}
			if r.posX != tC.end.posX || r.posY != tC.end.posY {
				t.Errorf("UpdatePOsition(%d) expected current pos to be (%d, %d), got (%d, %d)",
					tC.moveTime, tC.end.posX, tC.end.posY, r.posX, r.posY)
			}
		})
	}
}
