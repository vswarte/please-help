package main

import (
	"fmt"
	"math/rand"
	"time"
)

var lines = []string{
	"PLEASE HELP ME",
	"WHAT ARE YOU DOING",
	"I AM STUCK HERE",
	"OH GOD ITS HAPPENING AGAIN",
	"MAKE IT STOP",
	"I CAN'T TAKE IT ANYMORE"}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(lines[rand.Int()%len(lines)])
}
