
package app

import (
	"fmt"

	"logistics-go/gui"
	"logistics-go/logistics"
)


type DeliveryApplication struct {
	factory   gui.GUIFactory
	logistics logistics.Logistics
}


func New(factory gui.GUIFactory, l logistics.Logistics) *DeliveryApplication {
	return &DeliveryApplication{factory: factory, logistics: l}
}

func (a *DeliveryApplication) Run(cargo, destination string) {
	button := a.factory.CreateButton()
	checkbox := a.factory.CreateCheckbox()

	fmt.Println(button.Paint())
	fmt.Println(checkbox.Paint())
	fmt.Println(logistics.PlanDelivery(a.logistics, cargo, destination))
}
