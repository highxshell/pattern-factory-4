package soldier_adapter

import (
	"fmt"
)

type RomanResident struct {
}

func (b *RomanResident) SpeakingRoman() {
	fmt.Println("*speaks Roman*")
}
