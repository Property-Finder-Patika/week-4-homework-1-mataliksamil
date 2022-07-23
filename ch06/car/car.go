package car

import "fmt"

type Car struct {
	make, model string
	year        int
	speed       int
	distance    int
	convertible bool
}

//func New(.....) *Car{
//
//}

func (c *Car) Accelerate(speed int) {
	c.speed = speed
}

func (c *Car) Stop() {
	c.speed = 0
}

func (c *Car) Move(distance int) {
	c.distance += distance
}

func (c *Car) Print() {
	fmt.Printf("%#v\n", c)
}

func main() {
	myCar := Car{"Mercedes", "G", 2022, 0, 0, false}
	myCar.print()

	myCar.accelerate(100)
	myCar.move(200)
	myCar.print()

	myCar.accelerate(150)
	myCar.move(100)
	myCar.print()

	myCar.stop()
	myCar.print()

	println()

	yourCar := Car{"BMW", "i8", 2020, 200, 26200, true}
	yourCar.print()
}
