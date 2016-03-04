package main

import "fmt"
import "Go-Nes/etai/pinckney/nintendo/opcode"

func main() {
	fmt.Println("Hello, Go")
	codes := opcode.OpCodeTable()
	for _,element := range codes {
		fmt.Println(element.Name())
	}
	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}
}





