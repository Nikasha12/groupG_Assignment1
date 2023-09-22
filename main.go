package main

import "fmt"

func main() {
	fmt.Println("Hi")
        fmt.Println(rnumberIncrement())
}

func numberIncrement(){
	initialNumber := 5
	incrementedNumber := increment(initialNumber)

	fmt.Printf("The incremented number is: %d\n", incrementedNumber)
} 
