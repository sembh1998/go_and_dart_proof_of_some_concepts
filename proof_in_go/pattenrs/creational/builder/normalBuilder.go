package builder

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) SetWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) SetDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) SetNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) GetHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

func (b *NormalBuilder) GetType() string {
	return NormalBuilderType
}
