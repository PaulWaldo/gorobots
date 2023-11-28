package main

import "testing"

func Test_GetDamage_GivesProperDamage(t *testing.T) {
	want := 40
	r := Robot{damage: want}
	got := r.GetDamage()
	if got != want {
		t.Errorf("GetDamage() = %v, want %v", got, want)
	}

}
