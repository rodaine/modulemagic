package renamedmodule

import "testing"

func Test(t *testing.T) {
				diff := Sub(3, 2)

				if diff != 1 {
								t.Fatal(diff)
				}
}
