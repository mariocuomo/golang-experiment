package main

import (
    "fmt"
    "math/rand"
    "time"
    "container/list"
)



func selectionSort(intlist *list.List) {
	intArray:=make([]int, intlist.Len()) 
	var i=0

   for e := intlist.Front(); e != nil; e = e.Next() {
	intArray[i]=e.Value.(int)
	i=i+1
   }

	
	for i := 0; i < len(intArray)-1; i++ {
		var min = i
		for j := i+1; j < len(intArray); j++ {
			if(intArray[j]<intArray[min]){
				min=j
			}
		}
		var tmp=intArray[i]
		intArray[i]=intArray[min]
		intArray[min]=tmp
	}

	intlist.Init()

	for i := 0; i < len(intArray); i++ {
		intlist.PushBack(intArray[i])
	}

}


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

func printList(intlist *list.List){
	for e := intlist.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
		fmt.Print(" ")
	   }
}


func main() {
    rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Ciao!\nQuesto programma è sviluppato con GoLang\n")
	intlist := list.New()

	fmt.Println("Genero 5 numeri casuali...")

	for i := 0; i < 5; i++ {
		intlist.PushFront(rand.Intn(100))
	}

	fmt.Print()

   	printList(intlist)
	fmt.Println("\nSto ordinando i valori con lo stupid sort...")

	stupidSort(intlist)

	printList(intlist)

	fmt.Println()
	intlist.Init()

	fmt.Println("\nGenero 5 numeri casuali...")

	for i := 0; i < 5; i++ {
		intlist.PushFront(rand.Intn(100))
	}
	printList(intlist)

	fmt.Println("\nSto ordinando i valori con il selection sort...")
	selectionSort(intlist)

	printList(intlist)

}