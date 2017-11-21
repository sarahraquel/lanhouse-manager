package main

import (
	"fmt"
	"time"
	"strconv"
	"math/rand"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func user(channel chan string, id int){
	computer:= <- channel
	myrand := random(15, 30)
	fmt.Println(strconv.Itoa(id) + " está online no " + computer + " por " + strconv.Itoa(myrand) + " minutos")
	time.Sleep(time.Duration(myrand/10) * time.Second)
	fmt.Println(strconv.Itoa(id) + " terminou no " + computer)
	channel <- computer
}


func main() {	
	channel := make(chan string, 8)
	for i := 0; i < 8; i++ {
		channel <- "Pc " + strconv.Itoa(i)
	}
	fmt.Println("Lan House is open")
		
	for i := 1; i < 27; i++ {
		go user(channel, i)
	}
	
	time.Sleep(time.Duration(40) * time.Second)
	fmt.Println("Hello, 世界")
}
