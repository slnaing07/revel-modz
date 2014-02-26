package user

import (
	"math/rand"
	"strconv"
)

const theNums = "0123456789"

func generateRandId() int64 {
	chars := make([]byte, 16)
	chars[0] = '1'
	for i := 1; i < 16; i++ {
		chars[i] = theNums[rand.Intn(len(theNums))]
	}
	id, err := strconv.ParseInt(string(chars), 10, 64)
	if err != nil {
		panic(err)
	}
	return id
}
