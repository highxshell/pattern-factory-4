package main

import (
	"fmt"
	"strategy-as1/config"
	"strategy-as1/factory"
	"strategy-as1/observer"
	"strategy-as1/soldier_adapter"
	"strategy-as1/soldier_strategy"
)

func main() {
	//singleton
	instanceOfDB := config.GetInstanceOfDB()
	instanceOfDB.Info()
	configSingleton := config.GetInstanceOfDB()
	configSingleton.Info()
	//decorator/strategy
	basicSoldier := soldier_strategy.BasicSoldier{}
	bowSoldier := soldier_strategy.BowSoldier{Soldier: &basicSoldier}
	shieldSoldier := soldier_strategy.ShieldSoldier{Soldier: &basicSoldier}
	shieldBowSoldier := soldier_strategy.ShieldSoldier{
		Soldier: &soldier_strategy.BowSoldier{
			Soldier: &basicSoldier,
		},
	}
	fmt.Println()

	soldier1 := soldier_strategy.SoldierBehavior{SB: &basicSoldier}
	soldier2 := soldier_strategy.SoldierBehavior{SB: &bowSoldier}
	soldier3 := soldier_strategy.SoldierBehavior{SB: &shieldSoldier}
	soldier4 := soldier_strategy.SoldierBehavior{SB: &shieldBowSoldier}

	soldier1.DisplayStats()
	soldier2.DisplayStats()
	soldier3.DisplayStats()
	soldier4.DisplayStats()

	//adapter
	fmt.Println()
	adapter := &soldier_adapter.Adapter{}
	greekResident := &soldier_adapter.GreekResident{}
	adapter.TranslateTwoLanguages(greekResident)
	romanResident := &soldier_adapter.RomanResident{}
	romanTranslator := &soldier_adapter.RomanTranslator{
		RomanResident: romanResident,
	}
	adapter.TranslateTwoLanguages(romanTranslator)

	//factory
	fmt.Println()
	archer, _ := factory.GetSoldier("Archer")
	shieldbearer, _ := factory.GetSoldier("Shieldbearer")

	printDetails(archer)
	printDetails(shieldbearer)

	//observer
	fmt.Println()
	legendaryBowItem := observer.NewItem("bow of the Galadhrim")
	legendaryShieldItem := observer.NewItem("Boromir's shield")
	itemManager := observer.NewItemManager(legendaryBowItem)
	itemManager1 := observer.NewItemManager(legendaryShieldItem)
	observerPlayer1 := &observer.Customer{Id: "player1@gmail.com"}
	observerPlayer2 := &observer.Customer{Id: "player2@gmail.com"}

	itemManager.Register(observerPlayer1)
	itemManager.Register(observerPlayer2)
	itemManager.UpdateAvailability()
	itemManager1.Register(observerPlayer2)
	itemManager.DeRegister(observerPlayer2)
	itemManager.UpdateAvailability()
	itemManager1.UpdateAvailability()
}

//factory

func printDetails(i soldier_strategy.ISoldier) {
	fmt.Printf("Soldier: %s", i.GetName())
	fmt.Println()
	fmt.Printf("Attack: %d", i.GetAttack())
	fmt.Println()
	fmt.Printf("HP: %d", i.GetHP())
	fmt.Println()
}
