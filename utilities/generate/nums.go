package generate

import "math/rand"

func NumsGenerate(min, max int, length int) []int {
	//rand.Seed(time.Now().UnixNano())
	sl := make([]int, 0)
	for {
		n := min + rand.Intn(max-min+1)
		//fmt.Println(gen)
		sl = append(sl, n)
		if len(sl) > length {
			break
		}
	}
	return sl
}
