package logistics

import (
	"fmt"

	"logistics-go/transport"
)

type Logistics interface {
	CreateTransport() transport.Transport
}

func PlanDelivery(l Logistics, cargo, destination string) string {
	t := l.CreateTransport()
	return t.Deliver(cargo, destination)
}

type RoadLogistics struct{}

func (RoadLogistics) CreateTransport() transport.Transport {
	return transport.Truck{}
}

type SeaLogistics struct{}

func (SeaLogistics) CreateTransport() transport.Transport {
	return transport.Ship{}
}

func NewLogistics(mode string) (Logistics, error) {
	switch mode {
	case "ROAD":
		return RoadLogistics{}, nil
	case "SEA":
		return SeaLogistics{}, nil
	default:
		return nil, fmt.Errorf("unsupported delivery mode %q: expected ROAD or SEA", mode)
	}
}
