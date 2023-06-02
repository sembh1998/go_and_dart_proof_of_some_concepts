package builder_test

import (
	"fmt"
	"testing"

	"github.com/sembh1998/go_and_dart_proof_of_some_concepts/proof_in_go/pattenrs/creational/builder"
	"github.com/stretchr/testify/assert"
)

func TestBuilderSuccess(t *testing.T) {
	normalBuilder := builder.GetBuilder(builder.NormalBuilderType)
	iglooBuilder := builder.GetBuilder(builder.IglooBuilderType)

	director := builder.NewDirector(normalBuilder)
	assert.Equal(t, builder.NormalBuilderType, director.GetBuilderType())
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.GetDoorType())
	assert.Equal(t, "Wooden Door", normalHouse.GetDoorType())
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.GetWindowType())
	assert.Equal(t, "Wooden Window", normalHouse.GetWindowType())
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.GetFloor())
	assert.Equal(t, 2, normalHouse.GetFloor())

	director.SetBuilder(iglooBuilder)
	assert.Equal(t, builder.IglooBuilderType, director.GetBuilderType())
	iglooHouse := director.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.GetDoorType())
	assert.Equal(t, "Snow Door", iglooHouse.GetDoorType())
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.GetWindowType())
	assert.Equal(t, "Snow Window", iglooHouse.GetWindowType())
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.GetFloor())
	assert.Equal(t, 1, iglooHouse.GetFloor())

	// Output:
	// Normal House Door Type: Wooden Door
	// Normal House Window Type: Wooden Window
	// Normal House Num Floor: 2
	//
	// Igloo House Door Type: Snow Door
	// Igloo House Window Type: Snow Window
	// Igloo House Num Floor: 1
}
