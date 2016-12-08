package main

import (
        "fmt"
        "math/rand"
        "time"
)

var lines = [4]string{"PLEASE HELP ME", "WHAT ARE YOU DOING", "I AM STUCK HERE", "OH GOD ITS HAPPENING AGAIN"}

func main() {
        rand.Seed(time.Now().Unix())
        fmt.Println(lines[rand.Int()%len(lines)])
}
