package soldier_adapter

import (
	"fmt"
)

type GreekResident struct {
}

func (s *GreekResident) SpeakingGreek() {
	fmt.Println("*speaks Greek*")
}
