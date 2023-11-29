package main

import "testing"

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
func Test_Drive_RecordsParameters(t *testing.T) {
	speed := 0
	degree := 0
	r := Robot{}
	r.Drive(degree, speed)
}
