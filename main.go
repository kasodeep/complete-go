package main // This programs is built into an executable code.

import (
	"flag"
	"fmt"

	"github.com/kasodeep/complete-go/variables"
)

// This function is the start of the program.
func main() {
	fmt.Println("----- Starting Go Demo -----")

	// specify the package.
	packageName := flag.String("package", "variables", "This flag calls the specified package.")
	flag.Parse()

	switch *packageName {
	case "variables":
		variables.Index()
	}
}

// fmt.Println()
// 	fmt.Println("----- Conditionals Demo -----")
// 	conditionals.IfDemo()
// 	conditionals.SwitchDemo()

// 	fmt.Println()
// 	fmt.Println("----- Functions Demo -----")
// 	functions.FunctionsDemo()
// 	functions.TestReports()

// 	fmt.Println()
// 	fmt.Println("----- Structs Demo -----")
// 	structs.StructsDemo()
// 	structs.Anonymous()
// 	structs.EmbeddedDemo()
// 	structs.StructMethods()
// 	structs.SizeDemo()
// 	structs.EmptyStruct()

// 	fmt.Println()
// 	fmt.Println("----- Interfaces Demo -----")
// 	interfaces.InterfaceDemo()
// 	interfaces.StructInterfaceTypes()

// 	fmt.Println()
// 	fmt.Println("----- Errors Demo -----")
// 	errors.ErrorDemo()
// 	errors.CustomError()
// 	errors.ErrorPackage()

// 	fmt.Println()
// 	fmt.Println("----- Loops Demo -----")
// 	loops.ForDemo()

// 	fmt.Println()
// 	fmt.Println("----- Arrays Demo -----")
// 	arrays.ArrayDemo()
// 	arrays.VariadicDemo()
// 	arrays.MatrixDemo()
// 	arrays.RangeDemo()

// 	fmt.Println()
// 	fmt.Println("----- Maps Demo -----")
// 	maps.MapsDemo()

// 	fmt.Println()
// 	fmt.Println("----- Pointers Demo -----")
// 	pointers.PointersDemo()
// 	pointers.ReceiverDemo()
