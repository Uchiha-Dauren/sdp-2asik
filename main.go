package main

import (
	"fmt"
	"os"
	"strings"

	"logistics-go/app"
	"logistics-go/gui"
	"logistics-go/logistics"
)

const (
	usage    = "Usage: logistics-go <ROAD|SEA> <WINDOWS|MACOS>"
	cargo    = "laboratory equipment"
	destAddr = "Aktau"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(usage)
		os.Exit(1)
	}

	modeArg := strings.ToUpper(os.Args[1])
	platformArg := strings.ToUpper(os.Args[2])

	deliveryLogistics, err := logistics.NewLogistics(modeArg)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	factory, err := gui.NewFactory(platformArg)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Delivery mode:", modeArg)
	fmt.Println("UI platform:", platformArg)

	application := app.New(factory, deliveryLogistics)
	application.Run(cargo, destAddr)
}
