package main

import (
    "fmt"
    "math/rand"
    "time"
    "container/list"
)


func mergeSort_rec(intArray []int, p int, r int) {
	if (p < r){
        var q = (p+r)/2
	    mergeSort_rec(intArray, p, q)
	    mergeSort_rec(intArray, q+1, r)
	    merge(intArray, p, q, r)
	}


}

func merge(intArray []int, p int, q int, r int){
	var i=0;
	var j=0;
	var k=0;

	intArray_Support:=make([]int, len(intArray))

	i=p;
	j=q+1;

	for {
        if !(i<=q && j<=r) {
                break
        }
        if intArray[i]<intArray[j] {
			intArray_Support[k]=intArray[i];
			i++;
		} else{
			intArray_Support[k]=intArray[j];
			j++;
		}
		k++;
	}


	for {
		if !(i<=q) {
            break
        }
		intArray_Support[k]=intArray[i];
		i++;
		k++;
	}


	for {
		if !(j<=r) {
            break
        }
		intArray_Support[k]=intArray[j];
		j++;
		k++;
	}

	for k=p;k<=r;k++ {
		intArray[k]=intArray_Support[k-p];
	}

}

func mergeSort(intlist *list.List) {
	intArray:=make([]int, intlist.Len()) 
	var i=0

	for e := intlist.Front(); e != nil; e = e.Next() {
		intArray[i]=e.Value.(int)
		i=i+1
   	}

	mergeSort_rec(intArray, 0, len(intArray)-1)

	intlist.Init()

	for i := 0; i < len(intArray); i++ {
		intlist.PushBack(intArray[i])
	}
}




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

	fmt.Println()
	intlist.Init()

	fmt.Println("\nGenero 5 numeri casuali...")

	for i := 0; i < 5; i++ {
		intlist.PushFront(rand.Intn(100))
	}
	printList(intlist)

	fmt.Println("\nSto ordinando i valori con il merge sort...")
	mergeSort(intlist)

	printList(intlist)

}