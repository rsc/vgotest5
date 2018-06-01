package main // import "github.com/rsc/vgotest5"

import (
	"fmt"

	"github.com/myitcv/vgo_example_compat/v2"
	"github.com/myitcv/vgo_example_compat/v2/sub"
)

func main() {
	fmt.Printf("X=%d, Y=%d\n", vgo_example_compat.X, sub.Y)
}
