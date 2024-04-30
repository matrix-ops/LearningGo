package main

type location_wugui struct {
	x, y int
}

func x_left(x *int) {
	*x -= 1
}
func x_right(x *int) {
	*x += 1
}
func y_left(y *int) {
	*y -= 1
}
func y_right(y *int) {
	*y += 1
}
func main() {
	wugui := location_wugui{
		x: 10,
		y: 12,
	}
	x_left(&wugui.x)
}
