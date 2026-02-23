package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	go task("A")
	go task("B")

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second)

}
