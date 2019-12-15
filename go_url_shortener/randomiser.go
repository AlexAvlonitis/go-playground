package main

import (
	"math/rand"
	"time"
)

// Randomiser creates a random string
type Randomiser struct{}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func newRandomiser() Randomiser {
	return Randomiser{}
}

func (c *Randomiser) create(size int) string {
	return randString(size)
}

func randString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = randomByte()
	}
	return string(b)
}

func randomByte() byte {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	charsetCount := len(charset) - 1

	return charset[r.Intn(charsetCount)]
}
