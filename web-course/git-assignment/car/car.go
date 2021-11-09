package car

type Car struct {
	Brand        Brand
	Model        string
	Color        Color
	currentSpeed *int
}

type Color string

const (
	Red   Color = "red"
	Blue        = "blue"
	green       = "green"
)

type Brand string

const (
	BMW    Brand = "bmw"
	Benz         = "benz"
	Nissan       = "nissan"
)

func NewCar(bn Brand, model string, color Color) Car {
	cr := new(int)
	*cr = 0
	return Car{
		Brand:        bn,
		Model:        model,
		Color:        color,
		currentSpeed: cr,
	}
}

func (c Car) SpeedUp() {
	*c.currentSpeed++
}

func (c Car) SlowDown() {
	if *c.currentSpeed == 0 {
		return
	}
	*c.currentSpeed--
}

func (c Car) GetSpeed() int {
	return *c.currentSpeed
}
