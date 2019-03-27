package main

import "fmt"

func bSearch(value int) int {
	a:=[]int{1,3,4,5,8,9,11,42,67}

	low:=0
	high:=len(a)-1

	index:=-1
	for  {
		mid:=int((low+high)/2)

		if value==a[mid]{
			index=mid
			break
		}

		if low>=high{
			break
		}

		if a[mid]>value{
			high=mid-1
		}else {
			low=mid+1
		}
	}
	return index
}

func main() {
  i:= bSearch(67)
  fmt.Println(i)
}
