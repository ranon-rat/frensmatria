package main

import (
	"fmt"
	"time"
)

func General(a int, id string) {
	for {
		fmt.Println("sup", a, id)
		time.Sleep(time.Second)
	}
}
func Routine(a int) {
	go General(a, "nodo a")
	go General(a, "nodo b")

}
func main() {
	Routine(1)
	Routine(2)
	select {}

}
