package rand_numb

import (
	"math/rand"
)

func Random_n(number int) (num int) {
	num = rand.Intn(number)
	return
}
