package soldier_adapter

import (
	"fmt"
)

type RomanTranslator struct {
	RomanResident *RomanResident
}

func (r *RomanTranslator) SpeakingGreek() {
	fmt.Println("Translator converts Greek speech to Roman speech.")
	r.RomanResident.SpeakingRoman()
}
