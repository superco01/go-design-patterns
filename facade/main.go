package facade

import "fmt"

type FuelEngine struct{}

func (FuelEngine) checkFuel() {
	fmt.Println("Checking fuel")
}

func (FuelEngine) injectFuel() {
	fmt.Println("Injecting fuel")
}

type IgnitionEngine struct{}

func (IgnitionEngine) iginiteSparkPlug() {
	fmt.Println("Igniting spark plug")
}

type CarFacade struct {
	FuelEngine     FuelEngine
	IgnitionEngine IgnitionEngine
}

func NewCarFacade() CarFacade {
	return CarFacade{
		FuelEngine: NewCarFacade().FuelEngine,
		IgnitionEngine: NewCarFacade().IgnitionEngine,
	}
}

func (c CarFacade) Start()  {
	c.FuelEngine.checkFuel()
	c.FuelEngine.injectFuel()
	c.IgnitionEngine.iginiteSparkPlug()
}

func main()  {
	car := NewCarFacade()
	car.Start()
}
