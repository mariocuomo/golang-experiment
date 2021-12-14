package main

import (
    "fmt"
    "math/rand"
    "time"
    "container/list"
)

type Node struct {
	value    int
	left     *Node
	right     *Node
}


/*
=================================
          TREE SORT
=================================
*/
func treeSort(intlist *list.List) {
	intArray:=make([]int, intlist.Len()) 
	var i=0

	for e := intlist.Front(); e != nil; e = e.Next() {
		intArray[i]=e.Value.(int)
		i=i+1
	}


	btree := Node{
        value: intArray[0],
        left:    nil,
        right:    nil,
    }



    for i := 1; i < len(intArray); i++ {
		tmp := (&btree) 
		for {
			if (intArray[i] < (tmp).value){
				if(tmp.left!=nil){
					tmp=tmp.left
				}else{
					node := Node{
				        value: intArray[i],
				        left:    nil,
				        right:    nil,
				    }
				    tmp.left=&node
				    break
				}
			}else{
				if(tmp.right!=nil){
					tmp=tmp.right
				}else{
					node := Node{
				        value: intArray[i],
				        left:    nil,
				        right:    nil,
				    }
				    tmp.right=&node
				    break
				}
			}
	        
		}
	}

	intArray=visitTree(&btree)

	intlist.Init()

	for i := 0; i < len(intArray); i++ {
		intlist.PushBack(intArray[i])
	}
}

func visitTree(btree *Node) []int{
	if(btree==nil){
		ll:= make([]int, 1)
		return ll
	}

	if(btree.left==nil && btree.right==nil){
		ll:= make([]int, 1)
		ll[0]=btree.value
		return ll
	}

	if(btree.left==nil){
		ll:= make([]int, 1)
		ll[0]=btree.value

		right := visitTree(btree.right)
		var result = make([]int, 1 + len(right))

 		copy(result[1:], right[:])
		result[0]=ll[0]

		return result
	}

	if(btree.right==nil){
		ll:= make([]int, 1)
		ll[0]=btree.value

		left := visitTree(btree.left)
		var result = make([]int, 1 + len(left))

 		copy(result[:], left[:])
		result[len(result)-1]=ll[0]

		return result
	}

	left := visitTree(btree.left)
	right := visitTree(btree.right)	
 	var result = make([]int, len(left) + len(right) + 1)
 	copy(result[:], left[:])
 	result[len(left)]=btree.value
	copy(result[len(left)+1:], right[:])

	return result
}


/*
=================================
          COUNTING SORT
=================================
*/
func countingSort(intlist *list.List) {
	intArray:=make([]int, intlist.Len()) 
	var i=0
	var max=intlist.Front().Value.(int)
	var min=intlist.Front().Value.(int)

   for e := intlist.Front(); e != nil; e = e.Next() {
	intArray[i]=e.Value.(int)
	if (intArray[i]>max){
		max=intArray[i]
	}
	if (intArray[i]<min){
		min=intArray[i]
	}
	i=i+1
   }

   if(min<0){
   	min=min*(-1)
   }

	countingArray:=make([]int, max+min+1) 

	for i := 0; i <len(intArray); i++ {
		countingArray[intArray[i]+min]=countingArray[intArray[i]+min]+1
	}

	var j=0
	for i := 0; i <len(countingArray); i++ {
		if(countingArray[i]!=0){
			var times=countingArray[i]
			for k := 0; k<times;k++{
				intArray[j]=i-min
				j=j+1
			}
		}
		
	}


	intlist.Init()

	for i := 0; i < len(intArray); i++ {
		intlist.PushBack(intArray[i])
	}
}

/*
=================================
          BUCKET SORT
=================================
*/
func bucketSort(intlist *list.List) {
	intArray:=make([]int, intlist.Len()) 
	var i=0
	var max=intlist.Front().Value.(int)
	var min=intlist.Front().Value.(int)

   for e := intlist.Front(); e != nil; e = e.Next() {
		intArray[i]=e.Value.(int)
		if (intArray[i]>max){
			max=intArray[i]
		}
		if (intArray[i]<min){
			min=intArray[i]
		}
		i=i+1
   }

   if(min<0){
   	min=min*(-1)
   }

	bucketArray:=make([]*list.List, max+min+1)


	
	for i := 0; i <len(intArray); i++ {
		if(bucketArray[intArray[i]+min] == nil){
			bucketArray[intArray[i]+min]=list.New()
		}
		bucketArray[intArray[i]+min].PushFront(intArray[i])
	}

	
	var j=0
	for i := 0; i <len(bucketArray); i++ {
		if(bucketArray[i]!=nil){
			for e := bucketArray[i].Front(); e != nil; e = e.Next() {
				intArray[j]=e.Value.(int)
				j=j+1
   			}
		}
	}

	intlist.Init()
	

	for i := 0; i < len(intArray); i++ {
		intlist.PushBack(intArray[i])
	}
	
}


