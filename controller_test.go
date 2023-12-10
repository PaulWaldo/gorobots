package main

import "testing"

func Test_Drive_RecordsParameter(t *testing.T) {
	speed := 30
	degree := 45

	// c := mockController{
	// 	driveMock: func(r *Robot, degree, speed int) {
	// 		r.moveDegree = degree
	// 		r.moveSpeed = speed
	// 	},
	// }
	c:=controller{}
	r := Robot{controller: c}
	r.Drive(degree, speed)

	t.Run("degree", func(t *testing.T) {
		want := degree
		got := r.moveDegree
		if got != want {
			t.Errorf("Drive(%v, %v) moveDegree = %v, want %v", degree, speed, got, want)
		}
	})
	t.Run("speed", func(t *testing.T) {
		want := speed
		got := r.moveSpeed
		if got != want {
			t.Errorf("Drive(%v, %v) moveSpeed = %v, want %v", degree, speed, got, want)
		}
	})
}
// func Test_Controller_MoveRobot(t *testing.T) {
// 	c := controller{}
// 	r := Robot{}
// 	c.MoveRobot(r)
// }
