package main

import (
	"testing"
)

func Test_Damage_ReturnsDamage(t *testing.T) {
	want := 40
	r := Robot{damage: want}
	got := r.Damage()
	if got != want {
		t.Errorf("Damage() = %v, want %v", got, want)
	}

}

func Test_Scan_ReturnsClosestRobot(t *testing.T) {
	want := 0
	degree := 0
	resolution := 0
	r := Robot{}
	got := r.Scan(degree, resolution)
	if got != want {
		t.Errorf("Scan(%v, %v) = %v, want %v", degree, resolution, got, want)
	}
}
func Test_Canon_ReturnsWhetherMissileWasFired(t *testing.T) {
	want := true
	distance := 0
	degree := 0
	r := Robot{}
	got := r.Canon(degree, distance)
	if got != want {
		t.Errorf("Canon(%v, %v) = %v, want %v", degree, distance, got, want)
	}
}

type mockController struct {
	driveMock func(r *Robot, degree, speed int)
}

func (mc mockController) Drive(r *Robot, degree, speed int) {
	mc.driveMock(r, degree, speed)
}

// func Test_Drive_RecordsParameter(t *testing.T) {
// 	speed := 30
// 	degree := 45

// 	// c := mockController{
// 	// 	driveMock: func(r *Robot, degree, speed int) {
// 	// 		r.moveAngle = degree
// 	// 		r.moveSpeed = speed
// 	// 	},
// 	// }
// 	r := Robot{ /*controller: c*/ }
// 	r.Drive(degree, speed)

// 	t.Run("degree", func(t *testing.T) {
// 		want := degree
// 		got := int(math.Round(r.moveAngle))
// 		if got != want {
// 			t.Errorf("Drive(%v, %v) moveDegree = %v, want %v", degree, speed, got, want)
// 		}
// 	})
// 	t.Run("speed", func(t *testing.T) {
// 		want := speed
// 		got := int(math.Round(r.moveSpeed))
// 		if got != want {
// 			t.Errorf("Drive(%v, %v) moveSpeed = %v, want %v", degree, speed, got, want)
// 		}
// 	})
// }
