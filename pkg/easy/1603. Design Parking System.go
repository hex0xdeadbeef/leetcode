package easy

// https://leetcode.com/problems/design-parking-system/

type CarType int

const (
	_ = iota
	small
	medium
	big
)

type ParkingSystem struct {
	places map[CarType]int
}

func Constructor(s, m, b int) ParkingSystem {
	return ParkingSystem{places: map[CarType]int{small: s, medium: m, big: b}}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	var (
		key = CarType(carType)
	)
	ps.places[key]--

	return ps.places[key] >= 0
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
