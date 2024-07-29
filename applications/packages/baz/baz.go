package baz

import (
	"coverage-bug-example/applications/packages/bar"
)

func Baz(bar bar.BarStruct) {
	if bar.Bar == "" {
		println(bar.Bar)
	}
}
