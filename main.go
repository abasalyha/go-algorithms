package main

import "fmt"


func swap(items *[]int, left int, right int) {

	if left != right {
		var temp = (*items)[left]
		(*items)[left] = (*items)[right]
		(*items)[right] = temp
	}
}


func bubble_sort(items []int){
	var swapped = true
	for swapped != false {
		swapped = false;
		var i = 1
		for i < len(items) {
			if items[i - 1] > items[i]{
				swap(&items, i - 1, i)
				swapped = true;
			}
			i++
		}

	}
}

func main(){
	x:= []int{7, 3, 4, 6}
	bubble_sort(x)
	fmt.Println(x)

}
