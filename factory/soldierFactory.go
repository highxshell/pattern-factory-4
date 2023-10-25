package factory

import (
	"fmt"
	"strategy-as1/soldier_strategy"
)

func GetSoldier(soldierType string) (soldier_strategy.ISoldier, error) {
	if soldierType == "Archer" {
		return soldier_strategy.NewBowSoldier(), nil
	}
	if soldierType == "Shieldbearer" {
		return soldier_strategy.NewShieldSoldier(), nil
	}
	return nil, fmt.Errorf("wrong soldier type passed")
}
