package main

type controller interface {
	driver
}

type driver interface {
	Drive(r *Robot, degree, speed int)
}
