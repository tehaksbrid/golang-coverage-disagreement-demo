package foo

import (
	"coverage-bug-example/applications/packages/bar"
	"coverage-bug-example/applications/packages/baz"
)

func main() {
	barStruct := bar.BarStruct{Bar: "bar"}
	baz.Baz(barStruct)
}
