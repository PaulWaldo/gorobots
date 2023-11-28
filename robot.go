package main

type Robot struct {
	damage int
}

func (r Robot) GetDamage() int {
	return r.damage
}
