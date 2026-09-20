package transport

import "fmt"

type Transport interface {
	Deliver(cargo, destination string) string
}

type Truck struct{}

func (Truck) Deliver(cargo, destination string) string {
	return fmt.Sprintf("Truck delivers %s to %s warehouse", cargo, destination)
}

type Ship struct{}

func (Ship) Deliver(cargo, destination string) string {
	return fmt.Sprintf("Ship delivers %s to %s port", cargo, destination)
}
