package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	testList := [100]string{}
	for i := 0; i < 100; i++ {
		testList[i] = RandStringRunes(256)
	}
	searchFor := rand.Intn(100)
	start := time.Now()
	for _, each := range testList {
		if each == testList[searchFor] {
			break
		}
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
