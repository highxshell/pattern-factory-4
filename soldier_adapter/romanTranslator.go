package soldier_adapter

import (
	"fmt"
	"strategy-as1/soldier_strategy"
)

type RomanTranslator struct {
	RomanSoldier *soldier_strategy.BowSoldier
}

func (r *RomanTranslator) SpeakingGreek() {
	fmt.Println("Translator converts Greek speech to Roman speech.")
	r.RomanSoldier.SpeakingRoman()
}
