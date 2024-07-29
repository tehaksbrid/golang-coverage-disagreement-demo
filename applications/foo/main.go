package foo

import (
	"golang-coverage-disagreement-demo/applications/packages/bar"
	"golang-coverage-disagreement-demo/applications/packages/baz"
)

func main() {
	barStruct := bar.BarStruct{Bar: "bar"}
	baz.Baz(barStruct)
}
