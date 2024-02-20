package main

import (
	"fmt"
	"math/rand"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
