package soldier_strategy

import "fmt"

type BowSoldier struct {
	Soldier ISoldier
}

func (b *BowSoldier) Info() {
	fmt.Println("I'm a Soldier with Bow")
}

func (b *BowSoldier) Attack() int {
	return b.Soldier.Attack() + 15
}

func (b *BowSoldier) HealthPoints() int {
	return b.Soldier.HealthPoints() - 20
}

func (b *BowSoldier) SpeakingRoman() {
	fmt.Println("*speaks Roman*")
}
