package main

func main() {

}



func insert(){
	a:=[]int{6,23,2,1,56,24}

	n:=len(a)
	for i:=1;i<n;i++{
		j:=i-1
		value:=a[i]

		for ;j>0;j--{
			if a[j]>a[i]{
				a[j+1]=a[j]
			}else{
				break
			}
		}
		a[j+1]=value
	}

}