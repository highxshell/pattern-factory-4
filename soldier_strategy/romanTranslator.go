package soldier_strategy

import "fmt"

type RomanTranslator struct {
	RomanSoldier *BowSoldier
}

func (r *RomanTranslator) SpeakingGreek() {
	fmt.Println("Translator converts Greek speech to Roman speech.")
	r.RomanSoldier.SpeakingRoman()
}
