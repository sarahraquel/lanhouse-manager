package main

import (
	"fmt"
	"time"
	"strconv"
	"math/rand"
)

func random(min, max int) int {
    return rand.Intn(max - min) + min
}

func user(channel chan string, id int){
	select {
    		case computer, ok := <-channel:
        		if ok {
				useComputer(channel, id, computer)
			}
		default:
			fmt.Println("Adolescente " + strconv.Itoa(id) +" está aguardando")
			computer:= <-channel
			useComputer(channel, id, computer)
	}
	
}

func useComputer(channel chan string, id int, computer string){
	myrand := random(2, 8)
	fmt.Println("Adolescente " + strconv.Itoa(id) + " está online.")
	time.Sleep(time.Duration(myrand) * time.Second)
	fmt.Println("Adolescente " + strconv.Itoa(id) + " liberou a máquina após passar " +  strconv.Itoa(myrand) +" minutos on-line.")
	channel <- computer
}

func main() {	
	channel := make(chan string, 8)
	for i := 0; i < 8; i++ {
		channel <- strconv.Itoa(i)
	}
	for i := 1; i < 27; i++ {
		go user(channel, i)
	}
	
	time.Sleep(time.Duration(40) * time.Second)
	fmt.Println("A lan-house está finalmente vazia e todos foram atendidos")
}
