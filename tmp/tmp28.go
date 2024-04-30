package main

// func goRoutineTest28(v int, c chan int) {
// 	c <- v
// }

func main() {
	c := make(chan int)
	<-c
}
