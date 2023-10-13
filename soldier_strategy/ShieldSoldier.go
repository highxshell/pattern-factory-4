package soldier_strategy

import (
	"fmt"
)

type ShieldSoldier struct {
	Soldier ISoldier
}

func (s *ShieldSoldier) Info() {
	fmt.Println("I'm a Soldier with Shield")
}

func (s *ShieldSoldier) Attack() int {
	return s.Soldier.Attack() - 5
}

func (s *ShieldSoldier) HealthPoints() int {
	return s.Soldier.HealthPoints() + 50
}

func (s *ShieldSoldier) SpeakingGreek() {
	fmt.Println("*speaks Greek*")
}
