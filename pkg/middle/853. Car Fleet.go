package middle

import (
	"fmt"
	"slices"
)

func CarFleet(target int, position []int, speed []int) int {
	type (
		pair struct {
			pos      int
			velocity int

			stepsLeft int
		}
	)

	var (
		cars = make([]pair, len(position))

		stack = make([]pair, 0, len(position))
	)

	for i := 0; i < len(position); i++ {
		newCar := pair{pos: position[i], velocity: speed[i]}

		if remainder := (target - newCar.pos) % newCar.velocity; remainder == 1 {
			newCar.stepsLeft = (target-newCar.pos)/newCar.velocity + 1
		} else {
			newCar.stepsLeft = (target - newCar.pos) / newCar.velocity
		}

		cars[i] = newCar
	}

	slices.SortFunc(cars, func(a, b pair) int {
		return a.pos - b.pos
	})

	for _, v := range cars {
		fmt.Printf("%+v\n", v)
	}

	for i := len(cars) - 1; i > -1; i-- {
		switch {
		case len(stack) == 0:
			stack = append(stack, cars[i])
		
		}

	}

	return -1

}
