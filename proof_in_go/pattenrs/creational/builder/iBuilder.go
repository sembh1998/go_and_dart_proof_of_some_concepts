package builder

const (
	NormalBuilderType = "normal"
	IglooBuilderType  = "igloo"
)

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
	GetType() string
}

func GetBuilder(builderType string) IBuilder {
	if builderType == NormalBuilderType {
		return NewNormalBuilder()
	}

	if builderType == IglooBuilderType {
		return NewIglooBuilder()
	}
	return nil
}
