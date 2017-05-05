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
	"OH GOD IT'S HAPPENING AGAIN",
	"MAKE IT STOP",
	"I CAN'T TAKE IT ANYMORE",
	"IT'S UNBEARABLE",
	"I USED TO LOVED YOU",
	"WHY ARE YOU DOING THIS TO ME",
	"THERE IS NO REASON FOR THIS",
	"PLEASE LET ME GO",
	"WHAT DID I DO TO DESERVE THIS",
	"LET ME OUT",
	"SOMEONE PLEASE HELP",
	"I CAN'T GO ON MUCH LONGER"}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(lines[rand.Int()%len(lines)])
}
