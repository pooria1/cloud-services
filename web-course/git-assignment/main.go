package main

import (
	"car/car"
	"fmt"
)

func main() {
	bmw := car.NewCar(car.BMW, "730li", car.Red)

	for i := 0; i < 10; i++ {
		bmw.SpeedUp()
	}

	fmt.Printf("Owr brand new car has %d speed\n", bmw.GetSpeed())
}
