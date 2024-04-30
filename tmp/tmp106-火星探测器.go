package main

type command int
type RoverDriver struct {
	commandc chan command
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{commandc: make(chan command)}
	go r.drive()
	return r
}
func (r *RoverDriver) drive() {

}
func main() {

}
