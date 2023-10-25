package soldier_strategy

import (
	"fmt"
)

type BasicSoldier struct {
	attack int
	HP     int
	name   string
}

// factory
func (b *BasicSoldier) setAttack(attack int) {
	b.attack = attack
}

func (b *BasicSoldier) setName(name string) {
	b.name = name
}

func (b *BasicSoldier) setHP(HP int) {
	b.HP = HP
}

func (b *BasicSoldier) GetAttack() int {
	return b.attack
}
func (b *BasicSoldier) GetName() string {
	return b.name
}

func (b *BasicSoldier) GetHP() int {
	return b.HP
}

//factory

func (b *BasicSoldier) Info() {
	fmt.Println("I'm a basic Soldier")
}

func (b *BasicSoldier) Attack() int {
	return 10
}

func (b *BasicSoldier) HealthPoints() int {
	return 100
}
