package ainames

import (
	"fmt"

	"github.com/kasworld/rand"
)

func GetRandomName(rnd *rand.Rand) string {
	lastname := LastNames[rnd.Intn(len(LastNames))]
	firstname := FirstNames[rnd.Intn(len(FirstNames))]
	return fmt.Sprintf("%s %s", lastname, firstname)
}
