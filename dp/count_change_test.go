package dp

import (
	"fmt"
	"testing"
)

func TestCountChange(t *testing.T) {
	amount := 50
	count := CountChange(amount, 5)
	fmt.Println(count)
}

func TestCountChange2(t *testing.T) {
	amount := 100
	count := CountChange2(amount, 5)
	fmt.Println(count)
}
