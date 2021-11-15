package main

import (
    "fmt"
    "math/rand"
    "time"
    "container/list"
)

func stupidSort(intlist *list.List) {
	intArray:=make([]int, intlist.Len()) 
	var i=0

   for e := intlist.Front(); e != nil; e = e.Next() {
	intArray[i]=e.Value.(int)
	i=i+1
   }


	for !isSorted(intArray, len(intArray) ){
		rand.Seed(time.Now().UTC().UnixNano())

		rand.Shuffle(len(intArray), func(i, j int) {
		intArray[i], intArray[j] = intArray[j], intArray[i]
		});
	}


	intlist.Init()

	for i := 0; i < len(intArray); i++ {
		intlist.PushBack(intArray[i])
	}

}

func isSorted(intArray []int,sizes int) bool{
	var sorted bool

	for i := 0; i < sizes-1; i++ {
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
	intlist := list.New()

	fmt.Println("Genero 10 numeri casuali...")

	for i := 0; i < 5; i++ {
		intlist.PushFront(rand.Intn(100))
	}

	fmt.Println("Ecco i numeri generati:")

   for e := intlist.Front(); e != nil; e = e.Next() {
	fmt.Print(e.Value)
	fmt.Print(" ")
   }


   	
	fmt.Println()
	fmt.Println("\nSto ordinando i valori con lo stupid sort...")

	stupidSort(intlist)

	fmt.Println("Ecco i numeri ordinati:")
   for e := intlist.Front(); e != nil; e = e.Next() {
	fmt.Print(e.Value)
	fmt.Print(" ")
   }
}