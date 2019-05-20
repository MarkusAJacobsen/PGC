package internal

import (
	"math/rand"
	"time"
)

type Id struct {
	sRand *rand.Rand
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const DefaultLength uint8 = 12

func GetIdGenerator() (idGen Id) {
	return Id{
		sRand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (id *Id) GenId(length *uint8) (gId string) {
	if length == nil {
		tmp := DefaultLength // Cannot find address of constant
		length = &tmp
	}
	b := make([]byte, *length)
	for i := range b {
		b[i] = charset[id.sRand.Intn(len(charset))]
	}
	return string(b)
}

// Inspo: https://www.calhoun.io/creating-random-strings-in-go/