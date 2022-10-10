package ip

import "fmt"

func generate(i, j, k, l int) string {
	return fmt.Sprint(i) + "." + fmt.Sprint(j) + "." + fmt.Sprint(k) + "." + fmt.Sprint(l)
}

func GenerateAll() []string {
	print("Generating all ips...")
	var ips []string
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			for k := 0; k < 256; k++ {
				for l := 0; l < 256; l++ {
					ips = append(ips, generate(i, j, k, l))
				}
			}
		}
	}
	return ips
}
