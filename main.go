package main // This programs is built into an executable code.

import (
	"Go/arrays"
	"Go/conditionals"
	"Go/errors"
	"Go/functions"
	"Go/interfaces"
	"Go/loops"
	"Go/structs"
	"Go/variables"
	"fmt"
)

func main() {
	fmt.Println("----- Starting Go Demo -----")
	fmt.Println("Hello, Go!")

	fmt.Println()
	fmt.Println("----- Variables Demo -----")
	variables.Types()
	variables.ShortNotation()
	variables.TypeConversion()
	variables.Constants()
	variables.LogPrinting()
	variables.RuneDemo()

	fmt.Println()
	fmt.Println("----- Conditionals Demo -----")
	conditionals.IfDemo()
	conditionals.SwitchDemo()

	fmt.Println()
	fmt.Println("----- Functions Demo -----")
	functions.FunctionsDemo()
	functions.TestReports()

	fmt.Println()
	fmt.Println("----- Structs Demo -----")
	structs.StructsDemo()
	structs.Anonymous()
	structs.EmbeddedDemo()
	structs.StructMethods()
	structs.SizeDemo()
	structs.EmptyStruct()

	fmt.Println()
	fmt.Println("----- Interfaces Demo -----")
	interfaces.InterfaceDemo()
	interfaces.StructInterfaceTypes()

	fmt.Println()
	fmt.Println("----- Errors Demo -----")
	errors.ErrorDemo()
	errors.CustomError()
	errors.ErrorPackage()

	fmt.Println()
	fmt.Println("----- Loops Demo -----")
	loops.ForDemo()

	fmt.Println()
	fmt.Println("----- Arrays Demo -----")
	arrays.ArrayDemo()
	arrays.VariadicDemo()
	arrays.MatrixDemo()
	arrays.RangeDemo()
}