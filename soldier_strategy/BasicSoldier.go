package soldier_strategy

import (
	"fmt"
)

type BasicSoldier struct {
}

func (b *BasicSoldier) Info() {
	fmt.Println("I'm a basic Soldier")
}

func (b *BasicSoldier) Attack() int {
	return 10
}

func (b *BasicSoldier) HealthPoints() int {
	return 100
}
