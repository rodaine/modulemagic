package renamedmodule

import (
				"github.com/rodaine/modulemagic"
				"github.com/rodaine/renamedmodule/v2/foo"
)

func Sub(a, b int) int {
				return modulemagic.Add(a, foo.Negate(b))
}

func Div(a, b int) int {
				r := a % b
				return (a-r)/b
}
