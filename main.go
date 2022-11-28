package main

import (
	"ConfKafka/kaffka"
	"fmt"
	"time"
)

func main() {

	fmt.Println("Okey...")

	kaffka.StartKafka()
	fmt.Println("Kaffka started")
	time.Sleep(10 * time.Minute)

}
