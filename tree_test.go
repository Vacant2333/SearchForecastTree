package SearchForecastTree

import (
	"fmt"
	"testing"
)

func TestNewTree(t *testing.T) {
	s := " 123"
	for _, c := range s {
		c = 0
		fmt.Printf("--%c %T %d\n", c, c, c)
	}
}
