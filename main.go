package main

import (
	"fmt"
	"strategy-as1/soldier_strategy"
)

func main() {
	//singleton
	soldier := soldier_strategy.GetBasicSoldier()
	singletonSoldier := soldier_strategy.GetBasicSoldier()
	basicSoldier := soldier
	BasicSoldier1 := singletonSoldier
	//strategy/decorator
	bowSoldier := soldier_strategy.BowSoldier{Soldier: basicSoldier}
	shieldSoldier := soldier_strategy.ShieldSoldier{Soldier: basicSoldier}
	shieldBowSoldier := soldier_strategy.ShieldSoldier{
		Soldier: &soldier_strategy.BowSoldier{
			Soldier: basicSoldier,
		},
	}
	fmt.Println()

	soldier1 := soldier_strategy.SoldierBehavior{SB: basicSoldier}
	soldier2 := soldier_strategy.SoldierBehavior{SB: &bowSoldier}
	soldier3 := soldier_strategy.SoldierBehavior{SB: &shieldSoldier}
	soldier4 := soldier_strategy.SoldierBehavior{SB: &shieldBowSoldier}
	soldier5 := soldier_strategy.SoldierBehavior{SB: BasicSoldier1}

	soldier1.DisplayStats()
	soldier2.DisplayStats()
	soldier3.DisplayStats()
	soldier4.DisplayStats()
	soldier5.DisplayStats()

	//adapter
	fmt.Println()
	adapter := &soldier_strategy.Adapter{}
	soldier6 := &soldier_strategy.ShieldSoldier{}
	adapter.TranslateTwoLanguages(soldier6)
	soldier7 := &soldier_strategy.BowSoldier{}
	romanTranslator := &soldier_strategy.RomanTranslator{
		RomanSoldier: soldier7,
	}
	adapter.TranslateTwoLanguages(romanTranslator)
}
