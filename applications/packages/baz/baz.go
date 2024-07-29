package baz

import (
	"golang-coverage-disagreement-demo/applications/packages/bar"
)

func Baz(bar bar.BarStruct) {
	if bar.Bar == "" {
		println(bar.Bar)
	}
}
