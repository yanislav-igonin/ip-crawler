package ip

import (
	"fmt"
	"math/rand"
)

func GetRandom() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}
