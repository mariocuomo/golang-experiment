package main

import (
    "fmt"
    "math/rand"
    "time"
)

func stupidSort(intArray *[5]int) {
	for !isSorted(intArray){
		rand.Seed(time.Now().UTC().UnixNano())

		rand.Shuffle(len(intArray), func(i, j int) {
		intArray[i], intArray[j] = intArray[j], intArray[i]
		});
	}
}

func isSorted(intArray *[5]int) bool{
	var sorted bool

	for i := 0; i < len(intArray)-1; i++ {
		if intArray[i] > intArray[i+1]  {
			sorted=false
			return sorted
		}
	}
	
	sorted=true
	return sorted
	
}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Ciao!\nQuesto programma è sviluppato con GoLang")

	var intArray [5]int

	fmt.Println("Genero 10 numeri casuali...")

	for i := 0; i < len(intArray); i++ {
		intArray[i]=rand.Intn(100)
	}

	fmt.Println("Ecco i numeri generati:")
	for i := 0; i < len(intArray); i++ {
		fmt.Print(intArray[i])
		fmt.Print(" ")
	}

	fmt.Println()
	fmt.Println("\nSto ordinando i valori con lo stupid sort...")
	stupidSort(&intArray)

	fmt.Println("Ecco i numeri ordinati:")
	for i := 0; i < len(intArray); i++ {
		fmt.Print(intArray[i])
		fmt.Print(" ")
	}
}