/*
=================================
           BUBBLE SORT
=================================
*/
func bubbleSort(intlist *list.List) {
	intArray:=make([]int, intlist.Len()) 
	var i=0

   for e := intlist.Front(); e != nil; e = e.Next() {
	intArray[i]=e.Value.(int)
	i=i+1
   }

	
	for i := len(intArray)-1; i >=0; i-- {
		for j := 0; j < i; j++ {
			if(intArray[j]>intArray[j+1]){
				var tmp=intArray[j+1]
				intArray[j+1]=intArray[j]
				intArray[j]=tmp
			}
		}
	}

	intlist.Init()

	for i := 0; i < len(intArray); i++ {
		intlist.PushBack(intArray[i])
	}
}


/*
=================================
           INSERTION SORT
=================================
*/
func insertionSort(intlist *list.List) {
	intArray:=make([]int, intlist.Len()) 
	var i=0

   for e := intlist.Front(); e != nil; e = e.Next() {
	intArray[i]=e.Value.(int)
	i=i+1
   }

	
	for i := 0; i <len(intArray)-1; i++ {
		for j := i+1; j > 0; j-- {
			if(intArray[j]<intArray[j-1]){
				var tmp = intArray[j]
				intArray[j]=intArray[j-1]
				intArray[j-1]=tmp
			}
		}
	}

	intlist.Init()

	for i := 0; i < len(intArray); i++ {
		intlist.PushBack(intArray[i])
	}
}



/*
=================================
           MERGE SORT
=================================
*/
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



/*
=================================
        SELECTION SORT
=================================
*/
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



/*
=================================
           STUPID SORT
=================================
*/
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



/*
=================================
         SUPPORT FUNCTIONS
=================================
*/
func printList(intlist *list.List){
	for e := intlist.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
		fmt.Print(" ")
	   }
}

func printArray(intArray []int,sizes int){
	for i := 0; i < sizes; i++ {
		fmt.Print(intArray[i])
		fmt.Print(" ")
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



/*
================================
          MAIN FUNCTION
================================
*/

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Hello!\nThis program is developed using GoLang!\n")
	intlist := list.New()

	fmt.Println("I am generating 5 random numbers...")

	for i := 0; i < 5; i++ {
		intlist.PushFront(rand.Intn(100))
	}

	fmt.Print()

   	printList(intlist)
	fmt.Println("\nI'm sorting them using stupid sort...")

	stupidSort(intlist)

	printList(intlist)

	fmt.Println("\n")
	intlist.Init()

	fmt.Println("I am generating 5 random numbers...")

	for i := 0; i < 5; i++ {
		intlist.PushFront(rand.Intn(100))
	}
	printList(intlist)

	fmt.Println("\nI'm sorting them using selection sort...")
	selectionSort(intlist)

	printList(intlist)

	fmt.Println("\n")
	intlist.Init()

	fmt.Println("I am generating 5 random numbers...")

	for i := 0; i < 5; i++ {
		intlist.PushFront(rand.Intn(100))
	}
	printList(intlist)

	fmt.Println("\nI'm sorting them using merge sort...")
	mergeSort(intlist)

	printList(intlist)

	fmt.Println("\n")
	intlist.Init()

	fmt.Println("I am generating 5 random numbers...")

	for i := 0; i < 5; i++ {
		intlist.PushFront(rand.Intn(100))
	}
	printList(intlist)

	fmt.Println("\nI'm sorting them using bubble sort...")
	bubbleSort(intlist)

	printList(intlist)

	fmt.Println("\n")
	intlist.Init()

	fmt.Println("I am generating 5 random numbers...")

	for i := 0; i < 5; i++ {
		intlist.PushFront(rand.Intn(100))
	}
	printList(intlist)

	fmt.Println("\nI'm sorting them using tree sort...")
	treeSort(intlist)

	printList(intlist)

	fmt.Println("\n")
	intlist.Init()

	fmt.Println("I am generating 5 random numbers...")

	for i := 0; i < 5; i++ {
		intlist.PushFront(rand.Intn(100))
	}
	printList(intlist)

	fmt.Println("\nI'm sorting them using counting sort...")
	countingSort(intlist)

	printList(intlist)

	fmt.Println("\n")
	intlist.Init()

	fmt.Println("I am generating 5 random numbers...")

	for i := 0; i < 5; i++ {
		intlist.PushFront(rand.Intn(100))
	}
	printList(intlist)

	fmt.Println("\nI'm sorting them using insertion sort...")
	insertionSort(intlist)

	printList(intlist)

	fmt.Println("\n")
	intlist.Init()

	fmt.Println("I am generating 5 random numbers...")

	for i := 0; i < 5; i++ {
		intlist.PushFront(rand.Intn(100))
	}
	printList(intlist)

	fmt.Println("\nI'm sorting them using bucket sort...")
	bucketSort(intlist)

	printList(intlist)

	fmt.Println("\n")

}