package ip

import "fmt"

func Generate(i, j, k, l int) string {
	return fmt.Sprint(i) + "." + fmt.Sprint(j) + "." + fmt.Sprint(k) + "." + fmt.Sprint(l)
}
