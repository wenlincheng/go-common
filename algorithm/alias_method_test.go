package algorithm

import (
	"fmt"
	"testing"
)

func TestAliasMethod(t *testing.T) {
	probability := []float64{0.5, 0.4, 0.05, 0.04, 0.007, 0.003}
	str := []string{"1积分", "2积分", "3积分", "9积分", "19积分", "29积分"}
	aliasMethod := AliasMethod{}
	err := aliasMethod.Initialize(probability)
	if err != nil {
		return
	}
	for i := 0; i < 10; i++ {
		fmt.Print(str[aliasMethod.Next()])
	}
}
