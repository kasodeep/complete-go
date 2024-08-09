package main // This programs is built into an executable code.

import (
	"flag"

	"github.com/kasodeep/complete-go/arrays"
	"github.com/kasodeep/complete-go/channels"
	"github.com/kasodeep/complete-go/conditionals"
	"github.com/kasodeep/complete-go/errors"
	files "github.com/kasodeep/complete-go/file-handling"
	"github.com/kasodeep/complete-go/functions"
	"github.com/kasodeep/complete-go/generics"
	"github.com/kasodeep/complete-go/interfaces"
	"github.com/kasodeep/complete-go/loops"
	"github.com/kasodeep/complete-go/maps"
	"github.com/kasodeep/complete-go/mutexes"
	"github.com/kasodeep/complete-go/pointers"
	"github.com/kasodeep/complete-go/strings"
	"github.com/kasodeep/complete-go/structs"
	"github.com/kasodeep/complete-go/variables"
)

// This function is the start of the program.
func main() {
	// specify the package.
	packageName := flag.String("package", "channels", "This flag calls the specified package.")
	flag.Parse()

	// calling the appropriate package.
	switch *packageName {
	case "variables":
		variables.Index()
	case "structs":
		structs.Index()
	case "conditionals":
		conditionals.Index()
	case "functions":
		functions.Index()
	case "interfaces":
		interfaces.Index()
	case "errors":
		errors.Index()
	case "loops":
		loops.Index()
	case "arrays":
		arrays.Index()
	case "maps":
		maps.Index()
	case "pointers":
		pointers.Index()
	case "strings":
		strings.Index()
	case "mutexes":
		mutexes.Index()
	case "generics":
		generics.Index()
	case "files":
		files.Index()
	case "channels":
		channels.Index()
	}
}
